package controllers

type actionType int

const (
	addAction actionType = iota
	updateAction
	deleteAction
	getAction
)
