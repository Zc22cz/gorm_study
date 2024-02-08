package main

//
//import (
//	"gorm.io/gorm"
//	"time"
//)
//
//type Tag struct {
//	ID   uint
//	Name string
//	//Articles []Article `gorm:"many2many:article_tags;"` // 用于反向引用
//}
//
//type Article struct {
//	ID    uint
//	Title string
//	Tags  []Tag `gorm:"many2many:article_tags;"`
//}
//
//type ArticleTag struct {
//	ArticleID uint `gorm:"primaryKey"`
//	TagID     uint `gorm:"primaryKey"`
//	CreatedAt time.Time
//}
//
//func (at *ArticleTag) BeforeCreate(db *gorm.DB) error {
//	at.CreatedAt = time.Now()
//	return nil
//}
//
//func main() {
//	//// 设置Article的Tags表为ArticleTag
//	DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
//	//// 如果tag要反向应用Article，那么也得加上
//	//// DB.SetupJoinTable(&Tag{}, "Articles", &ArticleTag{})
//	//DB.AutoMigrate(&Tag{}, &Article{}, &ArticleTag{})
//
//	//自定义连接表
//
//	//添加文章，关联已有tag
//	//var tags []Tag
//	//DB.Find(&tags, []int{1, 2})
//	//DB.Create(&Article{
//	//	Title: "golang基础",
//	//	Tags:  tags,
//	//})
//
//	//给已有文章替换tag
//	//var article Article
//	//DB.Preload("Tags").Take(&article, 3)
//	//
//	//var tags []Tag
//	//DB.Find(&tags, []int{3})
//	//DB.Model(&article).Association("Tags").Replace(&tags)
//
//	//创建一篇文章和新的tag
//	//DB.Create(&Article{
//	//	Title: "golang基础",
//	//	Tags: []Tag{
//	//		{Name: "golang"},
//	//		{Name: "后端"},
//	//	},
//	//})
//
//	//创建一篇文章，使用已有的tag
//	//var tags []Tag
//	//DB.Find(&tags, "name in ?", []string{"后端"})
//	//DB.Create(&Article{
//	//	Title: "c++基础",
//	//	Tags:  tags,
//	//})
//
//	//创建一篇文章，既使用已有的tag，也使用新的tag
//	//var ExistTags []Tag
//	//DB.Find(&ExistTags, "name in ?", []string{"后端"})
//	//DB.Create(&Article{
//	//	Title: "java后端",
//	//	Tags: append([]Tag{
//	//		{Name: "java"},
//	//	}, ExistTags...),
//	//})
//
//	//查询
//	//var article Article
//	//DB.Preload("Tags").Take(&article, 1)
//	//fmt.Println(article)
//
//	//更新 两种方法
//	//先删除原有的
//	//var article Article
//	//DB.Preload("Tags").Take(&article, 1)
//	//DB.Model(&article).Association("Tags").Delete(article.Tags)
//	//再添加新的
//	//var tag Tag
//	//DB.Take(&tag, 2)
//	//DB.Model(&article).Association("Tags").Append(&tag)
//	//或者使用replace
//	//var article Article
//	//DB.Preload("Tags").Take(&article, 1)
//	//var tag Tag
//	//DB.Take(&tag, 1)
//	//DB.Model(&article).Association("Tags").Replace(&tag)
//}
