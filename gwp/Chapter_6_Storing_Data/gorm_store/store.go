package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        int64 `gorm:"primaryKey"`
	Content   string
	Author    string    `gorm:"not null"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
}

type Comment struct {
	ID        int64 `gorm:"primaryKey"`
	Content   string
	Author    string `gorm:"not null"`
	PostID    int64  `gorm:"index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	dsn := "host=192.168.1.243 user=gwp password=cw2AEiTZBn8QtYP- dbname=gwp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = Db.AutoMigrate(&Post{}, &Comment{})
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Good post!", Author: "Joe"}
	Db.Model(&post).Association("Comments").Append(&comment)

	var readPost Post
	Db.Where("author = ?", "Sau Sheong").First(&readPost)

	var comments []Comment
	Db.Model(&readPost).Association("Comments").Find(&comments)
	if len(comments) > 0 {
		fmt.Println(comments[0])
	} else {
		fmt.Println("No comments found for the post.")
	}
}
