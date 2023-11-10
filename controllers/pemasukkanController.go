package controllers

import (
	"aplikasi_pencatatan/database"
	"aplikasi_pencatatan/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllDataPemasukkan(c echo.Context) error {
	var Pemasukkan []models.Pemasukan

	if err := database.DB.Raw("SELECT * FROM pemasukans").Scan(&Pemasukkan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, Pemasukkan)
}

func AddDataPemasukkan(c echo.Context) error {
	var Pemasukkan models.Pemasukan

	if err := c.Bind(&Pemasukkan); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	database.DB.Create(&Pemasukkan)
	return c.JSON(http.StatusOK, Pemasukkan)
}
