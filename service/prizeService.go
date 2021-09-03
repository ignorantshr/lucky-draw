package service

import "lucky-draw/models"

type PrizeService struct{}

func (service *PrizeService) Add(prize *models.Prize) error {
	return models.AddPrize(prize)
}

func (service *PrizeService) Update(prize *models.Prize) error {
	return models.UpdatePrize(prize)
}

func (service *PrizeService) Del(id int64) error {
	return models.DelPrize(id)
}

func (service *PrizeService) Get(prize *models.Prize) ([]*models.Prize, error) {
	return models.GetPrize(prize)
}

func (service *PrizeService) GetAll() ([]*models.Prize, error) {
	return models.GetAllPrize()
}
