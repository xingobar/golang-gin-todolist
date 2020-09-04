package user

// user 共用 message
var Message = map[string]string{
	"Username.required": "請輸入用戶姓名",
	"Email.required": "請輸入電子郵件",
	"Password.required": "請輸入密碼",
}

// 註冊 validation
type RegisterValidation struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}


