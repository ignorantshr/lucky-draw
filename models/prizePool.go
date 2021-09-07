package models

import (
	"encoding/json"
	"fmt"
	"log"
	"lucky-draw/result"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type DrawType int

// 抽奖类型
const (
	ProbabilityType DrawType = iota + 1
	NumberType
)

type PrizePool struct {
	BaseModel
	// Id   int64  `json:"id"`
	// Name string `json:"name"`
	Type   DrawType `json:"type"`
	Prizes []*Prize `json:"prizes" gorm:"-"`
}

func (PrizePool) TableName() string {
	return "prize_pool"
}

func (p *PrizePool) String() string {
	jsonBtyes, _ := json.Marshal(p)
	return fmt.Sprintf("%s", jsonBtyes)
}

func AddPrizePool(prizePool *PrizePool) error {
	normalPool(prizePool)
	if err := db.Create(prizePool).Error; err != nil {
		log.Println(err)
		return fmt.Errorf(result.ADD_ERROR, prizePool.Name)
	}
	log.Printf("add PrizePool %v", prizePool)
	return nil
}

func UpdatePrizePool(prizePool *PrizePool) error {
	if !idCheck(&prizePool.BaseModel) {
		return result.PARAM_INVALID
	}
	if err := db.Save(prizePool).Error; err != nil {
		log.Println(err)
		return fmt.Errorf(result.UPDATE_ERROR, prizePool.Name)
	}
	log.Printf("update PrizePool %v", prizePool)
	return nil
}

func DelPrizePool(id int64) error {
	return db.Transaction(func(tx *gorm.DB) error {
		var prizePool *PrizePool
		if err := tx.Where(id).Take(&prizePool).Error; err != nil {
			log.Println(err)
			return fmt.Errorf(result.NOT_EXIST_ERROR)
		}
		if err := db.Delete(&PrizePool{}, id).Error; err != nil {
			log.Println(err)
			return fmt.Errorf(result.DEL_ERROR, prizePool.Name)
		}

		// 奖池-奖品 关联表
		poolPrize := &PrizePoolPrize{
			PrizePoolId: id,
		}
		if err := db.Delete(poolPrize).Error; err != nil {
			log.Println(err)
			return fmt.Errorf(result.DEL_ERROR, id)
		}

		log.Printf("delete PrizePool %v", prizePool)
		return nil
	})
}

// 根据名称查询
func GetPrizePool(prizePool *PrizePool) ([]*PrizePool, error) {
	var ps []*PrizePool
	err := db.Where("name like ?", "%"+prizePool.Name+"%").Find(&ps).Error
	if err != nil {
		log.Println(err)
	}
	return ps, err
}

func GetAllPrizePool() ([]*PrizePool, error) {
	var ps []*PrizePool
	err := db.Find(&ps).Error
	if err != nil {
		log.Println(err)
	}
	return ps, err
}

// 获取奖池的详细信息，包括具体的奖品内容、概率、数量
func InfoPrizePool(id int64) (*PrizePool, error) {
	var prizePool *PrizePool
	var mapper []*PrizePoolPrize
	var prizes []*Prize

	var err error

	if err = db.Take(&prizePool, id).Error; err != nil {
		return nil, err
	}

	queryMapper := &PrizePoolPrize{PrizePoolId: id}
	if err = db.Where(queryMapper).Find(&mapper).Error; err != nil {
		return nil, err
	}

	prizeInfos := make(map[int64]PrizePoolPrize, len(mapper))
	prizeIds := make([]int64, len(mapper))
	for i, v := range mapper {
		prizeInfos[v.PrizeId] = *v
		prizeIds[i] = v.PrizeId
	}

	// 当 slice 为空时会查询全部记录，所以最好不要使用简写方式
	tempRe := db.Where("id in ?", prizeIds).Find(&prizes)
	if err = tempRe.Error; err != nil {
		return nil, err
	}
	if tempRe.RowsAffected != int64(len(mapper)) {
		return nil, fmt.Errorf("奖池中缺少对应的奖品")
	}

	for _, v := range prizes {
		v.Probability = prizeInfos[v.Id].PrizeProbability
		v.Number = prizeInfos[v.Id].PrizeNumber
	}

	prizePool.Prizes = prizes
	return prizePool, nil
}

func normalPool(prizePool *PrizePool) {
	if prizePool.Type == 0 {
		prizePool.Type = 1
	}
}
