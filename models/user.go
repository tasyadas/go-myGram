package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age       int    `gorm:"not null;check:age > 8" json:"age" valid:"required~Your age is required,gt(8)~Age must be more than 8"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
