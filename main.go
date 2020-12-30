package main

import (
	"encoding/json"
	"fmt"
	"github.com/ie/supermarket-server/handler"
	"github.com/ie/supermarket-server/model"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/customers", handler.Read)
	e.POST("/customers", handler.Customer{}.Create)
	e.PUT("/customers/:cID", handler.Update)
	e.DELETE("/customers/:cID", handler.DeleteUser)

	e.GET("/report/:month", handler.GetReport)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println("error in start server : ", err)
	}

	c := model.Customer{
		Name:    "haniye",
		ID:      1,
		Address: "tehran",
		Tel:     1234,
		//cRegisterDate:
	}
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
