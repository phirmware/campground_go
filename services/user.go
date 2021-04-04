package services

import (
	"campground_go/models"
	"campground_go/utils"
	"errors"
)

const (
	Hash_key = "Hash_key_secret"
)

type UserSerice struct {
	UserDB *models.UserDB
	hmac   utils.HMAC
}

func NewUserService() *UserSerice {
	user := models.NewUser()
	hmac := utils.NewHmac(Hash_key)
	return &UserSerice{
		UserDB: user,
		hmac:   hmac,
	}
}

func (us *UserSerice) Authenticate(user *models.User) (*models.User, error) {
	foundUser, err := us.UserDB.FindByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	hashedPassword := us.hmac.Hash(user.Password)
	if foundUser.PasswordHash != hashedPassword {
		return nil, errors.New("Password is incorrect")
	}

	return foundUser, nil
}
