package controllers

import (
	"aplikasi_pencatatan/database"
	"aplikasi_pencatatan/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllDataPengeluaran(c echo.Context) error {
	var Pengeluaran []models.Pengeluaran

	if err := database.DB.Raw("SELECT * FROM pengeluarans").Scan(&Pengeluaran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, Pengeluaran)
}

func AddDataPengeluaran(c echo.Context) error {
	var Pengeluaran models.Pengeluaran

	if err := c.Bind(&Pengeluaran); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	database.DB.Create(&Pengeluaran)
	return c.JSON(http.StatusOK, Pengeluaran)
}
