package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Base     `valid:"required"`
	Name     string     `json:"name" valid:"notnull"`
	Email     string     `json:"email" valid:"notnull"`
}

func (User *User) isValid() error {
	_, err := govalidator.ValidateStruct(User)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(ID string, name string, email string) (*User, error) {
	User := User{
		ID: ID,
		Name: name,
		Email: email,
	}

	User.ID = uuid.NewV4().String()
	User.CreatedAt = time.Now()

	err := User.isValid()
	if err != nil {
		return nil, err
	}

	return &User, nil
}