package service

import "lucky-draw/models"

type PrizePoolService struct{}

func (service *PrizePoolService) Add(prizePool *models.PrizePool) error {
	return models.AddPrizePool(prizePool)
}

func (service *PrizePoolService) Update(prizePool *models.PrizePool) error {
	return models.UpdatePrizePool(prizePool)
}

func (service *PrizePoolService) Del(id int64) error {
	return models.DelPrizePool(id)
}

func (service *PrizePoolService) AddPrize(prizePool *models.PrizePool) error {
	return models.AddPrize2Pool(prizePool)
}

func (service *PrizePoolService) UpdatePrize(prizePool *models.PrizePool) error {
	for _, v := range prizePool.Prizes {
		return models.UpdatePrize4Pool(prizePool.Id, v)
	}
	return nil
}

func (service *PrizePoolService) DelPrize(prizePool *models.PrizePool) error {
	return models.DelPrize4Pool(prizePool)
}

func (service *PrizePoolService) Get(prizePool *models.PrizePool) ([]*models.PrizePool, error) {
	return models.GetPrizePool(prizePool)
}

func (service *PrizePoolService) Info(prizePool *models.PrizePool) (*models.PrizePool, error) {
	return models.InfoPrizePool(prizePool.Id)
}

func (service *PrizePoolService) GetAll() ([]*models.PrizePool, error) {
	return models.GetAllPrizePool()
}
