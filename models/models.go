package models

import "time"

// блюдо
type Dish struct {
	Name   string
	Recipe string
}

// ингредиенты для блюд
type Trash struct {
	Name     string
	SomeData interface{} // не используется
}

// ингредиенты для блюд
type Ingredient struct {
	Name string
	Exp  time.Time
}

// ExpDate - срок годности
// [#02]
func (ingr Ingredient) ExpDate() time.Time {
	return time.Now().Add(time.Hour * 48)
}
