package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/afiefafian/todo-api/src/entity"
	"log"
	"strings"
	"time"
)

type registrationServices struct {
	userRepo         entity.UserRepository
	registrationRepo entity.RegistrationRepository
	tokenRepo        entity.TokenRepository
	cache            entity.CacheRepository
}

// NewRegistrationServices create new registration services
func NewRegistrationServices(u entity.UserRepository, r entity.RegistrationRepository, t entity.TokenRepository, c entity.CacheRepository) entity.RegistrationServices {
	return &registrationServices{
		userRepo:         u,
		registrationRepo: r,
		tokenRepo:        t,
		cache:            c,
	}
}

var cacheKey = map[string]string{
	"cooldown": "cooldown",
	"retry":    "retry",
}

// Register user data to temporary and send register otp to user
func (r *registrationServices) Register(ctx context.Context, registration *entity.Registration) (string, error) {
	var (
		user  entity.User
		err   error
		email = strings.Trim(registration.Email, "")
		// Cache Key
		cooldownKey = fmt.Sprintf("%s:%s", email, cacheKey["cooldown"])
		retryKey    = fmt.Sprintf("%s:%s", email, cacheKey["retry"])
	)

	if user, err = r.userRepo.GetByEmail(ctx, email); err != nil {
		return "", err
	}
	if user != (entity.User{}) {
		return "", errors.New("invalidField: email:Email already registered")
	}

	// Check register status
	if err = r.registerStatus(ctx, email); err != nil {
		return "", err
	}

	// Check retry
	retry := r.cache.Get(ctx, retryKey)
	if retry.Val() != "" && retry.Val() == "0" {
		ttl := r.cache.TTL(ctx, retryKey)
		errMsg := fmt.Sprintf("Too much retry, please wait %d minutes", int64(time.Duration(ttl)/time.Minute))
		return "", errors.New(errMsg)
	}

	// Check cooldown timer
	cooldown := r.cache.Get(ctx, cooldownKey)
	if cooldown.Val() != "" {
		ttl := r.cache.TTL(ctx, cooldownKey)
		errMsg := fmt.Sprintf("Wait until %d seconds", int64(time.Duration(ttl)/time.Second))
		return "", errors.New(errMsg)
	}

	// Set cooldown timer
	if err = r.cache.SetWithTTL(ctx, &entity.Cache{Key: cooldownKey, Value: email, TTL: 60}); err != nil {
		log.Printf("Fail set token lifetime : %s", err)
		return "", err
	}

	// Set or decrement limiter value for 15 minutes
	if retry.Val() == "" {
		if err = r.cache.SetWithTTL(ctx, &entity.Cache{Key: retryKey, Value: "3", TTL: 15 * 60}); err != nil {
			log.Printf("Fail set request limiter : %s", err)
			return "", err
		}
	} else if retry.Val() > "0" {
		if err = r.cache.Decrement(ctx, retryKey); err != nil {
			log.Printf("Fail increment request limiter : %s", err)
			return "", err
		}
	}

	// If email is not used
	// Generate register code
	var regCode string
	if regCode, err = r.generateOTPCode(6); err != nil {
		log.Printf("Failed generate otp code : %s", err)
		return "", err
	}

	// Save register  code and data
	// Deactivate all register token by identifier
	if err = r.tokenRepo.DeactivateAllByIdentifierAndType(ctx, email, "register"); err != nil {
		log.Printf("Fail deactive token : %s", err)
		return "", err
	}

	if err = r.tokenRepo.CreateNewToken(ctx, &entity.Token{Identifier: email, Type: "register", Code: regCode}); err != nil {
		log.Printf("Fail save token : %s", err)
		return "", err
	}

	if err = r.registrationRepo.StoreOrUpdateIfEmailExist(ctx, registration); err != nil {
		log.Printf("Fail upsert user register data : %s", err)
		return "", err
	}

	// Set register code to redis for 3 minutes
	if err = r.cache.SetWithTTL(ctx, &entity.Cache{Key: email, Value: regCode, TTL: 180}); err != nil {
		log.Printf("Fail set token lifetime : %s", err)
		return "", err
	}

	return regCode, nil
}

