package repo

import "github.com/ratheeshkumar25/ultimate-rent-car/ultimate-rent-user-service/internal/model"

// CreateUser implements createUser and returns error if any
func (u *UserRepository) CreateUser(user *model.User) (uint32, error) {
	//create user in the database, if any return error otherwise return user ID
	if err := u.DB.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// VerifyUser implements verify the user.
func (u *UserRepository) VerifyUser(userID uint32) error {
	//Declare a user variable
	var user model.User
	//Find the user by userID
	if err := u.DB.Model(&model.User{}).Where("user_id = ?", userID).First(&user).Error; err != nil {
		return err
	}
	//Update the IsVerified field to true
	user.IsVerified = true
	return u.DB.Save(&user).Error
}

// GetUserByEmail implements UserRepoInter.
func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.DB.Model(&model.User{}).Where("email = ?", &email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers implements UserRepoInter.
func (u *UserRepository) GetAllUsers() ([]model.User, error) {
	var user []model.User
	if err := u.DB.Model(&model.User{}).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil

}

// UpdateUser implements UserRepoInter.
func (u *UserRepository) UpdateUser(user *model.User) error {
	if err := u.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateWallet implements UserRepoInter.
func (u *UserRepository) UpdateWallet(userID uint, amount float64) error {
	if err := u.DB.Model(&model.User{}).Where("user_id = ?", userID).Update("wallet_balance", amount).Error; err != nil {
		return err
	}
	return nil
}

// GetWalletBalance implements UserRepoInter.
func (u *UserRepository) GetWalletBalance(userID uint) (float64, error) {
	var user model.User
	if err := u.DB.Model(&model.User{}).Where("user_id = ?", userID).First(&user).Error; err != nil {
		return 0, err
	}
	return user.Wallet, nil
}
