package handlers

import (
	"net/http"
	"project/database"
	"project/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func GetShifts(c echo.Context) error {
	var shifts []models.Shift
	database.Database.Find(&shifts)
	database.Database.Preload(clause.Associations).Find(&shifts)
	return c.JSON(http.StatusOK, shifts)
}

func CreateShift(c echo.Context) error {
	shift := new(models.Shift)
	if err := c.Bind(shift); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	database.Database.Create(shift)
	return c.JSON(http.StatusOK, shift)

}
