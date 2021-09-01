package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Prize struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (p *Prize) String() string {
	jsonBtyes, _ := json.Marshal(p)
	return fmt.Sprintf("%s", jsonBtyes)
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:Lenovo123-@tcp(10.221.5.7:3306)/lucky_draw?charset=utf8&loc=Local")

	// register model
	orm.RegisterModel(new(Prize))

	// create table
	// orm.RunSyncdb("default", false, true)
}

func dotx(f func(myTxOrm orm.TxOrmer) error) error {
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		return f(txOrm)
	})
	return err
}

func Add(prize *Prize) error {
	return dotx(func(txOrm orm.TxOrmer) error {
		_, err := txOrm.Insert(prize)
		log.Printf("add prize %v", prize)
		return err
	})
}

func Update(prize *Prize) error {
	if !argCheck(prize) {
		return orm.ErrArgs
	}
	return dotx(func(txOrm orm.TxOrmer) error {
		_, err := txOrm.Update(prize)
		log.Printf("update prize %v", prize)
		return err
	})
}

func Del(id int64) error {
	return dotx(func(txOrm orm.TxOrmer) error {
		prize := &Prize{Id: id}
		if err := txOrm.Read(prize); err != nil {
			return orm.ErrMissPK
		}
		log.Printf("delete prize %v", prize)
		_, err := txOrm.Delete(prize)
		return err
	})
}

func Get(prize *Prize) ([]*Prize, error) {
	var ps []*Prize
	o := orm.NewOrm()
	qs := o.QueryTable(prize)
	if prize.Id != 0 {
		qs = qs.Filter("id", prize.Id)
	}
	if strings.TrimSpace(prize.Name) != "" {
		qs = qs.Filter("name", prize.Name)
	}
	_, err := qs.All(&ps)
	return ps, err
}

func GetAll() ([]*Prize, error) {
	var ps []*Prize
	o := orm.NewOrm()
	_, err := o.QueryTable(Prize{}).All(&ps)
	return ps, err
}

func argCheck(prize *Prize) bool {
	return prize.Id != 0
}
