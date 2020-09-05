package tag

type CreateTagValidation struct {
	Title string `form:"title" json:"title" binding:"required"`
}

var Message = map[string]string {
	"Title.required": "請輸入標籤名稱",
}