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

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}
	e.GET("/customers", handler.Read)
	e.POST("/customers", handler.Customer{}.Create2)

	c := model.Customer{
		CName:    "haniye",
		CID:      1,
		CAddress: "tehran",
		CTel:     1234,
		//cRegisterDate:
	}
	b, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
