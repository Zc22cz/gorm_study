package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var mysqlLogger logger.Interface

func init() {
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "gorm"     //数据库名
	timeout := "10s"     //连接超时，10秒

	mysqlLogger = logger.Default.LogMode(logger.Info)

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true,
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "f_",
		//	SingularTable: true,
		//	NoLowerCase:   true,
		//},
		//Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	DB = db
}

//type Student struct {
////	ID    uint    `gorm:"size:10"`
////	Name  string  `gorm:"size:16"`
////	Age   int     `gorm:"size:4"`
////	Email *string `gorm:"size:128"` //指针用来存空值
////	Date  string  `gorm:"defalut:2022-12-08;comment:日期"`
////}

//func main() {
//	//DB = DB.Session(&gorm.Session{
//	//	Logger: mysqlLogger,
//	//})
//	DB.Debug().AutoMigrate(&Student{})
//}
