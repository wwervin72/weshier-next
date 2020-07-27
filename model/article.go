package model

// ArticleModel article model
type ArticleModel struct {
	BaseModel
	Title        string `zh:"文章标题" json:"title" gorm:"not null;column:title"`
	Abstract     string `zh:"文章摘要" json:"abstract" gorm:"column:abstract"`
	Content      string `zh:"文章内容" json:"content" gorm:"column:content"`
	AllowComment uint8  `zh:"允许评论" json:"allowComment" gorm:"column:allow_comment"`
	Password     string `zh:"允许评论" json:"password" gorm:"column:password"`
	Thumbnail    string `zh:"缩略图" json:"thumbnail" gorm:"column:thumbnail"`
	Topping      uint8  `zh:"置顶" json:"topping" gorm:"column:topping"`
	AuthorID     uint64 `zh:"作者" json:"authorId" gorm:"column:author_id"`
	CategoryID   uint64 `zh:"分类" json:"categoryId" gorm:"column:category_id"`
}

// TableName article table name
func (am *ArticleModel) TableName() string {
	return "ws_article"
}

// Create create article
func (am *ArticleModel) Create() error {
	return DB.Self.Create(&am).Error
}
