package handlers

import (
	"net/http"
	"project/database"
	"project/models"

	"github.com/labstack/echo/v4"
)

func GetProducts(c echo.Context) error {
	var products []models.Product
	database.Database.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	database.Database.Create(product)
	return c.JSON(http.StatusOK, product)

}
