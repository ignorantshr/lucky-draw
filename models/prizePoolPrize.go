package models

import (
	"fmt"
	"lucky-draw/result"

	"github.com/beego/beego/v2/client/orm"
)

type PrizePoolPrize struct {
	Id               int64 `json:"id"`
	PrizePoolId      int64
	PrizeId          int64
	PrizeProbability int
	PrizeNumber      int64
}

var mapperOrm orm.Ormer
var mapperQs orm.QuerySeter

func init() {
	mapperOrm = orm.NewOrm()
	mapperQs = mapperOrm.QueryTable(new(PrizePoolPrize))
}

// 添加奖品到奖池
func AddPrize2Pool(prizePool *PrizePool) error {
	id := prizePool.Id
	if !argCheck(id) {
		return orm.ErrArgs
	}

	for _, v := range prizePool.Prizes {
		if isExist(id, v.Id) {
			return fmt.Errorf(result.EXIST_ERROR)
		}
	}

	return dotx(func(myTxOrm orm.TxOrmer) error {
		size := len(prizePool.Prizes)
		mappers := make([]*PrizePoolPrize, size)
		for i, v := range prizePool.Prizes {
			mappers[i] = genRecord(id, v)
		}
		if num, err := myTxOrm.InsertMulti(size, mappers); err != nil || num == 0 {
			return fmt.Errorf(result.POOL_ADD_PRIZE_ERROR, id, mappers)
		}
		return nil
	})
}

// 更新奖池的奖品
func UpdatePrize4Pool(id int64, prize *Prize) error {
	if !argCheck(id) || !argCheck(prize.Id) {
		return orm.ErrArgs
	}

	if !isExist(id, prize.Id) {
		return fmt.Errorf(result.NOT_EXIST_ERROR)
	}

	return dotx(func(myTxOrm orm.TxOrmer) error {
		mapper := genRecord(id, prize)
		if num, err := myTxOrm.
			QueryTable(PrizePoolPrize{}).
			Filter("PrizePoolId", mapper.PrizePoolId).Filter("PrizeId", mapper.PrizeId).
			Update(orm.Params{
				"PrizeProbability": mapper.PrizeProbability,
				"PrizeNumber":      mapper.PrizeNumber,
			}); err != nil || num != 1 {
			return fmt.Errorf(result.POOL_UPDATE_PRIZE_ERROR, id, mapper)
		}
		return nil
	})
}

// 从奖池删除奖品
func DelPrize4Pool(prizePool *PrizePool) error {
	poolId := prizePool.Id
	if !argCheck(poolId) {
		return orm.ErrArgs
	}

	prizeIds := make([]*int64, len(prizePool.Prizes))
	for i, v := range prizePool.Prizes {
		if !isExist(poolId, v.Id) {
			return fmt.Errorf(result.NOT_EXIST_ERROR)
		}
		prizeIds[i] = &v.Id
	}

	return dotx(func(myTxOrm orm.TxOrmer) error {
		if num, err := myTxOrm.QueryTable(PrizePoolPrize{}).Filter("PrizePoolId", poolId).Filter("PrizeId__in", prizeIds).Delete(); err != nil || num == 0 {
			return fmt.Errorf(result.POOL_DEL_PRIZE_ERROR, poolId, prizePool.Prizes)
		}
		return nil
	})
}

func argCheck(id int64) bool {
	return id > 0
}

func isExist(poolId int64, prizeId int64) bool {
	return mapperQs.Filter("PrizePoolId", poolId).Filter("PrizeId", prizeId).Exist()
}

func normal(prize *Prize) {
	if prize.Probability < 0 {
		prize.Probability = 0
	}
}

func genRecord(id int64, prize *Prize) *PrizePoolPrize {
	normal(prize)
	return &PrizePoolPrize{
		PrizePoolId:      id,
		PrizeId:          prize.Id,
		PrizeProbability: prize.Probability,
		PrizeNumber:      prize.Number,
	}
}
