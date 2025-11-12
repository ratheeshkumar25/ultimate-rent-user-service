package model

import (
	"time"

	"gorm.io/gorm"
)

// DocumentType represents the type of document
type DocumentType string

const (
	DocumentTypeAadhar   DocumentType = "aadhar"
	DocumentTypePancard  DocumentType = "pancard"
	DocumentTypeLicence  DocumentType = "licence"
	DocumentTypePassPort DocumentType = "passport"
)

// Document represents the document model
type Document struct {
	ID           uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       uint           `json:"user_id" gorm:"not null;index" validate:"required"`
	DocumentType DocumentType   `json:"document_type" gorm:"not null;type:varchar(20)" validate:"required,oneof=aadhar pancard licence"`
	FileName     string         `json:"file_name" gorm:"not null" validate:"required,min=1,max=255"`
	FilePath     string         `json:"file_path" gorm:"not null;uniqueIndex" validate:"required,min=1,max=500"`
	FileSize     int64          `json:"file_size" gorm:"not null" validate:"required,min=1"`
	MimeType     string         `json:"mime_type" gorm:"not null" validate:"required,min=1,max=100"`
	IsVerified   bool           `json:"is_verified" gorm:"default:false"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	// Relations
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// TableName returns the table name for Document model
func (Document) TableName() string {
	return "documents"
}
