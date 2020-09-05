package user_service

import (
	"golang-gin-todolist/interfaces"
	"golang-gin-todolist/models"
	"golang-gin-todolist/repository"
)

type UserService struct {
	userRepository interfaces.IUserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

// 註冊
func (s *UserService) Register(user models.User) bool{
	ok, _ := s.userRepository.Create(user)
	return ok
}

// 檢查資料是否存在
func (s *UserService) CheckExistByEmail(email string) (bool, error) {
	return s.userRepository.CheckExistByEmail(email)
}

// 根據信箱以及密碼取得會員
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

// 取得會員文章
func (s *UserService) GetArticles(userId string) ([]models.Article, error){
	var article []models.Article
	if err := models.Db.Preload("Tags").Where("user_id = ?", userId).Find(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}
