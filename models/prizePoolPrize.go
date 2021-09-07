package models

import (
	"fmt"
	"log"
	"lucky-draw/result"

	"github.com/beego/beego/v2/client/orm"
)

type PrizePoolPrize struct {
	PrizePoolId      int64 `gorm:"primaryKey;autoIncrement:false"`
	PrizeId          int64 `gorm:"primaryKey;autoIncrement:false"`
	PrizeProbability int
	PrizeNumber      int64
}

func (PrizePoolPrize) TableName() string {
	return "prize_pool_prize"
}

// 添加奖品到奖池
func AddPrize2Pool(prizePool *PrizePool) error {
	id := prizePool.Id
	if !argCheck(id) {
		return result.PARAM_INVALID
	}

	for _, v := range prizePool.Prizes {
		if isExist(id, v.Id) {
			return fmt.Errorf(result.EXIST_ERROR)
		}
	}

	size := len(prizePool.Prizes)
	mappers := make([]*PrizePoolPrize, size)
	for i, v := range prizePool.Prizes {
		mappers[i] = genRecord(id, v)
	}
	if err := db.CreateInBatches(mappers, size).Error; err != nil {
		log.Println(err)
		return fmt.Errorf(result.POOL_ADD_PRIZE_ERROR, id)
	}
	return nil
}

// 更新奖池的奖品
func UpdatePrize4Pool(id int64, prize *Prize) error {
	if !argCheck(id) || !argCheck(prize.Id) {
		return orm.ErrArgs
	}

	if !isExist(id, prize.Id) {
		return fmt.Errorf(result.NOT_EXIST_ERROR)
	}

	mapper := genRecord(id, prize)
	if err := db.Model(mapper).Updates(mapper).Error; err != nil {
		log.Println()
		return fmt.Errorf(result.POOL_UPDATE_PRIZE_ERROR, id)
	}
	return nil
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

	if err := db.Where("prize_pool_id = ? and prize_id in ?", poolId, prizeIds).Delete(PrizePoolPrize{}).Error; err != nil {
		return fmt.Errorf(result.POOL_DEL_PRIZE_ERROR, poolId)
	}
	return nil
}

func argCheck(id int64) bool {
	return id > 0
}

func isExist(poolId int64, prizeId int64) bool {
	re := db.Where(&PrizePoolPrize{PrizePoolId: poolId, PrizeId: prizeId}).Limit(1).Find(&PrizePoolPrize{})
	return re.RowsAffected > 0
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
