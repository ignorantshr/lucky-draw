package impl

import "lucky-draw/models"

type PrizeServiceImpl struct{}

func (impl *PrizeServiceImpl) Add(prize *models.Prize) error {
	return models.AddPrize(prize)
}

func (impl *PrizeServiceImpl) Update(prize *models.Prize) error {
	return models.UpdatePrize(prize)
}

func (impl *PrizeServiceImpl) Del(id int64) error {
	return models.DelPrize(id)
}

func (impl *PrizeServiceImpl) Get(prize *models.Prize) ([]*models.Prize, error) {
	return models.GetPrize(prize)
}

func (impl *PrizeServiceImpl) GetAll() ([]*models.Prize, error) {
	return models.GetAllPrize()
}
