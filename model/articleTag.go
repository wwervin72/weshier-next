package model

// ArticleTagModel article_tag model
type ArticleTagModel struct {
	BaseModel
	TagID     uint64 `json:"tagId" gorm:"no null;column:tag_id"`
	ArticleID uint64 `json:"articleId" gorm:"no null;column:article_id"`
}

// TableName article_tag table name
func (atm *ArticleTagModel) TableName() string {
	return "ws_article_tag"
}
