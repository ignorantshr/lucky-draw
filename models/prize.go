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

type Prize struct {
	BaseModel
	// Id   int64  `json:"id"`
	// Name string `json:"name"`
	Url string `json:"url"`

	Probability int   `json:"probability" orm:"-"` // 概率
	Number      int64 `json:"number" orm:"-"`      // 数量
}

func (p *Prize) String() string {
	jsonBtyes, _ := json.Marshal(p)
	return fmt.Sprintf("%s", jsonBtyes)
}

func AddPrize(prize *Prize) error {
	return dotx(func(txOrm orm.TxOrmer) error {
		_, err := txOrm.Insert(prize)
		if err != nil {
			log.Println(err)
			return fmt.Errorf(result.ADD_ERROR, prize.Name)
		}
		log.Printf("add prize %v", prize)
		return err
	})
}

func UpdatePrize(prize *Prize) error {
	if !idCheck(&prize.BaseModel) {
		return orm.ErrArgs
	}
	return dotx(func(txOrm orm.TxOrmer) error {
		num, err := txOrm.Update(prize)
		if err == nil && num != 1 {
			return fmt.Errorf(result.UPDATE_ERROR, prize.Id)
		}
		log.Printf("update prize %v", prize)
		return err
	})
}

func DelPrize(id int64) error {
	return dotx(func(txOrm orm.TxOrmer) error {
		prize := &Prize{
			BaseModel: BaseModel{Id: id},
		}
		if err := txOrm.Read(prize); err != nil {
			return fmt.Errorf(result.NOT_EXIST_ERROR)
		}
		num, err := txOrm.Delete(prize)
		if err == nil && num != 1 {
			return fmt.Errorf(result.DEL_ERROR, prize.Id)
		}
		log.Printf("delete prize %v", prize)
		return err
	})
}

func GetPrize(prize *Prize) ([]*Prize, error) {
	var ps []*Prize
	o := orm.NewOrm()
	qs := o.QueryTable(prize)
	if prize.Id != 0 {
		qs = qs.Filter("id", prize.Id)
	}
	if strings.TrimSpace(prize.Name) != "" {
		qs = qs.Filter("name__contains", prize.Name)
	}
	_, err := qs.All(&ps)
	return ps, err
}

func GetAllPrize() ([]*Prize, error) {
	var ps []*Prize
	o := orm.NewOrm()
	_, err := o.QueryTable(Prize{}).All(&ps)
	return ps, err
}
