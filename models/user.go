package models

import (
	"fmt"
	"go-myGram/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age       int    `gorm:"not null;" json:"age" valid:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Validate() (bool, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return false, err
	}

	if u.Age <= 8 {
		return false, fmt.Errorf("Age must be greater than 8")
	}

	return true, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	valid, err := u.Validate()
	if !valid {
		return err
	}

	u.Password = helpers.HashPass(u.Password)
	return nil
}
