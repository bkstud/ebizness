package handlers

import (
	"net/http"
	"project/database"
	"project/models"

	"github.com/labstack/echo/v4"
)

func GetCustomers(c echo.Context) error {
	var customers []models.Customer
	database.Database.Find(&customers)
	return c.JSON(http.StatusOK, customers)
}

func CreateCustomer(c echo.Context) error {
	customer := new(models.Customer)
	if err := c.Bind(customer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	database.Database.Create(customer)
	return c.JSON(http.StatusOK, customer)

}
