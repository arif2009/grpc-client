package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name string `json:"name"`
}

type Users []*User

// Find user by its ID
func GetUserByID(id uint) (*User, error) {
	user := &User{}
	if tx := DB.Where("id", id).First(user); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, tx.Error
		}
	}

	return user, nil
}

// Get a list of all users
func GetUsers() (Users, error) {
	list := Users{}
	if tx := DB.Find(&list); tx.Error != nil {
		return nil, tx.Error
	}

	return list, nil
}
