package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Info struct {
	Status string `json:"status"`
	Addr   string `json:"addr"`
	Age    int    `json:"age"`
}

// Scan 从数据库中读取出来
func (i *Info) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	info := Info{}
	err := json.Unmarshal(bytes, &info)
	*i = info
	return err
}

// Value 存入数据库
func (i Info) Value() (driver.Value, error) {
	fmt.Printf("入库前, %#v, %T\n", i, i)
	return json.Marshal(i)
}

type AutoModel struct {
	ID   uint
	Name string
	Info Info `gorm:"type:string"`
}

func main() {
	//DB.AutoMigrate(&AutoModel{})

	//DB.Create(&AutoModel{
	//	Name: "杰",
	//	Info: Info{
	//		Status: "success",
	//		Addr:   "湖北宜昌",
	//		Age:    20,
	//	},
	//})

	var user AutoModel
	DB.Take(&user)
	fmt.Println(user)
}
