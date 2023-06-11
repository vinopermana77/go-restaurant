package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// struct = reusable
type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
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
	// return c.JSON(http.StatusOK, foodMenu) - Simple

	// Spesifik
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}
func getDrinkMenu(c echo.Context) error {
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
	// return c.JSON(http.StatusOK, drinkMenu) - Simple

	// Spesifik
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main() {
	e := echo.New()
	// localhost:14045
	e.GET("/menu/food", getDrinkMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
