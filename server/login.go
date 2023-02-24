package server

import (
	"errors"
	"github.com/Akilan1999/remotegameplay/server/auth"
	"gorm.io/gorm"
)

func AuthLogin(db *gorm.DB, user *Users) (string, error) {
	// Generate Hash from password
	password, err := auth.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	match, err := CheckIfEmailAndPasswordMatch(db, user.EmailID, password)
	if err != nil {
		return "", err
	}

	if match == "success" {
		return "success", nil
	}

	return "", errors.New("Something is wrong with from our side. Email us: me@akilan.io to find out more. ")
}
