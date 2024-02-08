package main

//type User struct {
//	ID       uint      `gorm:"size:4"`
//	Name     string    `gorm:"size:8"`
//	Articles []Article // 用户拥有的文章列表
//}
//
//type Article struct {
//	ID     uint   `gorm:"size:4"`
//	Title  string `gorm:"size:16"`
//	UserID uint   `gorm:"size:4"` // 属于   这里的类型要和引用的外键类型一致，包括大小
//	User   User   // 属于
//}

func main() {
	//创建用户，并且创建文章
	//DB.Create(&User{
	//	Name: "杰",
	//	Articles: []Article{
	//		{Title: "golang"},
	//		{Title: "c++"},
	//	},
	//})

	//创建文章，关联已有用户
	//DB.Create(&Article{
	//	Title:  "java",
	//	UserID: 1,
	//})

	//DB.Create(&Article{
	//	Title: "php",
	//	User: User{
	//		Name: "张",
	//	},
	//})

	//给现有用户绑定文章
	//var user User
	//DB.Take(&user, 1)
	//var article Article
	//DB.Take(&article, 5)
	//
	////user.Articles = []Article{article}
	////DB.Save(&user)
	//DB.Model(&user).Association("Articles").Append(&article)

	//给现有文章关联用户
	//var user User
	//DB.Take(&user, 1)
	//
	//var article Article
	//DB.Take(&article, 5)
	//
	//DB.Model(&article).Association("User").Append(&user)
	//
	//DB.AutoMigrate(&User{}, &Article{})

	//预加载
	//var user User
	//DB.Preload("Articles").Take(&user, 1)
	//fmt.Println(user)

	//带条件的预加载
	//var user User
	//DB.Preload("Articles", "id > ?", 3).Take(&user, 1)
	//fmt.Println(user)

	//自定义预加载
	//var user User
	//DB.Preload("Articles", func(db *gorm.DB) *gorm.DB {
	//	return db.Where("id in ?", []int{1, 2})
	//}).Take(&user, 1)
	//fmt.Println(user)

	//级联删除
	//var user User
	//DB.Take(&user, 1)
	//DB.Select("Articles").Delete(&user)

	//清除外键关系
	//var user User
	//DB.Preload("Articles").Take(&user, 2)
	//DB.Model(&user).Association("Articles").Delete(&user.Articles)
	//DB.Delete(&uesr)
}
