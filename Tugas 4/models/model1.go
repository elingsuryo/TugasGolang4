package models

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := "username:password@tcp(127.0.0.1:3306)/dbname=Logincharset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open("mysql", dsn)
    if err != nil {
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&User{}, &Todo{})
    DB = db
}
