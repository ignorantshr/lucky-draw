package controllers

import (
	"encoding/json"
	"errors"
	"lucky-draw/models"
	"lucky-draw/result"
	"lucky-draw/service"
	"lucky-draw/service/impl"

	beego "github.com/beego/beego/v2/server/web"
)

var prizePoolService service.PrizePoolService

func init() {
	prizePoolService = &impl.PrizePoolServiceImpl{}
}

type PrizePoolController struct {
	beego.Controller
}

// @router /prizePool/add [post]
func (p *PrizePoolController) Add() {
	p.execute(addAction)
}

// @router /prizePool/update [post]
func (p *PrizePoolController) Update() {
	p.execute(updateAction)
}

// @router /prizePool/delete [post]
func (p *PrizePoolController) Delete() {
	p.execute(deleteAction)
}

// @router /prizePool/get [post]
func (p *PrizePoolController) Get() {
	p.execute(getAction)
}

// @router /prizePool/getAll [get]
func (p *PrizePoolController) GetAll() {
	ps, err := prizePoolService.GetAll()
	if err != nil {
		p.serverJson(result.ERROR(err.Error()))
	} else {
		p.serverJson(result.OK_RESULT(ps))
	}
}

func (p *PrizePoolController) execute(action actionType) {
	var prizePool models.PrizePool
	var err error
	if err = json.Unmarshal(p.Ctx.Input.RequestBody, &prizePool); err == nil {
		var ps []*models.PrizePool
		switch action {
		case addAction:
			err = prizePoolService.Add(&prizePool)
		case updateAction:
			err = prizePoolService.Update(&prizePool)
		case deleteAction:
			err = prizePoolService.Del(prizePool.Id)
		case getAction:
			ps, err = prizePoolService.Get(&prizePool)
		default:
			err = errors.New("no action to execute")
		}

		if err != nil {
			p.serverJson(result.ERROR(err.Error()))
		} else {
			p.serverJson(result.OK_RESULT(ps))
		}
	} else {
		p.serverJson(result.ERROR(err.Error()))
	}
}

func (p *PrizePoolController) serverJson(obj interface{}) {
	p.Data["json"] = obj
	p.ServeJSON()
}
