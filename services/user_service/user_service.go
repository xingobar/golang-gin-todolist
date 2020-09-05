package user_service

import (
	"golang-gin-todolist/models"
)

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

// 註冊
func (s *UserService) Register(user models.User) error{
	if err := models.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 檢查資料是否存在
func (s *UserService) CheckExistByEmail(email string) (bool, error) {
	var user models.User
	if err := models.Db.Where("email = ?",email).First(&user).Error; err != nil {
		return false, err
	}

	// 判斷是否有資料
	if user == (models.User{}) {
		return true, nil
	}
	return false, nil
}

// 根據信箱以及密碼取得會員
func (s *UserService) GetUserByEmailAndPassword(email string, password string) (*models.User, error) {
	var user models.User

	if err := models.Db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}