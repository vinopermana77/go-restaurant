package menu

import "github.com/vinopermana77/go-restaurant-app/internal/model"

type Repository interface {
	// GetMenu = mengambil menu dengan menuType apa aja
	// return sebuah model.MenuItem, dan error jika ada
	GetMenu(menuType string) ([]model.MenuItem, error)
}