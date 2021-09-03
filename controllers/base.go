package controllers

type actionType int

const (
	addAction actionType = iota
	updateAction
	delAction
	getAction
	infoAction

	poolAddPrizeAction
	poolUpdatePrizeAction
	poolDelPrizeAction
)
