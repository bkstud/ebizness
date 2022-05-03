package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type State string

const (
	Available State = "available"
	Basket    State = "basket"
	Bough     State = "bought"
)

type Product struct {
	Name  string
	Price uint
	State State
	Id    uint
}

func find_product_by_id(id uint, arr []Product) int {
	for i := range arr {
		if arr[i].Id == id {
			return i
		}
	}
	return -1
}

func update_states(new []Product) {
	for _, p := range new {
		id := find_product_by_id(p.Id, available)
		available[id].State = p.State
	}
}

var available = []Product{
	{"banana", 10, Available, 0},
	{"apple", 5, Available, 1},
	{"orange", 2, Available, 2}}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/products", func(c echo.Context) error {
		fmt.Println("[get products]")
		return c.JSON(http.StatusOK, available)
	})

	e.POST("/basket", func(c echo.Context) error {
		products := new([]Product)
		if err := c.Bind(products); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		fmt.Println("[post products]")
		fmt.Println(products)
		update_states(*products)
		return c.JSON(http.StatusOK, products)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
