package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug        string    `gorm:"type:varchar(255);unique;not null" json:"slug"`
	Description string    `gorm:"type:text" json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	Posts []Post `gorm:"foreignKey:CategoryID" json:"posts,omitempty"`
}

func (Category) TableName() string {
	return "categories"
}
