package impl

import "lucky-draw/models"

type PrizeServiceImpl struct{}

func (impl *PrizeServiceImpl) Add(prize *models.Prize) error {
	return models.Add(prize)
}

func (impl *PrizeServiceImpl) Update(prize *models.Prize) error {
	return models.Update(prize)
}

func (impl *PrizeServiceImpl) Del(id int64) error {
	return models.Del(id)
}

func (impl *PrizeServiceImpl) Get(prize *models.Prize) ([]*models.Prize, error) {
	return models.Get(prize)
}

func (impl *PrizeServiceImpl) GetAll() ([]*models.Prize, error) {
	return models.GetAll()
}
