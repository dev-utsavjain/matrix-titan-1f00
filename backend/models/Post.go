package models

import (
	"time"
	"github.com/google/uuid"
)

type Post struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title         string    `gorm:"type:varchar(255);not null" json:"title"`
	Slug          string    `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Content       string    `gorm:"type:text" json:"content"`
	Excerpt       string    `gorm:"type:varchar(500)" json:"excerpt"`
	FeaturedImage string    `gorm:"type:text" json:"featuredImage"`
	ReadTime      int       `gorm:"type:int" json:"readTime"`
	Views         int       `gorm:"type:int;default:0" json:"views"`
	Featured      bool      `gorm:"type:boolean;default:false" json:"featured"`
	Status        string    `gorm:"type:varchar(20);check:status IN ('draft','published')" json:"status"`
	AuthorID      uuid.UUID `gorm:"type:uuid;not null" json:"authorId"`
	CategoryID    uuid.UUID `gorm:"type:uuid" json:"categoryId"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Post) TableName() string {
	return "posts"
}
