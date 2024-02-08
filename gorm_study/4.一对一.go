package main

//type User struct {
//	ID       uint
//	Name     string
//	Age      int
//	Gender   bool
//	UserInfo UserInfo // 通过UserInfo可以拿到用户详情信息
//}
//
//type UserInfo struct {
//	UserID uint // 外键
//	ID     uint
//	Addr   string
//	Like   string
//}

func main() {
	//DB.AutoMigrate(&User{}, &UserInfo{})
	//DB.Create(&User{
	//	Name:   "杰",
	//	Age:    20,
	//	Gender: true,
	//	UserInfo: UserInfo{
	//		Addr: "湖北",
	//		Like: "cs",
	//	},
	//})

	//var user User
	//DB.Preload("UserInfo").Take(&user)
	//fmt.Println(user)

	//var userInfo UserInfo
	//DB.Preload("User").Take(&userInfo)
	//fmt.Println(userInfo)

	//var user User
	//DB.Take(&user)
	//DB.Select("UserInfo").Delete(&user, 1)
}
