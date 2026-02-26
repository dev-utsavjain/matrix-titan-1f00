package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID            string    `gorm:"type:uuid;primaryKey" json:"id"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Slug          string    `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Content       string    `gorm:"type:text" json:"content"`
	Excerpt       string    `gorm:"type:varchar(500)" json:"excerpt"`
	Featured      bool      `gorm:"type:boolean;default:false" json:"featured"`
	FeaturedImage string    `gorm:"type:text" json:"featuredImage,omitempty"`
	ReadTime      int       `gorm:"type:integer" json:"readTime"`
	Status        string    `gorm:"type:varchar(20);not null" json:"status"` // draft, published
	Views         int       `gorm:"type:integer;default:0" json:"views"`
	AuthorID      string    `gorm:"type:uuid;not null" json:"authorId"`
	CategoryID    string    `gorm:"type:uuid" json:"categoryId"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	Author   User     `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}

func (Post) TableName() string {
	return "posts"
}
