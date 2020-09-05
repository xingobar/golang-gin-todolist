package interfaces

import "golang-gin-todolist/models"

type IUserRepository interface {
	// 新增會員
	Create(user models.User) (bool, error)

	// 檢查email是否存在
	CheckExistByEmail(email string) (bool, error)

	// 根據電子郵件取得會員
	GetUserByEmail(email string) (*models.User, error)
}