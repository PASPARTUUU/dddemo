package cleaning

import (
	"dddemo/domains/kitchen"
	"dddemo/domains/kitchen/kitchenusecase"
	"dddemo/domains/shop"
	"dddemo/domains/shop/shopusecase"
	"dddemo/models"
)

// [#01] подтягивание зависимостей из разных доменов
type ServiceCleaning struct {
	kitchenUC kitchen.Cleaning
	shopUC    shop.Cleaning
}

func NewServiceCleaning(ingredientRepo kitchen.IngredientRepo, dishRepo shop.DishRepo) *ServiceCleaning {
	return &ServiceCleaning{
		kitchenUC: kitchenusecase.NewCleaningUseCase(ingredientRepo),
		shopUC:    shopusecase.NewCleaningUseCase(dishRepo),
	}
}

// [#01] имплементация юзкейса работающего с юзкейсами из разных доменов
func (h ServiceCleaning) CleanShopAndKitchen() ([]models.Trash, error) {
	trash := []models.Trash{}

	kichTrash, _ := h.kitchenUC.CleanIt()
	trash = append(trash, kichTrash...)

	shopTrash, _ := h.shopUC.CleanIt()
	trash = append(trash, shopTrash...)

	return trash, nil
}
