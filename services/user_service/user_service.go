package user_service

import "golang-gin-todolist/models"

type userService struct {

}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) Register(user models.User) {

}