// ResendCode register otp to user
func (r *registrationServices) ResendCode(ctx context.Context, email string) (string, error) {
	trimmedEmail := strings.Trim(email, "")

	if email == "" {
		return "", errors.New("invalidField: email:Email is empty")
	}

	cooldownKey := fmt.Sprintf("%s:cooldown", trimmedEmail)
	retryKey := fmt.Sprintf("%s:retry", trimmedEmail)

	// Check retry
	retry := r.cache.Get(ctx, retryKey)
	if retry.Val() != "" && retry.Val() == "0" {
		ttl := r.cache.TTL(ctx, retryKey)
		errMsg := fmt.Sprintf("tooMuchRetry: Too much retry, please wait %d minutes", int64(time.Duration(ttl)/time.Minute))
		return "", errors.New(errMsg)
	}

	// Check cooldown timer
	cooldown := r.cache.Get(ctx, cooldownKey)
	if cooldown.Val() != "" {
		ttl := r.cache.TTL(ctx, cooldownKey)
		errMsg := fmt.Sprintf("Wait until %d seconds", int64(time.Duration(ttl)/time.Second))
		return "", errors.New(errMsg)
	}

	var register entity.Registration
	var err error
	if register, err = r.registrationRepo.GetByEmail(ctx, email); err != nil {
		return "", err
	}

	var user entity.User
	if user, err = r.userRepo.GetByEmail(ctx, email); err != nil {
		return "", err
	}

	if register == (entity.Registration{}) {
		return "", errors.New("Please register from the beginning")
	}

	if register.IsRegistered || user != (entity.User{}) {
		return "", errors.New("invalidField: email:Email already registered")
	}

	// Check register status
	if err = r.registerStatus(ctx, email); err != nil {
		return "", err
	}

	// Set cooldown timer
	if err = r.cache.SetWithTTL(ctx, &entity.Cache{Key: cooldownKey, Value: email, TTL: 60}); err != nil {
		log.Printf("Fail set token lifetime : %s", err)
		return "", err
	}

	// Set cooldown timer
	if err = r.cache.SetWithTTL(ctx, &entity.Cache{Key: cooldownKey, Value: email, TTL: 60}); err != nil {
		log.Printf("Fail set token lifetime : %s", err)
		return "", err
	}

	// Set or decrement limiter value for 15 minutes
	if retry.Val() == "" {
		if err = r.cache.SetWithTTL(ctx, &entity.Cache{Key: retryKey, Value: "3", TTL: 15 * 60}); err != nil {
			log.Printf("Fail set request limiter : %s", err)
			return "", err
		}
	} else if retry.Val() > "0" {
		if err = r.cache.Decrement(ctx, retryKey); err != nil {
			log.Printf("Fail increment request limiter : %s", err)
			return "", err
		}
	}

	// Generate register code
	var regCode string
	if regCode, err = r.generateOTPCode(6); err != nil {
		log.Printf("Failed generate otp code : %s", err)
		return "", err
	}

	// Save register  code and data
	// Deactivate all register token by identifier
	if err = r.tokenRepo.DeactivateAllByIdentifierAndType(ctx, email, "register"); err != nil {
		log.Printf("Fail deactive token : %s", err)
		return "", err
	}

	if err = r.tokenRepo.CreateNewToken(ctx, &entity.Token{Identifier: email, Type: "register", Code: regCode}); err != nil {
		log.Printf("Fail save token : %s", err)
		return "", err
	}

	// Set register code to redis for 3 minutes
	if err = r.cache.SetWithTTL(ctx, &entity.Cache{Key: email, Value: regCode, TTL: 180}); err != nil {
		log.Printf("Fail set token lifetime : %s", err)
		return "", err
	}

	return regCode, nil
}

