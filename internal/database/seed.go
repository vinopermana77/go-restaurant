package database

import (
	"github.com/vinopermana77/go-restaurant-app/internal/model"
	"github.com/vinopermana77/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}
	drinkMenu := []model.MenuItem{
		{
			Name:      "Teh Botol",
			OrderCode: "teh_botol",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
	}

	// Create
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}