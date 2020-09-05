package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang-gin-todolist/interfaces"
	"golang-gin-todolist/models"
)

type userRepository struct {

}

func NewUserRepository() interfaces.IUserRepository {
	return &userRepository{}
}

func (repository *userRepository) Create(user models.User) (bool, error){
	if err := models.Db.Create(&user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (repository *userRepository) CheckExistByEmail(email string) (bool, error) {
	var user models.User
	if err := models.Db.Where("email = ?",email).First(&user).Error; err != nil {
		// 判斷是否有資料
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

func (repository *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := models.Db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}