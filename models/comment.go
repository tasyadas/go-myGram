package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	PhotoID   string `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo_id"`
	Message   string `gorm:"not null;type:varchar(191)" json:"message" valid:"required~Message cannot be empty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
