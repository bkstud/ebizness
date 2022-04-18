package handlers

import (
	"net/http"
	"project/database"
	"project/models"

	"github.com/labstack/echo/v4"
)

func GetWorkers(c echo.Context) error {
	var workers []models.Worker
	database.Database.Find(&workers)
	return c.JSON(http.StatusOK, workers)
}

func CreateWorker(c echo.Context) error {
	worker := new(models.Worker)
	if err := c.Bind(worker); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	database.Database.Create(worker)
	return c.JSON(http.StatusOK, worker)

}
