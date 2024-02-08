package main

import (
	"gorm.io/gorm"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8" json:"name"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func main() {
	//var studentList []Student
	//DB.Find(&studentList).Delete(&studentList)
	//studentList = []Student{
	//	{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
	//	{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
	//	{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
	//	{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
	//	{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
	//	{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
	//	{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
	//	{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
	//	{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	//}
	//DB.Create(&studentList)

	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})

	//#查询用户名是枫枫的
	//select * from students where name = '枫枫';
	//#查询用户名不是枫枫的
	//select * from students where not name = '枫枫';
	//#查询用户名包含 如燕，李元芳的
	//select * from students where name in ('如燕', '李元芳');
	//#查询姓李的
	//select * from students where name like '李%';
	//#查询年龄大于23，是qq邮箱的
	//select * from students where age > 23 and  email like '%@qq.com';
	//#查询是qq邮箱的，或者是女的
	//select * from students where gender = false or  email like '%@qq.com';

	//where函数
	//var users []Student
	//// 查询用户名是枫枫的
	//DB.Where("name = ?", "枫枫").Find(&users)
	//DB.Find(&users, "name = ?", "枫枫")
	//fmt.Println(users)
	//// 查询用户名不是枫枫的
	//DB.Where("name <> ?", "枫枫").Find(&users)
	//DB.Where("not name = ?", "枫枫").Find(&users)
	//DB.Not("name = ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名包含 如燕，李元芳的
	//DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
	//fmt.Println(users)
	//// 查询姓李的
	//DB.Where("name like ?", "李%").Find(&users)
	//fmt.Println(users)
	//// 查询年龄大于23，是qq邮箱的
	//DB.Where("age > ? and email like ?", "23", "%@qq.com").Find(&users)
	//DB.Where("age > ?", "23").Where("email like ?", "%@qq.com").Find(&users)
	//fmt.Println(users)
	//// 查询是qq邮箱的，或者是女的
	//DB.Where("gender = ? or email like ?", false, "%@qq.com").Find(&users)
	//DB.Where("gender = ?", false).Or("email like ?", "%@qq.com").Find(&users)
	//fmt.Println(users)

	////使用结构体查询，自动过滤零值
	//DB.Where(&Student{Name: "李元芳", Age: 0}).Find(&users)
	////使用map查询，不会过滤零值
	//DB.Where(map[string]any{"name": "李元芳", "age": 0}).Find(&users)

	//DB.Select("name", "age").Find(&users)
	//DB.Select([]string{"name","age"}).Find(&users)

	//type User struct {
	//	Name string
	//	Age  int
	//}
	//var students []Student
	//var u []User
	//DB.Select("name", "age").Find(&students).Scan(&u)
	//fmt.Println(u)

	//DB.Order("age desc").Find(&studentList)

	//DB.Limit(2).Offset(0).Find(&studentList)
	//DB.Limit(2).Offset(2).Find(&studentList)
	//DB.Limit(2).Offset(4).Find(&studentList)
	//通用写法
	//limit := 2
	//// 第几页
	//page := 1
	//offset := (page - 1) * limit
	//DB.Limit(limit).Offset(offset).Find(&studentList)
	//fmt.Println(studentList)

	//var ageList []int
	//DB.Table("students").Select("age").Distinct("age").Scan(&ageList)
	//DB.Table("students").Select("distinct age").Scan(&ageList)
	//fmt.Println(ageList)

	//type Group struct {
	//	Count    int
	//	Gender   string
	//	NameList string
	//}

	//var groupList []Group
	////DB.Model(Student{}).Select("count(id) as count", "gender").Group("gender").Scan(&groupList)
	////DB.Model(Student{}).Select("group_concat(name) as name_list", "count(id) as count", "gender", ).Group("gender").Scan(&groupList)
	//DB.Raw("Select group_concat(name) as name_list, count(id) as count, gender from students group by gender").Scan(&groupList)
	//fmt.Println(groupList)

	//DB.Raw("select * from students where age > (select avg(age) from students)").Find(&studentList)
	/*DB.Where("age > (?)", DB.Model(Student{}).Select("avg(age)")).Find(&studentList)
	fmt.Println(studentList)*/

	//var users []Student
	//DB.Scopes(Age23).Find(&users)
	//fmt.Println(users)
}

func Age23(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 23)
}

func PtrString(email string) *string {
	return &email
}
