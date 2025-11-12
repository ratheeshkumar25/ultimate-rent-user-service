package model

import (
	"time"

	"gorm.io/gorm"
)

// Address represents the address model
type Address struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id" gorm:"not null;index" validate:"required"`
	House     string         `json:"house" gorm:"not null" validate:"required,min=1,max=100"`
	Street    string         `json:"street" gorm:"not null" validate:"required,min=1,max=100"`
	City      string         `json:"city" gorm:"not null" validate:"required,min=2,max=50"`
	Zip       uint           `json:"zip" gorm:"not null" validate:"required,min=100000,max=999999"`
	State     string         `json:"state" gorm:"not null" validate:"required,min=2,max=50"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName returns the table name for Address model
func (Address) TableName() string {
	return "addresses"
}
