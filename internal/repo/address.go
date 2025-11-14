package repo

import "github.com/ratheeshkumar25/ultimate-rent-car/ultimate-rent-user-service/internal/model"

// Add Address Details of a user.
func (u *UserRepository) AddAddress(address *model.Address) error {
	// Create a new address record in the database
	if err := u.DB.Create(&address).Error; err != nil {
		return err
	}
	return nil

}

// DeleteAddress of users by addressID and userID.
func (u *UserRepository) DeleteAddress(addressID uint, userID uint) error {
	// Delete the address record from the database
	if err := u.DB.Where("address_id = ? AND user_id = ?", addressID, userID).Delete(&model.Address{}).Error; err != nil {
		return err
	}
	return nil
}

// Fetch List of Addresses for a user.
func (u *UserRepository) ListAddresses(userID uint) ([]model.Address, error) {
	// Retrieve all addresses associated with the userID
	var address []model.Address
	if err := u.DB.Where("user_id = ?", userID).Find(&address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

// UpdateAddress implements UserRepoInter.
func (u *UserRepository) UpdateAddress(address *model.Address) error {
	// Update the address record in the database
	if err := u.DB.Save(&address).Error; err != nil {
		return err
	}
	return nil
}
