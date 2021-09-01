package service

import (
	"lucky-draw/models"
)

type PrizeService interface {
	Add(prize *models.Prize) error
	Del(id int64) error
	Update(prize *models.Prize) error
	Get(prize *models.Prize) ([]*models.Prize, error)
	GetAll() ([]*models.Prize, error)
}
