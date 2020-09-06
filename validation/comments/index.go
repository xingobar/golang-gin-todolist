package comments

var Message = map[string]string {
	"Content.required": "請輸入留言內容",
	"ParentId.required": "請傳入留言編號",
	"ArticleId.required": "請輸入文章編號",
}

type CreateCommentValidation struct {
	Content string `json:"content" form:"content" binding:"required"`
	ParentId uint `json:"parent_id" form:"parent_id"`
	ArticleId uint `json:"article_id" form:"article_id" binding:"required"`
}