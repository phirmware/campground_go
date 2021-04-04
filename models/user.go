package models

import (
	adapter "campground_go/adapters"
	"campground_go/utils"
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	Hash_key = "Hash_key_secret"
)

type UserDB struct {
	postgres *adapter.Postgres
	hmac     utils.HMAC
}

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

func NewUser() *UserDB {
	postgres := adapter.NewPostgresAdapter()
	hmac := utils.NewHmac(Hash_key)
	// postgres.DestructiveReset(&User{})
	return &UserDB{
		postgres: postgres,
		hmac:     hmac,
	}
}

func (u *UserDB) Create(user *User) error {
	hashedPassword := u.hmac.Hash(user.Password)
	sanitizedUser := &User{
		Username:     user.Username,
		Password:     "",
		PasswordHash: hashedPassword,
	}
	if err := u.postgres.Create(sanitizedUser).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserDB) FindByUsername(username string) (*User, error) {
	var user User
	if err := u.postgres.FindByUsername(username, &user); err != nil {
		return nil, errors.New("Could not find user")
	}
	return &user, nil
}
