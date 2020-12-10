package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/afiefafian/todo-api/src/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type userAuthServices struct {
	userRepo entity.UserRepository
}

// NewUserAuthServices create new authentication services
func NewUserAuthServices(u entity.UserRepository) entity.UserAuthServices {
	return &userAuthServices{
		userRepo: u,
	}
}

func (u userAuthServices) Authentication(ctx context.Context, login *entity.UserLogin) (entity.User, string, error) {
	var (
		user  entity.User
		token string
		err   error
	)

	email := strings.Trim(login.Email, "")
	password := login.Password

	if user, err = u.userRepo.GetByEmail(ctx, email); err != nil {
		return user, "", err
	}

	// If user not found
	if user == (entity.User{}) {
		return user, "", errors.New("invalidField: email:Email is not registered")
	}

	// Check user password
	if ok := u.comparePasswords(user.Password, password); !ok {
		return user, "", errors.New("invalidField: password:Wrong email or password")
	}

	// Generate auth token
	identifier := fmt.Sprintf("%s", user.ID)
	if token, err = u.GenerateAuthToken(identifier); err != nil {
		return user, "", errors.New("invalidField: password:Wrong email or password")
	}

	// Hide user password
	user.Password = ""

	return user, token, nil
}

func (u userAuthServices) Logout(ctx context.Context) error {
	panic("implement me")
}

func (u userAuthServices) GenerateAuthToken(identifier string) (string, error) {
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    identifier,
		"exp":        time.Now().Add(time.Minute * 60 * 24 * 1).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	secret := viper.GetString("secret")
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u userAuthServices) comparePasswords(hashedPwd string, password string) bool {
	byteHash := []byte(hashedPwd)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		return false
	}

	return true
}
