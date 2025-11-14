package repo

import "github.com/ratheeshkumar25/ultimate-rent-car/ultimate-rent-user-service/internal/model"

// UploadDocument implements logic to upload a document for a user.
func (u *UserRepository) UploadDocument(userID uint, docType model.DocumentType, docPath string) error {
	// Create a new document record in the database
	if err := u.DB.Model(&model.Document{}).Create(&model.Document{
		UserID:       userID,
		DocumentType: docType,
		FilePath:     docPath,
	}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteDocument the uploaded document by its ID.
func (u *UserRepository) DeleteDocument(docID uint) error {
	// Delete the document record from the database
	if err := u.DB.Model(&model.Document{}).Where("document_id = ?", docID).Delete(&model.Document{}).Error; err != nil {
		return err
	}
	return nil
}

// GetDocuments all documents for a specific user.
func (u *UserRepository) GetDocuments(userID uint) ([]model.Document, error) {
	var documents []model.Document
	// Retrieve all documents associated with the userID
	if err := u.DB.Find(&documents, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return documents, nil

}
