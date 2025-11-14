package repo

import (
	"github.com/ratheeshkumar25/ultimate-rent-car/ultimate-rent-user-service/internal/model"
	"gorm.io/gorm"
)

// UserRepoInter defines the interface for user repository operations.
type UserRepoInter interface {
	//User Authentication sitories
	CreateUser(user *model.User) (uint32, error)
	GetUserByEmail(email string) (*model.User, error)
	VerifyUser(userID uint32) error
	UpdateUser(user *model.User) error

	//Document Management sitories
	UploadDocument(userID uint, docType model.DocumentType, docPath string) error
	GetDocuments(userID uint) ([]model.Document, error)
	DeleteDocument(docID uint) error

	//Address Management sitories
	AddAddress(address *model.Address) error
	UpdateAddress(address *model.Address) error
	DeleteAddress(addressID uint, userID uint) error
	ListAddresses(userID uint) ([]model.Address, error)

	//Wallet Management sitories
	GetWalletBalance(userID uint) (float64, error)
	UpdateWallet(userID uint, amount float64) error
	GetAllUsers() ([]model.User, error)
}

type UserRepository struct {
	//database connection
	DB *gorm.DB
}

// constructor function to create a new UserRepository
func NewUserRepository(db *gorm.DB) UserRepoInter {
	//return the UserRepository instance
	return &UserRepository{
		DB: db,
	}
}
