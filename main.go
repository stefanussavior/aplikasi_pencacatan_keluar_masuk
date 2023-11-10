package main

import (
	routes "aplikasi_pencatatan/Routes"
	"aplikasi_pencatatan/database"
)

func main() {
	database.ConnectionDatabase()
	routes.Routes()
}
