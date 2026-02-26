package models

import (
	"time"
	"gorm.io/gorm"
)

type Post struct {
	ID            string         `gorm:"type:uuid;primaryKey" json:"id"`
	Title         string         `gorm:"type:varchar(255);not null" json:"title"`
	Slug          string         `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Excerpt       string         `gorm:"type:varchar(500)" json:"excerpt"`
	Content       string         `gorm:"type:text;not null" json:"content"`
	FeaturedImage string         `gorm:"type:text" json:"featuredImage"`
	ReadTime      int            `gorm:"type:integer;default:0" json:"readTime"`
	Views         int            `gorm:"type:integer;default:0" json:"views"`
	Featured      bool           `gorm:"type:boolean;default:false" json:"featured"`
	Status        string         `gorm:"type:varchar(50);not null" json:"status"`
	AuthorID      string         `gorm:"type:uuid;not null" json:"authorId"`
	CategoryID    string         `gorm:"type:uuid" json:"categoryId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	
	Author   User     `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
	Category Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}

func (Post) TableName() string {
	return "posts"
}
