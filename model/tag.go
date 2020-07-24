package model

// TagModel tag model
type TagModel struct {
	BaseModel
	Name   string `json:"name" gorm:"no null" binding:"required" validate:"min=3,max=10"`
	Desc   string `json:"desc" validate:"max=100"`
	UserID uint64 `json:"userId" gorm:"no null;column:user_id"`
}

// TableName tag table name
func (tm *TagModel) TableName() string {
	return "ws_tag"
}
