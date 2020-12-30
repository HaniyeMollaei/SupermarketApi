package handler

import (
	"github.com/ie/supermarket-server/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type Request struct {
	CName    string
	CTel     int64
	CAddress string
}

type DeleteRequest struct {
	CID int
}

type Customer struct {
}

func (customer Customer) Create(c echo.Context) error {

	var req Request
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Input isn't valid")
	}
	var i int
	for i = 1; i <= len(model.Users); i++ {
		if model.Users[i].Name == req.CName && model.Users[i].Tel == req.CTel && model.Users[i].Address == req.CAddress {
			return c.JSON(http.StatusNotAcceptable, "user is registered before")
		}
	}
	m := model.Customer{
		Name:         req.CName,
		Tel:          req.CTel,
		Address:      req.CAddress,
		ID:           model.Seq,
		RegisterDate: time.Now(),
	}
	model.Users[model.Seq] = &m
	model.Seq++
	return c.JSON(http.StatusCreated, m)
}

func Read(c echo.Context) error {
	return c.JSON(http.StatusOK, model.Users)
}

func Update(c echo.Context) error {

	var req Request
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Input isn't valid")
	}
	cID, _ := strconv.Atoi(c.Param("cID"))
	model.Users[cID].Name = req.CName
	model.Users[cID].Address = req.CAddress
	model.Users[cID].Tel = req.CTel
	return c.JSON(http.StatusOK, model.Users[cID])
}

func DeleteUser(c echo.Context) error {

	cID, _ := strconv.Atoi(c.Param("cID"))
	var req DeleteRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Input isn't valid")
	}
	delete(model.Users, cID)
	return c.JSON(http.StatusOK, "{\"msg\" : \"success\"}")
}

func GetReport(c echo.Context) error {
	month, _ := strconv.Atoi(c.Param("month"))
	month *= month
	return nil
}
