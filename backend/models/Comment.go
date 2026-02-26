package models

import (
	"time"
	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	AuthorID  uuid.UUID `gorm:"type:uuid;not null" json:"authorId"`
	PostID    uuid.UUID `gorm:"type:uuid;not null" json:"postId"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}

func (Comment) TableName() string {
	return "comments"
}
