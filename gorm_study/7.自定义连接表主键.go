package main

//
//import (
//	"fmt"
//	"time"
//)
//
//type ArticleModel struct {
//	ID    uint
//	Title string
//	Tags  []TagModel `gorm:"many2many:article_tags;joinForeignKey:ArticleID;JoinReferences:TagID"`
//}
//
//type TagModel struct {
//	ID       uint
//	Name     string
//	Articles []ArticleModel `gorm:"many2many:article_tags;joinForeignKey:TagID;JoinReferences:ArticleID"`
//}
//
//type ArticleTagModel struct {
//	ArticleID uint `gorm:"primaryKey"` // article_id
//	TagID     uint `gorm:"primaryKey"` // tag_id
//	CreatedAt time.Time
//}
//
//func main() {
//	DB.SetupJoinTable(&ArticleModel{}, "Tags", &ArticleTagModel{})
//	DB.SetupJoinTable(&TagModel{}, "Articles", &ArticleTagModel{})
//	//err := DB.AutoMigrate(&ArticleModel{}, &TagModel{}, &ArticleTagModel{})
//	//fmt.Println(err)
//
//	DB.Create(&ArticleModel{
//		Title: "goland基础",
//		Tags: []TagModel{
//			{Name: "后端"},
//			{Name: "goland"},
//		},
//	})
//
//	var article []ArticleModel
//	DB.Preload("Tags").Find(&article, 1)
//	fmt.Println(article)
//}
