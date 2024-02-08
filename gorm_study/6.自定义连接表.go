package main

import (
	"fmt"
	"time"
)

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags"`
}

type Tag struct {
	ID   uint
	Name string
}

type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
	CreatedAt time.Time
	//如果加上下面这样的字段，就写的钩子函数，BeforeCreate在创建时获取用户名并赋值
	//username string
}

func main() {
	// 设置Article的Tags表为ArticleTag
	DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
	// 如果tag要反向应用Article，那么也得加上
	// DB.SetupJoinTable(&Tag{}, "Articles", &ArticleTag{})
	//err := DB.AutoMigrate(&Article{}, &Tag{}, &ArticleTag{})
	//fmt.Println(err)

	//添加文章并添加标签，并自动关联
	//DB.Create(&Article{
	//	Title: "goland基础",
	//	Tags: []Tag{
	//		{Name: "后端"},
	//		{Name: "goland"},
	//	},
	//})
	//添加文章，关联已有标签
	//var tags []Tag
	//DB.Find(&tags, "name in ?", []string{"后端"})
	//DB.Create(&Article{
	//	Title: "c++基础",
	//	Tags:  tags,
	//})
	//给已有文章关联标签
	//var article Article
	//DB.Take(&article, "title = ?", "java基础")
	//var tags []Tag
	//DB.Find(&tags, "name in ? ", []string{"java"})
	//DB.Model(&article).Association("Tags").Append(tags)
	//替换已有文章的标签
	//var article Article
	//DB.Take(&article, "title = ?", "java基础")
	//var tags []Tag
	//DB.Find(&tags, "name in ? ", []string{"后端"})
	//DB.Model(&article).Association("Tags").Replace(tags)
	//查询文章列表，显示标签
	var article1 []Article
	DB.Preload("Tags").Find(&article1)
	fmt.Println(article1)
}
