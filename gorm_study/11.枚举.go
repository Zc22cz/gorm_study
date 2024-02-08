package main

import (
	"encoding/json"
	"fmt"
)

type Status int

func (s Status) MarshalJSON() ([]byte, error) {
	var str string
	switch s {
	case Running:
		str = "Running"
	case Except:
		str = "Except"
	case OffLine:
		str = "Status"
	}
	return json.Marshal(str)
}

const (
	Running Status = 1
	OffLine Status = 2
	Except  Status = 3
)

type Host struct {
	ID     uint
	Status Status `grom:"size:8" json:"status"`
	IP     string `json:"ip"`
}

func main() {
	//DB.AutoMigrate(&Host{})

	//DB.Create(&Host{
	//	IP:     "127.0.0.1",
	//	Status: Running,
	//})

	var host Host
	DB.Take(&host)
	//fmt.Println(host)
	data, _ := json.Marshal(host)
	fmt.Println(string(data))
}
