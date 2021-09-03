package models

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
)

type BaseModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func init() {
	registerDB()
	// register model
	orm.RegisterModel(new(Prize))
	orm.RegisterModel(new(PrizePool))
	orm.RegisterModel(new(PrizePoolPrize))

	// create table
	// orm.RunSyncdb("default", false, true)
}

func registerDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:Lenovo123-@tcp(10.221.5.7:3306)/lucky_draw?charset=utf8&loc=Local")
}

func idCheck(model *BaseModel) bool {
	return model.Id > 0
}

func dotx(f func(myTxOrm orm.TxOrmer) error) error {
	o := orm.NewOrm()
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		return f(txOrm)
	})
	return err
}
