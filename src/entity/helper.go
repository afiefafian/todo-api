package entity

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(pwd string) string {
	//salt := viper.GetString(`salt`)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
