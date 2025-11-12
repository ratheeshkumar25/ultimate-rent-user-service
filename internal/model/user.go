package model

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Phone     string         `json:"phone" gorm:"uniqueIndex;not null" validate:"required,min=10,max=15"`
	Password  string         `json:"-" gorm:"not null" validate:"required,min=8"`
	FirstName string         `json:"first_name" gorm:"not null" validate:"required,min=2,max=50"`
	LastName  string         `json:"last_name" gorm:"not null" validate:"required,min=2,max=50"`
	Email     string         `json:"email" gorm:"uniqueIndex" validate:"omitempty,email"`
	AadharNo  string         `json:"aadhar_no" gorm:"uniqueIndex" validate:"omitempty,len=12"`
	PancardNo string         `json:"pancard_no" gorm:"uniqueIndex" validate:"omitempty,len=10"`
	Licence   string         `json:"licence" gorm:"uniqueIndex" validate:"omitempty,min=10,max=20"`
	Wallet    float64        `json:"wallet" gorm:"default:0"`
	IsBlocked bool           `json:"is_blocked" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// Relations
	Addresses []Address  `json:"addresses,omitempty" gorm:"foreignKey:UserID"`
	Documents []Document `json:"documents,omitempty" gorm:"foreignKey:UserID"`
}

// TableName returns the table name for User model
func (User) TableName() string {
	return "users"
}

// BeforeCreate hook to set default values
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Wallet == 0 {
		u.Wallet = 0.0
	}
	return nil
}
