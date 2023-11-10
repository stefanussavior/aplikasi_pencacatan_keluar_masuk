package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type JwtClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Nama        string `json:"nama"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Pemasukan   []Pemasukan
	Pengeluaran []Pengeluaran
}

type Pemasukan struct {
	gorm.Model
	UserID    uint    `json:"id_user"`
	Gaji      float32 `json:"gaji"`
	Tunjangan float32 `json:"tunjangan"`
	Bonus     float32 `json:"bonus"`
}

type Pengeluaran struct {
	gorm.Model
	UserID        uint    `json:"id_user"`
	SewaKos       float32 `json:"sewa_kos"`
	Makan         float32 `json:"makan"`
	Pakaian       float32 `json:"pakaian"`
	NontonBioskop float32 `json:"nonton_bioskop"`
}