// VerifyCode verify registration code and register user data from temporary
func (r *registrationServices) VerifyCode(ctx context.Context, email string, code string) error {
	var register entity.Registration
	var err error

	// Check user registration data
	if register, err = r.registrationRepo.GetByEmail(ctx, email); err != nil {
		return err
	}

	// Check if register data is null
	if register == (entity.Registration{}) {
		return errors.New("Please register from the beginning")
	}

	if register.IsRegistered {
		return errors.New("Email already registered")
	}

	// Check register status
	if err = r.registerStatus(ctx, email); err != nil {
		return errors.New(fmt.Sprintf("invalidField: code:%s", err))
	}

	// Check token is still active
	verifyCode := r.cache.Get(ctx, email)
	if verifyCode.Val() == "" {
		errMsg := fmt.Sprintf("invalidField: code:Verify code is expired")
		return errors.New(errMsg)
	}

	// Check token is valid
	if verifyCode.Val() != code {
		// Failed retry
		if err = r.setOrDecrementFailedVerifyRetry(ctx, email); err != nil {
			r.setFailedVerifyRetryTimedown(ctx, email)
			return errors.New(fmt.Sprintf("invalidField: code:%s", err))
		}

		errMsg := fmt.Sprintf("invalidField: code:Verify code is invalid")
		return errors.New(errMsg)
	}

	// Create new user from user temporary data
	newUser := entity.User{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  register.Password,
		Phone:     register.Phone,
	}

	if err = r.userRepo.Store(ctx, &newUser); err != nil {
		log.Printf("Failed create a new user : %s", err)
		return err
	}

	// Update register data status
	if err = r.registrationRepo.ChangeStatusToRegistered(ctx, email); err != nil {
		log.Printf("Failed change user status to registered : %s", err)
		return err
	}

	// Deactivate token in database
	if err = r.tokenRepo.DeactivateAllByIdentifierAndType(ctx, email, "register"); err != nil {
		log.Printf("Fail deactive token : %s", err)
		return err
	}

	// Deactivate all user token from in memory db
	if err = r.cache.Delete(ctx, fmt.Sprintf("%s*", email)); err != nil {
		log.Printf("Failed delete all register token : %s", err)
		return err
	}

	return nil
}

// generateOTPCode generate random number for otp code
func (r *registrationServices) generateOTPCode(length int) (string, error) {
	const otpChars = "1234567890"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

// setOrDecrementFailedVerifyRetry set or decrement failed retry
func (r *registrationServices) setOrDecrementFailedVerifyRetry(ctx context.Context, key string) error {
	failedRetryKey := fmt.Sprintf("%s:failedretry", key)
	failedRetry := r.cache.Get(ctx, failedRetryKey)
	if failedRetry.Val() == "" {
		err := r.cache.SetWithTTL(ctx, &entity.Cache{Key: failedRetryKey, Value: "9", TTL: 60 * 2})
		if err != nil {
			return nil
		}
		return nil
	}
	fmt.Println(failedRetry)
	err := r.cache.Decrement(ctx, failedRetryKey)
	if err != nil {
		return err
	}
	if failedRetry.Val() == "0" {
		return errors.New("Already failed 10 times, please wait for 10 minutes")
	}

	return nil
}

// setFailedVerifyRetryTimedown set failed code
func (r *registrationServices) setFailedVerifyRetryTimedown(ctx context.Context, key string) {
	failedRetryTimedownKey := fmt.Sprintf("%s:failedretrytimedown", key)
	err := r.cache.SetWithTTL(ctx, &entity.Cache{Key: failedRetryTimedownKey, Value: key, TTL: 60 * 10})
	if err != nil {
		log.Println("Failed verify retry timedown ", err)
	}
	return
}

// registerStatus check redis error
func (r *registrationServices) registerStatus(ctx context.Context, key string) error {
	failedRetryTimedownKey := fmt.Sprintf("%s:failedretrytimedown", key)
	status := r.cache.Get(ctx, failedRetryTimedownKey)
	if status.Val() != "" {
		return errors.New("Already failed 10 times, please wait for 10 minutes")
	}
	return nil
}
