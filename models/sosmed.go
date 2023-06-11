package models

import "time"

type Sosmed struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)" json:"name"`
	SosmedUrl string `gorm:"not null;type:varchar(191)" json:"sosmed_url"`
	UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
