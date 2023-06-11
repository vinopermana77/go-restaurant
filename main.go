package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	seedDB()
	e := echo.New()

	// localhost:14045
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)

	e.Logger.Fatal(e.Start(":14045"))
}

// struct = reusable
type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
		},
	}
	drinkMenu := []MenuItem{
		{
			Name:      "Teh Botol",
			OrderCode: "teh_botol",
			Price:     7000,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     5000,
		},
	}

	fmt.Println(foodMenu, drinkMenu)
	// connect database
	dbAddress := "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	// migrate database
	db.AutoMigrate(&MenuItem{})

}

func getFoodMenu(c echo.Context) error {
	// return c.JSON(http.StatusOK, foodMenu) - Simple

	// Spesifik
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "data": foodMenu,
	})
}
func getDrinkMenu(c echo.Context) error {
	// return c.JSON(http.StatusOK, drinkMenu) - Simple

	// Spesifik
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "data": drinkMenu,
	})
}
