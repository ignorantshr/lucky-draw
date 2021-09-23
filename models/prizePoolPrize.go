package models

import (
	"fmt"
	"log"
	"lucky-draw/result"
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

// 添加新的奖品到奖池
func AddNewPrize2Pool(pool *PrizePool) error {
	id := pool.Id
	if !argCheck(id) {
		return result.PARAM_INVALID
	}
	var err error

	for _, p := range pool.Prizes {
		if err = AddPrize(p); err != nil {
			return err
		}
		tmp, err := GetPrizeByName(p.Name)
		if err != nil {
			return err
		}
		p.Id = tmp.Id

		normal(p)
		mapper := &PrizePoolPrize{
			PrizePoolId:      id,
			PrizeId:          p.Id,
			PrizeProbability: p.Probability,
			PrizeNumber:      p.Number,
		}

		if err = db.Create(mapper).Error; err != nil {
			log.Println(err)
			return fmt.Errorf(result.POOL_ADD_PRIZE_ERROR, id)
		}
	}
	return nil
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
		return result.PARAM_INVALID
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
		return result.PARAM_INVALID
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

// 获取所有未附加到此奖池的奖品
func GetUnpoolPrizes(id int64) ([]*Prize, error) {
	var err error

	var mapper []*PrizePoolPrize = make([]*PrizePoolPrize, 0)
	if err = db.Where("prize_pool_id = ?", id).Find(&mapper).Error; err != nil {
		return nil, err
	}

	prizeIds := make([]int64, len(mapper))
	for i, v := range mapper {
		prizeIds[i] = v.PrizeId
	}

	log.Printf("now prizes: %v\n", prizeIds)
	prizes := make([]*Prize, 0)
	if len(prizeIds) == 0 {
		return GetAllPrize()
	}
	if err = db.Where("id not in ?", prizeIds).Find(&prizes).Error; err != nil {
		return nil, err
	}

	log.Printf("remain prizes: %s\n", prizes)

	return prizes, err
}

// 查询附加到此奖池的奖品
func GetPrizes(query *PoolPrizeQuery) ([]*Prize, error) {
	type PrizeResult struct {
		// 基本属性
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Url  string `json:"url"`

		// 附加属性
		Probability int   `json:"probability"` // 概率
		Number      int64 `json:"number"`      // 数量
	}
	var err error
	var pr PrizeResult
	prs := make([]*PrizeResult, 0)
	baseSql := "select prize.*, m.prize_probability as probability, m.prize_number as number from prize left join prize_pool_prize m on prize.id = m.prize_id "

	if query.PrizeId != 0 {
		if err = db.Raw(baseSql+"where m.prize_pool_id = ? and m.prize_id = ?",
			query.PoolId, query.PoolId).Scan(&pr).Error; err != nil {
			return nil, err
		}
		prs = append(prs, &pr)
	} else {
		if err = db.Raw(baseSql+"where m.prize_pool_id = ? and prize.name like ?",
			query.PoolId, "%"+query.PrizeName+"%").Scan(&prs).Error; err != nil {
			return nil, err
		}
	}

	prizes := make([]*Prize, len(prs))
	for i, v := range prs {
		prizes[i] = &Prize{
			BaseModel: BaseModel{
				Id:   v.Id,
				Name: v.Name,
			},
			Url:         v.Url,
			Probability: v.Probability,
			Number:      v.Number,
		}
	}

	return prizes, nil
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
	if prize.Number < 0 {
		prize.Number = 0
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
