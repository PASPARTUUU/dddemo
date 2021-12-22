package kitchen

import "errors"

var (
	ErrIngredientNotFound      = errors.New("ingredient not found") 
	ErrInsufficientIngredients = errors.New("insufficient ingredients") // недостаточное количество ингредиентов
)
