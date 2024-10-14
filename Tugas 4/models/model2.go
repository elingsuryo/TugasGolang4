package models

import "github.com/jinzhu/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
    Role     string
}

type Todo struct {
    gorm.Model
    Task   string
    Status string
    UserID uint
}
