package services

import (
	"campground_go/models"
	"campground_go/session"
	"campground_go/utils"
	"errors"
)

const (
	Hash_key = "Hash_key_secret"
)

type UserService struct {
	UserDB  *models.UserDB
	hmac    utils.HMAC
	session *session.Session
}

func NewUserService() *UserService {
	user := models.NewUser()
	hmac := utils.NewHmac(Hash_key)
	s := session.NewSession()
	return &UserService{
		UserDB:  user,
		hmac:    hmac,
		session: s,
	}
}

func (us *UserService) Authenticate(user *models.User) (*models.User, error) {
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
