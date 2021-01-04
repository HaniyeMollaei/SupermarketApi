package main

import (
	"fmt"
	"github.com/ie/supermarket-server/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/customers", handler.Read)
	e.POST("/customers", handler.Customer{}.Create)
	e.PUT("/customers/:cID", handler.Update)
	e.DELETE("/customers/:cID", handler.DeleteUser)

	e.GET("/report/:month", handler.GetReport)
	e.GET("/report", handler.GetGeneralReport)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println("error in start server : ", err)
	}

}
