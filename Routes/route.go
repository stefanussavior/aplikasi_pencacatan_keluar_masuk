package routes

import (
	"aplikasi_pencatatan/controllers"
	"aplikasi_pencatatan/database"

	"github.com/labstack/echo/v4"
)

func Routes() {
	database.ConnectionDatabase()

	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	//user routes
	e.GET("/getuser", controllers.GetAllUser)

	//pemasukkan routes
	e.GET("/getallpemasukkan", controllers.GetAllDataPemasukkan)
	e.POST("/addpemasukkan", controllers.AddDataPemasukkan)

	//pengeluaran routes
	e.GET("/getallpengeluaran", controllers.GetAllDataPengeluaran)
	e.POST("/addpengeluaran", controllers.AddDataPengeluaran)

	e.Logger.Fatal(e.Start(":4040"))
}
