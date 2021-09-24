package models

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	registerDB()
}

func registerDB() {
	host := "localhost"
	port := 3306
	user := "root"
	passwd := ""

	if host_str := os.Getenv("mysql_host"); host_str != "" {
		host = host_str
	}
	if port_str := os.Getenv("mysql_port"); port_str != "" {
		port, _ = strconv.Atoi(port_str)
	}
	if user_str := os.Getenv("mysql_user"); user_str != "" {
		user = user_str
	}
	if passwd_str := os.Getenv("mysql_passwd"); passwd_str != "" {
		passwd = passwd_str
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/lucky_draw?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, host, port)
	newDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = newDb
}

type BaseModel struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func idCheck(model *BaseModel) bool {
	return model.Id > 0
}

type PoolPrizeQuery struct {
	PoolId    int64  `json:"poolId"`
	PrizeId   int64  `json:"prizeId"`
	PrizeName string `json:"prizeName"`
}
