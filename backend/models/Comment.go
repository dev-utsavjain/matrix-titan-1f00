package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	PostID    uuid.UUID      `gorm:"type:uuid;not null" json:"postId"`
	AuthorID  uuid.UUID      `gorm:"type:uuid;not null" json:"authorId"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Post   Post `gorm:"foreignKey:PostID" json:"post,omitempty"`
	Author User `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
}

func (Comment) TableName() string {
	return "comments"
}
