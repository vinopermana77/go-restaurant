package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

func main() {
	seedDB()
	e := echo.New()

	// localhost:14045
	e.GET("/menu", getMenu)

	e.Logger.Fatal(e.Start(":14045"))
}

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

// struct = reusable
type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      MenuTypeFood,
		},
	}
	drinkMenu := []MenuItem{
		{
			Name:      "Teh Botol",
			OrderCode: "teh_botol",
			Price:     7000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     5000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinkMenu)

	// connect database
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	// Create
	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}

// Method
func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

// func getFoodMenu(c echo.Context) error {
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}
// 	var menuData []MenuItem
// 	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }

// func getDrinkMenu(c echo.Context) error {
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}
// 	var menuData []MenuItem
// 	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }
