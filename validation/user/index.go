package user

import "github.com/go-playground/validator"

type RegisterValidation struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (r *RegisterValidation) GetError(err validator.ValidationErrors) string {

	for _, item := range err {
		if item.Field() == "Username" {
			switch item.Tag() {
				case "required":
					return "請輸入用戶姓名"
			}
		} else if item.Field() == "Email" {
			switch item.Tag() {
				case "required":
					return "請輸入電子郵件"
			}
		} else if item.Field() == "Password" {
			switch item.Tag() {
				case "required":
					return "請輸入密碼"
			}
		}
	}
	return "參數錯誤"
}