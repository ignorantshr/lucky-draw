package impl

import "lucky-draw/models"

type PrizePoolServiceImpl struct{}

func (impl *PrizePoolServiceImpl) Add(prizePool *models.PrizePool) error {
	return models.AddPrizePool(prizePool)
}

func (impl *PrizePoolServiceImpl) Update(prizePool *models.PrizePool) error {
	return models.UpdatePrizePool(prizePool)
}

func (impl *PrizePoolServiceImpl) Del(id int64) error {
	return models.DelPrizePool(id)
}

func (impl *PrizePoolServiceImpl) Get(prizePool *models.PrizePool) ([]*models.PrizePool, error) {
	return models.GetPrizePool(prizePool)
}

func (impl *PrizePoolServiceImpl) GetAll() ([]*models.PrizePool, error) {
	return models.GetAllPrizePool()
}
