package models

import "time"

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;type:varchar(191)" json:"title"`
	Caption   string `gorm:"not null;type:varchar(191)" json:"caption"`
	PhotoUrl  string `gorm:"not null;type:varchar(191)" json:"photo_url"`
	UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
