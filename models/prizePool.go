package models

import (
	"encoding/json"
	"fmt"
	"log"
	"lucky-draw/result"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type DrawType int

const (
	ProbabilityType DrawType = iota + 1
	NumberType
)

type PrizePool struct {
	BaseModel
	// Id   int64  `json:"id"`
	// Name string `json:"name"`
	Type   DrawType `json:"type"`
	Prizes []*Prize `json:"prizes" orm:"-"`
}

func (p *PrizePool) String() string {
	jsonBtyes, _ := json.Marshal(p)
	return fmt.Sprintf("%s", jsonBtyes)
}

var poolOrm orm.Ormer
var poolQs orm.QuerySeter

func init() {
	poolOrm = orm.NewOrm()
	poolQs = poolOrm.QueryTable(new(PrizePool))
}

func AddPrizePool(prizePool *PrizePool) error {
	return dotx(func(txOrm orm.TxOrmer) error {
		normalPool(prizePool)
		_, err := txOrm.Insert(prizePool)
		if err != nil {
			log.Println(err)
			return fmt.Errorf(result.ADD_ERROR, prizePool.Name)
		}
		log.Printf("add PrizePool %v", prizePool)
		return err
	})
}

func UpdatePrizePool(prizePool *PrizePool) error {
	if !idCheck(&prizePool.BaseModel) {
		return orm.ErrArgs
	}
	return dotx(func(txOrm orm.TxOrmer) error {
		num, err := txOrm.Update(prizePool)
		if err == nil && num != 1 {
			return fmt.Errorf(result.UPDATE_ERROR, prizePool.Name)
		}
		log.Printf("update PrizePool %v", prizePool)
		return err
	})
}

func DelPrizePool(id int64) error {
	return dotx(func(txOrm orm.TxOrmer) error {
		prizePool := &PrizePool{
			BaseModel: BaseModel{Id: id},
		}
		if err := txOrm.Read(prizePool); err != nil {
			return fmt.Errorf(result.NOT_EXIST_ERROR)
		}
		num, err := txOrm.Delete(prizePool)
		if err == nil && num != 1 {
			return fmt.Errorf(result.DEL_ERROR, prizePool.Name)
		}

		PrizePoolPrize := &PrizePoolPrize{PrizePoolId: id}
		if _, err = txOrm.Delete(PrizePoolPrize); err != nil {
			return fmt.Errorf(result.DEL_ERROR, prizePool.Name)
		}

		log.Printf("delete PrizePool %v", prizePool)
		return err
	})
}

// 根据条件查询，比如 name、id
func GetPrizePool(prizePool *PrizePool) ([]*PrizePool, error) {
	var ps []*PrizePool
	var qs orm.QuerySeter
	if prizePool.Id != 0 {
		qs = poolQs.Filter("id", prizePool.Id)
	}
	if strings.TrimSpace(prizePool.Name) != "" {
		qs = poolQs.Filter("name__contains", prizePool.Name)
	}
	_, err := qs.All(&ps)
	return ps, err
}

func GetAllPrizePool() ([]*PrizePool, error) {
	var ps []*PrizePool
	_, err := poolOrm.QueryTable(PrizePool{}).All(&ps)
	// for _, pool := range ps {
	// 	poolOrm.QueryTable()
	// }
	return ps, err
}

// 获取奖池的详细信息，包括具体的奖品内容、概率、数量
func InfoPrizePool(id int64) (*PrizePool, error) {
	prizePool := new(PrizePool)
	mapper := make([]*PrizePoolPrize, 0)
	prizes := make([]*Prize, 0)

	var err error

	if err = poolQs.Filter("id", id).One(prizePool); err != nil {
		return nil, err
	}

	if _, err = mapperQs.Filter("PrizePoolId", id).All(&mapper); err != nil {
		return nil, err
	}

	prizeQs := poolOrm.QueryTable(Prize{})
	prizeInfos := make(map[int64]PrizePoolPrize, len(mapper))
	prizeIds := make([]int64, len(mapper))
	for i, v := range mapper {
		prizeInfos[v.PrizeId] = *v
		prizeIds[i] = v.PrizeId
	}

	var size int64
	if size, err = prizeQs.Filter("id__in", prizeIds).All(&prizes); err != nil {
		return nil, err
	}
	if size != int64(len(mapper)) {
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
