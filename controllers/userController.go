package controllers

import (
	"aplikasi_pencatatan/database"
	"aplikasi_pencatatan/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	var Users []models.User

	if err := database.DB.Preload("Pemasukan").Preload("Pengeluaran").Find(&Users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, Users)
}

func LoginUser(c echo.Context) error {
	Email := c.FormValue("email")
	Password := c.FormValue("password")

	token, err := generateJWTToken(Email, Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func generateJWTToken(Email string, Password string) (string, error) {
	claims := &models.JwtClaims{
		Email: Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
