package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Array []string

// Scan 从数据库中读取出来
func (i *Array) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	info := Array{}
	err := json.Unmarshal(bytes, &info)
	*i = info
	return err
}

// Value 存入数据库
func (i Array) Value() (driver.Value, error) {
	return json.Marshal(i)
}

type HostModel struct {
	ID    uint
	IP    string
	Ports Array `gorm:"type:string"`
}

func main() {
	//DB.AutoMigrate(&HostModel{})
	
	//DB.Create(&HostModel{
	//	IP:    "192.168.21.129",
	//	Ports: []string{"8080", "66"},
	//})

	//var host HostModel
	//DB.Take(&host)
	//fmt.Println(host)
}
