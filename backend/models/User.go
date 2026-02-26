package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Avatar    string    `gorm:"type:text" json:"avatar,omitempty"`
	Bio       string    `gorm:"type:text" json:"bio,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}
