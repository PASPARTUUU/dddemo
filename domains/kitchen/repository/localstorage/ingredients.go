package localstorage

import (
	"sync"

	"dddemo/domains/kitchen"
	"dddemo/domains/kitchen/aggregates"
	"dddemo/models"
)

type IngredientLocalStorage struct {
	ingredients map[string]aggregates.Ingredient
	mutex       *sync.Mutex
}

func NewIngredientLocalStorage() *IngredientLocalStorage {
	return &IngredientLocalStorage{
		ingredients: make(map[string]aggregates.Ingredient),
		mutex:       new(sync.Mutex),
	}
}

func (s *IngredientLocalStorage) RAddIngredient(ingr models.Ingredient) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if el, ex := s.ingredients[ingr.Name]; ex {
		el.Count++
		s.ingredients[ingr.Name] = el
	} else {
		s.ingredients[ingr.Name] = aggregates.Ingredient{
			Ingredient: ingr,
			Count:      1,
		}
	}

	return nil
}

func (s *IngredientLocalStorage) RTakeIngredients(names ...string) ([]models.Ingredient, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	res := []models.Ingredient{}
	for _, n := range names {
		if el, ex := s.ingredients[n]; ex {
			if el.Count == 0 {
				return nil, kitchen.ErrInsufficientIngredients
			}

			el.Count--
			s.ingredients[el.Name] = el

			res = append(res, el.Ingredient)
		} else {
			return nil, kitchen.ErrIngredientNotFound
		}
	}

	return res, nil
}

func (s *IngredientLocalStorage) RShowIngredients() ([]aggregates.Ingredient, error) {
	res := []aggregates.Ingredient{}

	for _, el := range s.ingredients {
		res = append(res, el)
	}

	return res, nil
}
