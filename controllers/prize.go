package controllers

import (
	"encoding/json"
	"errors"
	"lucky-draw/models"
	"lucky-draw/result"
	"lucky-draw/service"

	beego "github.com/beego/beego/v2/server/web"
)

var prizeService *service.PrizeService

func init() {
	prizeService = &service.PrizeService{}
}

type PrizeController struct {
	beego.Controller
}

// @router /lucky-draw/prize/add [post]
func (p *PrizeController) Add() {
	p.execute(addAction)
}

// @router /lucky-draw/prize/update [post]
func (p *PrizeController) Update() {
	p.execute(updateAction)
}

// @router /lucky-draw/prize/delete [post]
func (p *PrizeController) Delete() {
	p.execute(delAction)
}

// @router /lucky-draw/prize/get [post]
func (p *PrizeController) Get() {
	p.execute(getAction)
}

// @router /lucky-draw/prize/getAll [get]
func (p *PrizeController) GetAll() {
	ps, err := prizeService.GetAll()
	if err != nil {
		p.serverJson(result.ERROR(err.Error()))
	} else {
		p.serverJson(result.OK_RESULT(ps))
	}
}

func (p *PrizeController) execute(action actionType) {
	var prize models.Prize
	var err error
	if err = json.Unmarshal(p.Ctx.Input.RequestBody, &prize); err == nil {
		var ps []*models.Prize
		switch action {
		case addAction:
			err = prizeService.Add(&prize)
		case updateAction:
			err = prizeService.Update(&prize)
		case delAction:
			err = prizeService.Del(prize.Id)
		case getAction:
			ps, err = prizeService.Get(&prize)
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

func (p *PrizeController) serverJson(obj interface{}) {
	p.Data["json"] = obj
	p.ServeJSON()
}
