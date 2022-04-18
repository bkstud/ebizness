package handlers

import (
	"net/http"
	"project/database"
	"project/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func GetPurchases(c echo.Context) error {
	var purchases []models.Purchase
	database.Database.Find(&purchases)
	database.Database.Preload(clause.Associations).Find(&purchases)
	return c.JSON(http.StatusOK, purchases)
}

func CreatePurchase(c echo.Context) error {
	purchase := new(models.Purchase)
	if err := c.Bind(purchase); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	database.Database.Create(purchase)
	return c.JSON(http.StatusOK, purchase)

}
