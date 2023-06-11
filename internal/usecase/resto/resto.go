package resto

import (
	"github.com/vinopermana77/go-restaurant-app/internal/model"
	"github.com/vinopermana77/go-restaurant-app/internal/repository/menu"
)

type restoUsecase struct {
	// Mengambil menuRepo di folder Repository
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	// mengembalikan sebuah implementasi dari interface usecase ini
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}

// Implementasi
func (r *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}