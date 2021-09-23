package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"lucky-draw/models"
	"lucky-draw/result"
	"lucky-draw/service"
	"math/rand"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

var prizePoolService *service.PrizePoolService

func init() {
	rand.Seed(time.Now().UnixNano())
	prizePoolService = &service.PrizePoolService{}
}

type PrizePoolController struct {
	beego.Controller
}

// @router /lucky-draw/prizePool/add [post]
func (p *PrizePoolController) Add() {
	p.execute(addAction)
}

// @router /lucky-draw/prizePool/update [post]
func (p *PrizePoolController) Update() {
	p.execute(updateAction)
}

// @router /lucky-draw/prizePool/delete [post]
func (p *PrizePoolController) Delete() {
	p.execute(delAction)
}

// @router /lucky-draw/prizePool/addPrize [post]
func (p *PrizePoolController) AddPrize() {
	p.execute(poolAddPrizeAction)
}

// @router /lucky-draw/prizePool/updatePrize [post]
func (p *PrizePoolController) UpdatePrize() {
	p.execute(poolUpdatePrizeAction)
}

// @router /lucky-draw/prizePool/delPrize [post]
func (p *PrizePoolController) DelPrize4Pool() {
	p.execute(poolDelPrizeAction)
}

// @router /lucky-draw/prizePool/get [post]
func (p *PrizePoolController) Get() {
	p.execute(getAction)
}

// @router /lucky-draw/prizePool/info [post]
func (p *PrizePoolController) Info() {
	p.execute(infoAction)
}

// @router /lucky-draw/prizePool/getAll [get]
func (p *PrizePoolController) GetAll() {
	ps, err := prizePoolService.GetAll()
	if err != nil {
		p.serverJson(result.ERROR(err.Error()))
	} else {
		p.serverJson(result.OK_RESULT(ps))
	}
}

// @router /lucky-draw/prizePool/getPrizes [post]
func (p *PrizePoolController) GetPrizes() {
	var query models.PoolPrizeQuery
	var err error
	if err = json.Unmarshal(p.Ctx.Input.RequestBody, &query); err != nil {
		p.serverJson(result.ERROR(err.Error()))
	}
	prizes, err := prizePoolService.GetPrizes(&query)
	if err != nil {
		p.serverJson(result.ERROR(err.Error()))
	}
	p.serverJson(result.OK_RESULT(prizes))
}

// @router /lucky-draw/prizePool/getUnpoolPrizes [post]
func (p *PrizePoolController) GetUnpoolPrizes() {
	var prizePool *models.PrizePool
	var err error
	if err = json.Unmarshal(p.Ctx.Input.RequestBody, &prizePool); err == nil {
		prizes, err := prizePoolService.GetUnpoolPrizes(prizePool)

		if err != nil {
			p.serverJson(result.ERROR(err.Error()))
		} else if prizes != nil {
			p.serverJson(result.OK_RESULT(prizes))
		}
	}
}

// @router /lucky-draw/prizePool/draw [get]
func (p *PrizePoolController) Draw() {
	var err error
	pooId, err := p.GetInt64("id")
	if err != nil {
		log.Println("获取参数错误", err)
		p.serverJson(result.ERROR(err.Error()))
		return
	}

	var pool *models.PrizePool = &models.PrizePool{}
	pool.Id = pooId
	pool, err = prizePoolService.Info(pool)
	if err != nil {
		log.Println("获取参数错误", err)
		p.serverJson(result.ERROR(err.Error()))
		return
	}
	log.Printf("pool: %v\n", pool)

	prizes := make(map[int64]*models.Prize, len(pool.Prizes))
	for _, v := range pool.Prizes {
		prizes[v.Id] = v
	}

	switch pool.Type {
	case models.ProbabilityType:
		// 概率数字范围	 {id: [起始数字，终结数字]}
		probMapper := make(map[int64][2]int)
		end := 0
		for _, v := range pool.Prizes {
			probMapper[v.Id] = [2]int{end, end + v.Probability}
			end += v.Probability
		}
		log.Printf("priz map: %v", probMapper)
		var n int
		n = rand.Intn(end)
		for k, v := range probMapper {
			if v[0] <= n && n < v[1] {
				p.serverJson(result.OK_RESULT(prizes[k]))
				return
			}
		}
	case models.NumberType:
		// 数量范围	 {id: [起始数字，终结数字]}
		numMapper := make(map[int64][2]int64)
		var end int64 = 0
		var nullNum int64 = 0 // 空奖数量
		for _, v := range pool.Prizes {
			if v.Id == 1 {
				nullNum = v.Number
				continue
			}
			numMapper[v.Id] = [2]int64{end, end + v.Number}
			end += v.Number
		}
		log.Printf("prize map: %v", numMapper)
		var n int64
		n = rand.Int63n(end + nullNum)
		// 未中奖直接返回空值
		if n >= end {
			p.serverJson(result.OK())
			return
		}
		for k, v := range numMapper {
			if v[0] <= n && n < v[1] {
				prize := prizes[k]
				prize.Number--
				modifyPool := &models.PrizePool{}
				modifyPool.Id = pooId
				modifyPool.Prizes = []*models.Prize{prize}
				if err = prizePoolService.UpdatePrize(modifyPool); err != nil {
					log.Println("update prize for pool failed: ", err)
					p.serverJson(result.ERROR(err.Error()))
				} else {
					p.serverJson(result.OK_RESULT(prizes[k]))
				}
				return
			}
		}
	default:
		p.serverJson(result.OK())
		return
	}
}

func (p *PrizePoolController) execute(action actionType) {
	var prizePool models.PrizePool
	var err error
	if err = json.Unmarshal(p.Ctx.Input.RequestBody, &prizePool); err == nil {
		var ps []*models.PrizePool
		var pool *models.PrizePool
		switch action {
		case addAction:
			err = prizePoolService.Add(&prizePool)
		case updateAction:
			err = prizePoolService.Update(&prizePool)
		case delAction:
			err = prizePoolService.Del(prizePool.Id)
		case getAction:
			ps, err = prizePoolService.Get(&prizePool)
		case infoAction:
			pool, err = prizePoolService.Info(&prizePool)
		case poolAddPrizeAction:
			err = prizePoolService.AddPrize(&prizePool)
		case poolUpdatePrizeAction:
			err = prizePoolService.UpdatePrize(&prizePool)
		case poolDelPrizeAction:
			err = prizePoolService.DelPrize(&prizePool)
		default:
			err = errors.New("no action to execute")
		}

		if err != nil {
			p.serverJson(result.ERROR(err.Error()))
		} else {
			if ps != nil {
				p.serverJson(result.OK_RESULT(ps))
			} else if pool != nil {
				p.serverJson(result.OK_RESULT(pool))
			}
		}
		p.serverJson((result.OK()))
	} else {
		p.serverJson(result.ERROR(err.Error()))
	}
}

func (p *PrizePoolController) serverJson(obj interface{}) {
	p.Data["json"] = obj
	p.ServeJSON()
}
