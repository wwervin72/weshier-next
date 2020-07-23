package model

import "time"

// BaseModel base model define
type BaseModel struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;" json:"id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
