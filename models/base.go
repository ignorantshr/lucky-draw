package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	registerDB()
}

func registerDB() {
	dsn := "root:@tcp(localhost:3306)/lucky_draw?charset=utf8mb4&parseTime=True&loc=Local"
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
