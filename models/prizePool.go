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

type PrizePool struct {
	BaseModel
	// Id   int64  `json:"id"`
	// Name string `json:"name"`
	prizes []*Prize
}

func (p *PrizePool) String() string {
	jsonBtyes, _ := json.Marshal(p)
	return fmt.Sprintf("%s", jsonBtyes)
}

func AddPrizePool(prizePool *PrizePool) error {
	return dotx(func(txOrm orm.TxOrmer) error {
		num, err := txOrm.Insert(prizePool)
		if err == nil && num != 1 {
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
		log.Printf("delete PrizePool %v", prizePool)
		return err
	})
}

func GetPrizePool(prizePool *PrizePool) ([]*PrizePool, error) {
	var ps []*PrizePool
	o := orm.NewOrm()
	qs := o.QueryTable(prizePool)
	if prizePool.Id != 0 {
		qs = qs.Filter("id", prizePool.Id)
	}
	if strings.TrimSpace(prizePool.Name) != "" {
		qs = qs.Filter("name__contains", prizePool.Name)
	}
	_, err := qs.All(&ps)
	return ps, err
}

func GetAllPrizePool() ([]*PrizePool, error) {
	var ps []*PrizePool
	o := orm.NewOrm()
	_, err := o.QueryTable(PrizePool{}).All(&ps)
	// for _, pool := range ps {
	// 	o.QueryTable()
	// }
	return ps, err
}
