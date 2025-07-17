package entity

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	Price     int
	Qty       int
	BrandId   int
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.CreatedAt = time.Now()
	return nil
}

func (p *Product) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}
