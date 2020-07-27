package article

// CreateArticleReqStruct create article request body
type CreateArticleReqStruct struct {
	Title        string   `json:"title" binding:"required" validate:"min=2,max=50"`
	Abstract     string   `json:"abstract" binding:"required" validate:"min=1,max=200"`
	Content      string   `json:"content" binding:"required" validate:"min=1"`
	AllowComment uint8    `json:"allowComment" validate:"min=1,max=1"`
	Password     string   `json:"password" validate:"max=10"`
	Thumbnail    string   `json:"thumbnail"`
	Topping      uint8    `json:"topping"`
	CategoryID   uint64   `json:"categoryId" binding:"required"`
	Tags         []uint64 `json:"tags"`
}

// CreateArticleResStruct create article response body
type CreateArticleResStruct struct {
}
