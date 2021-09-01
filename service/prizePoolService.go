package service

import (
	"lucky-draw/models"
)

type PrizePoolService interface {
	Add(prizePool *models.PrizePool) error
	Update(prizePool *models.PrizePool) error
	Del(id int64) error
	Get(prizePool *models.PrizePool) ([]*models.PrizePool, error)
	GetAll() ([]*models.PrizePool, error)
}
