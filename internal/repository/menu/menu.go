package menu

import (
	"github.com/vinopermana77/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB // use db
}

func GetRepository(db *gorm.DB) Repository {
	// mengembalikan sebuah implementasi dari interface repository ini
	return &menuRepo{
		db: db,
	}
}

//GetMenu di ambil dari interface repository.go
func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem

	if err := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}