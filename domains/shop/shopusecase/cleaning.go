package shopusecase

import (
	"dddemo/domains/shop"
	"dddemo/models"
)

type CleaningUseCase struct {
	dishRepo shop.DishRepo
}

func NewCleaningUseCase(dishRepo shop.DishRepo) *CleaningUseCase {
	return &CleaningUseCase{
		dishRepo: dishRepo,
	}
}

func (uc CleaningUseCase) CleanIt() ([]models.Trash, error) {
	dishes, _ := uc.dishRepo.RGetDishes()

	dishNames := make([]string, 0, len(dishes))
	for _, el := range dishes {
		dishNames = append(dishNames, el.Name)
		uc.dishRepo.RDeleteDish(el.Name)
	}

	trash := make([]models.Trash, 0, len(dishNames))
	for _, el := range dishNames {
		trash = append(trash, models.Trash{Name: el})
	}

	return trash, nil
}
