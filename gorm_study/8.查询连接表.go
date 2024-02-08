package main

import (
	"fmt"
	"time"
)

type UserModel struct {
	ID       uint
	Name     string
	Collects []ArticleModel `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
}

type ArticleModel struct {
	ID    uint
	Title string
	// 这里也可以反向引用，根据文章查哪些用户收藏了
}

// UserCollectModel 用户收藏文章表
type UserCollectModel struct {
	UserID       uint         `gorm:"primaryKey"` // article_id
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"` // tag_id
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}

func main() {
	//DB.SetupJoinTable(&UserModel{}, "Collects", &UserCollectModel{})
	//err := DB.AutoMigrate(&UserModel{}, &ArticleModel{}, &UserCollectModel{})
	//fmt.Println(err)

	//var user UserModel
	//DB.Preload("Collects").Take(&user, "name = ?", "杰")
	//var collects []ArticleModel
	//DB.Model(&user).Limit(1).Association("Collects").Find(&collects)
	//fmt.Println(collects)

	var collects []UserCollectModel
	DB.Preload("UserModel").Preload("ArticleModel").Find(&collects, "user_id = ?", 1)
	fmt.Println(collects)
}
