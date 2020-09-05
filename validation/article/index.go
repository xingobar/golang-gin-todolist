package article

var Message = map[string]string {
	"Title.required": "請輸入標題",
	"Content.required": "請輸入內容",
	"Tags.required": "請傳入標籤",
}

type CreateArticleValidation struct {
	Title string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" binding:"required"`
	Tags []int `json:"tags[]" form:"tags[]" binding:"required"`
}