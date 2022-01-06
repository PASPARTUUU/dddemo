package models

import "time"

// Dish - блюдо-кушанье
type Dish struct {
	Name   string
	Recipe string
}

// Trash - 
type Trash struct {
	Name     string
	SomeData interface{} // не используется
}

// Ingredient - ингредиенты для блюд
type Ingredient struct {
	Name string
	Exp  time.Time
}

// ExpDate - срок годности
// [#02]
func (ingr Ingredient) ExpDate() time.Time {
	return time.Now().Add(time.Hour * 48)
}

// SidebarMarkup - структура инициализации раздела на боковой панели для домена
type SidebarMarkup struct {
	Name       string            // заголовок раздела
	LI         []SidebarMarkupLI // елементы-ссылки раздела
	Permission []string
}

type SidebarMarkupLI struct {
	Name       string
	Href       string
	Permission []string
}
