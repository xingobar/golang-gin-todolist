package user_service

import (
	"golang-gin-todolist/models"
)

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Register(user models.User) error{
	if err := models.Db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}