package model

// ArticleModel article model
type ArticleModel struct {
	BaseModel
	Title    string `json:"title" gorm:"not null;column:title" binding:"required" validate:"min=2,max=50"`
	Abstract string `json:"abstract" gorm:"column:abstract" binding:"required" validate:"max=200"`
}

// TableName article table name
func (am *ArticleModel) TableName() string {
	return "ws_article"
}
