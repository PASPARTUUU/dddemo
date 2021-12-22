package localstorage

import (
	"sync"

	"dddemo/domains/shop"
	"dddemo/domains/shop/aggregates"
)

type DishLocalStorage struct {
	dishes map[string]aggregates.Dish
	mutex  *sync.Mutex
}

func NewDishLocalStorage() *DishLocalStorage {
	return &DishLocalStorage{
		dishes: make(map[string]aggregates.Dish),
		mutex:  new(sync.Mutex),
	}
}

func (s *DishLocalStorage) RCreateDish(dish aggregates.Dish) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.dishes[dish.Name] = dish

	return nil
}

func (s *DishLocalStorage) RGetDishes() ([]aggregates.Dish, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	dishes := make([]aggregates.Dish, 0)
	for _, d := range s.dishes {
		dishes = append(dishes, d)
	}

	return dishes, nil
}

func (s *DishLocalStorage) RDeleteDish(dishName string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ex := s.dishes[dishName]; ex {
		delete(s.dishes, dishName)
		return nil
	}

	return shop.ErrDishNotFound
}
