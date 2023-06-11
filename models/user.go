package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"not null;unique;type:varchar(191)" json:"username"`
	Email     string `gorm:"not null;unique;type:varchar(191)" json:"email"`
	Password  string `gorm:"not null;type:varchar(191)" json:"password"`
	Age       int    `gorm:"not null;type:varchar(191)" json:"age"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
