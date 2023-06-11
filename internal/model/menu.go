package model

type MenuType string

// struct = reusable
type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}