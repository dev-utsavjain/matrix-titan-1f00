package models

import (
	"time"
	"gorm.io/gorm"
)

type Comment struct {
	ID        string         `gorm:"type:uuid;primaryKey" json:"id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	AuthorID  string         `gorm:"type:uuid;not null" json:"authorId"`
	PostID    string         `gorm:"type:uuid;not null" json:"postId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	
	Author User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Post   Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
}

func (Comment) TableName() string {
	return "comments"
}
