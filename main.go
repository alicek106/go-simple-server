package main

import (
	"github.com/labstack/echo"
    _ "github.com/go-sql-driver/mysql"
	"controllers"
)

func main(){
	e := echo.New()
	e.GET("/name", controllers.GetSingleRow)
	e.GET("/all", controllers.GetMultipleRow)
	e.PUT("/name", controllers.InsertRow)
	e.Logger.Fatal(e.Start(":1331"))
}