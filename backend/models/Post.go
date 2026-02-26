package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title         string         `gorm:"type:varchar(255);not null" json:"title"`
	Slug          string         `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Content       string         `gorm:"type:text;not null" json:"content"`
	Excerpt       string         `gorm:"type:varchar(500)" json:"excerpt"`
	FeaturedImage string         `gorm:"type:text" json:"featuredImage"`
	ReadTime      int            `gorm:"type:integer" json:"readTime"`
	Views         int            `gorm:"type:integer;default:0" json:"views"`
	Featured      bool           `gorm:"type:boolean;default:false" json:"featured"`
	Status        string         `gorm:"type:varchar(20);not null" json:"status"`
	AuthorID      uuid.UUID      `gorm:"type:uuid;not null" json:"authorId"`
	CategoryID    uuid.UUID      `gorm:"type:uuid" json:"categoryId"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	Author   User     `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

func (Post) TableName() string {
	return "posts"
}
