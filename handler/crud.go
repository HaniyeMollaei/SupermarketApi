package handler

import (
	"fmt"
	"github.com/ie/supermarket-server/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	Name string
	Id   int64
}
type Customer struct {
}

func (customer Customer) Create(c echo.Context) error {
	var m model.Customer
	if err := c.Bind(&m); err != nil {
		fmt.Println(err)
	}
	return nil
}

func (customer Customer) Create2(c echo.Context) error {
	var req Request

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Input isn't valid")
	}
	m := model.Customer{
		CName:    req.Name,
		CTel:     12345,
		CAddress: "Znj . Iran",
		CID:      req.Id,
		//cRegisterDate time.Time `json:"date"`
	}
	return c.JSON(http.StatusCreated, m)
}

func Read(c echo.Context) error {

	return c.JSON(http.StatusOK, "Hello World")
}

func Update(c echo.Context) error {
	return nil
}

func Delet(c echo.Context) error {
	return nil
}
