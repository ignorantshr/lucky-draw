package models

import (
	"encoding/json"
	"fmt"
	"log"
	"lucky-draw/result"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Prize struct {
	// 基本属性
	BaseModel
	// Id   int64  `json:"id"`
	// Name string `json:"name"`
	Url string `json:"url"`

	// 附加属性
	Probability int   `json:"probability" gorm:"-"` // 概率
	Number      int64 `json:"number" gorm:"-"`      // 数量
}

func (Prize) TableName() string {
	return "prize"
}

func (p *Prize) String() string {
	jsonBtyes, _ := json.Marshal(p)
	return fmt.Sprintf("%s", jsonBtyes)
}

func AddPrize(prize *Prize) error {
	if err := db.Create(prize).Error; err != nil {
		log.Println(err)
		return fmt.Errorf(result.ADD_ERROR, prize.Name)
	}
	log.Printf("add prize %v", prize)
	return nil
}

func UpdatePrize(prize *Prize) error {
	if !idCheck(&prize.BaseModel) {
		return result.PARAM_INVALID
	}

	if err := db.Save(prize).Error; err != nil {
		log.Println(err)
		return fmt.Errorf(result.ADD_ERROR, prize.Name)
	}
	log.Printf("update prize %v", prize)
	return nil
}

func DelPrize(id int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var prize *Prize
		if err := db.Where(id).Take(&prize).Error; err != nil {
			log.Println(err)
			return fmt.Errorf(result.NOT_EXIST_ERROR)
		}
		if err := db.Delete(&Prize{}, id).Error; err != nil {
			log.Println(err)
			return fmt.Errorf(result.DEL_ERROR, prize.Name)
		}
		// 奖池-奖品 关联表
		if err := db.Where("prize_id = ?", id).Delete(&PrizePoolPrize{}).Error; err != nil {
			log.Println(err)
			return fmt.Errorf(result.DEL_ERROR, id)
		}
		log.Printf("delete prize %v", prize)
		return nil
	})
}

func GetPrize(prize *Prize) ([]*Prize, error) {
	var ps []*Prize
	var err error
	if prize.Id != 0 {
		err = db.Where("id = ?", prize.Id).Find(&ps).Error
	} else {
		err = db.Where("name like ?", "%"+prize.Name+"%").Find(&ps).Error
	}
	if err != nil {
		log.Println(err)
	}
	return ps, err
}

func GetPrizeByName(name string) (*Prize, error) {
	var p Prize
	err := db.Where("name = ?", name).Find(&p).Error
	if err != nil {
		log.Println(err)
	}
	return &p, err
}

func GetAllPrize() ([]*Prize, error) {
	var ps []*Prize
	err := db.Find(&ps).Error
	if err != nil {
		log.Println(err)
	}
	return ps, err
}
