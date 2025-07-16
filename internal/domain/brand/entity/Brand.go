package entity

import (
	"gorm.io/gorm"
	"time"
)

type Brand struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
}

func (Brand) TableName() string {
	return "brands"
}

func (b *Brand) BeforeCreate(tx *gorm.DB) error {
	b.CreatedAt = time.Now()
	return nil
}

func (b *Brand) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
