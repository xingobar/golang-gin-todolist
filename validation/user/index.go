package user

// user 共用 message
var Message = map[string]string{
	"Username.required": "請輸入用戶姓名",
	"Email.required": "請輸入電子郵件",
	"Password.required": "請輸入密碼",
	"Password.min": "密碼不得小於6位數",
	"Email.email": "電子郵件格式不符",
	"ConfirmPassword.eqfield": "密碼不一致",
}

// 註冊 validation
type RegisterValidation struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"eqfield=Password"`

}

// 登入
type LoginValidation struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


