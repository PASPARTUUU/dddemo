package service

import "dddemo/models"

type Cleaning interface {
	// [#01] юзкейс предполагающий работу с разными доменами, от того и расположен в service
	CleanShopAndKitchen() ([]models.Trash, error)
}
