package resto

import "github.com/vinopermana77/go-restaurant-app/internal/model"

type Usecase interface {
	// GetMenu = mengambil menu dengan menuType apa aja
	// return sebuah model.MenuItem, dan error jika ada
	GetMenu(menuType string) ([]model.MenuItem, error)
}