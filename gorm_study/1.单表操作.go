package main

import (
	"gorm.io/gorm"
)

//type Student struct {
//	ID     uint   `gorm:"size:3"`
//	Name   string `gorm:"size:8" json:"name"`
//	Age    int    `gorm:"size:3"`
//	Gender bool
//	Email  *string `gorm:"size:32"`
//}
//
//func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
//	email := "test@qq.com"
//	user.Email = &email
//	return nil
//}

func main() {
	//DB.AutoMigrate(&Student{})

	//添加记录
	//email := "2393775920@qq.com"
	//s1 := Student{
	//	Name:   "杰",
	//	Age:    20,
	//	Gender: true,
	//	Email:  &email,
	//}
	//err := DB.Create(&s1).Error
	//fmt.Println(s1)

	//批量插入
	//var studentList []Student
	//
	//for i := 0; i < 5; i++ {
	//	studentList = append(studentList, Student{
	//		Name:   fmt.Sprintf("杰%d", i),
	//		Age:    20 + i,
	//		Gender: true,
	//		Email:  &email,
	//	})
	//}
	//
	//err := DB.Create(&studentList)
	//fmt.Println(err)

	//单条记录的查询
	//var student Student
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})

	//DB.Take(&student)
	//fmt.Println(student)
	//student = Student{}
	//DB.First(&student)
	//fmt.Println(student)
	//student = Student{}
	//DB.Last(&student)
	//fmt.Println(student)
	//student = Student{}

	//根据主键查询
	//err := DB.Take(&student, 2)
	//fmt.Println(student)
	//fmt.Println(err)

	//根据其他条件查询
	//DB.Take(&student, "name = ?", "杰")

	//student.ID = 2
	//count := DB.Take(&student).RowsAffected
	//DB.Take(&student)
	//fmt.Println(student, count)

	//查询多条记录
	//var studentList []Student
	//
	//count := DB.Find(&studentList).RowsAffected
	//fmt.Println(count)
	//for _, student := range studentList {
	//	fmt.Println(student)
	//}
	//data, _ := json.Marshal(studentList)
	//fmt.Println(string(data))

	//DB.Find(&studentList, []int{2, 3, 5})
	//DB.Find(&studentList, "name in (?)", []string{"杰", "杰1"})
	//fmt.Println(studentList)

	//var student Student
	//DB.Take(&student, 3)
	//save更新
	//student.Age = 6
	//student.Name = "杰666"
	//DB.Save(&student)

	//只更新指定字段
	//DB.Select("name").Save(&student)

	//var studentList []Student
	//批量更新单个字段
	//DB.Find(&studentList, []int{2, 3, 6}).Update("gender", false)
	//批量更新多个字段
	//DB.Find(&studentList, []int{2, 3, 6}).Updates(Student{Age: 66, Gender: true})
	//DB.Find(&studentList, []int{2, 3, 6}).Updates(map[string]any{
	//	"name": "杰123",
	//})

	//删除
	//var student Student
	//DB.Delete(&student, 2)
	//DB.Delete(&student, []int{4, 5})
	//
	//DB.Take(&student)
	//DB.Delete(&student)

	//DB.Create(&Student{
	//	Name: "杰杰",
	//	Age:  15,
	//})
}
