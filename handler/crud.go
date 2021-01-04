package handler

import (
	"github.com/ie/supermarket-server/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
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

	// we can handle duplicate users with this part if needed
	//
	//var i int
	//for i = 1; i <= len(model.Users); i++ {
	//	if model.Users[i].Name == req.CName && model.Users[i].Tel == req.CTel && model.Users[i].Address == req.CAddress {
	//		return c.JSON(http.StatusNotAcceptable, "user is registered before")
	//	}
	//}

	m := model.Customer{
		Name:         req.CName,
		Tel:          req.CTel,
		Address:      req.CAddress,
		ID:           model.Seq,
		RegisterDate: convertDateToString(time.Now()),
		Msg:          "success",
	}
	model.Users[model.Seq] = &m
	model.Seq++
	return c.JSON(http.StatusCreated, m)
}

func Read(c echo.Context) error {
	var (
		users = map[int]*model.CustomersListCell{}
		seq   = 1
	)

	name := c.QueryParam("cName")
	var i int
	for i = 1; i <= len(model.Users); i++ {
		if strings.Contains(model.Users[i].Name, name) {

			m := model.CustomersListCell{
				ID:           model.Users[i].ID,
				Name:         model.Users[i].Name,
				Address:      model.Users[i].Address,
				Tel:          model.Users[i].Tel,
				RegisterDate: model.Users[i].RegisterDate,
			}
			users[seq] = &m
			seq += 1

		}
	}
	if seq == 1 {
		m := model.Message{
			Msg: "error",
		}
		return c.JSON(http.StatusNotFound, m)
	}
	m := model.GetUsers{
		Size:      seq - 1,
		Customers: users,
		Msg:       "success",
	}
	return c.JSON(http.StatusOK, m)
}

func Update(c echo.Context) error {

	var req Request
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Input isn't valid")
	}
	cID, _ := strconv.Atoi(c.Param("cID"))

	var i int
	for i = 1; i <= len(model.Users); i++ {
		if cID == model.Users[i].ID {
			model.Users[cID].Name = req.CName
			model.Users[cID].Address = req.CAddress
			model.Users[cID].Tel = req.CTel
			model.Users[cID].Msg = "success"
			return c.JSON(http.StatusOK, model.Users[cID])
		}
	}

	m := model.Message{
		Msg: "error",
	}
	return c.JSON(http.StatusBadRequest, m)

}

func DeleteUser(c echo.Context) error {

	cID, _ := strconv.Atoi(c.Param("cID"))
	var req DeleteRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Input isn't valid")
	}
	var i int
	for i = 1; i <= len(model.Users); i++ {
		if cID == model.Users[i].ID {
			delete(model.Users, cID)
			m := model.Message{
				Msg: "success",
			}
			return c.JSON(http.StatusOK, m)
		}
	}

	m := model.Message{
		Msg: "error",
	}

	return c.JSON(http.StatusNotFound, m)
}

func GetGeneralReport(c echo.Context) error {
	usersCount := [...]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var i int
	for i = 1; i <= len(model.Users); i++ {
		var _, mon, _ = convertStringToDate(model.Users[i].RegisterDate)
		switch mon {
		case 1:
			usersCount[0]++
		case 2:
			usersCount[1]++
		case 3:
			usersCount[2]++
		case 4:
			usersCount[3]++
		case 5:
			usersCount[4]++
		case 6:
			usersCount[5]++
		case 7:
			usersCount[6]++
		case 8:
			usersCount[7]++
		case 9:
			usersCount[8]++
		case 10:
			usersCount[9]++
		case 11:
			usersCount[10]++
		case 12:
			usersCount[11]++
		default:
		}
	}
	m := model.GeneralReport{
		January:   usersCount[0],
		February:  usersCount[1],
		March:     usersCount[2],
		April:     usersCount[3],
		May:       usersCount[4],
		June:      usersCount[5],
		July:      usersCount[6],
		August:    usersCount[7],
		September: usersCount[8],
		October:   usersCount[9],
		November:  usersCount[10],
		December:  usersCount[11],
	}
	return c.JSON(http.StatusCreated, m)
}

func GetReport(c echo.Context) error {
	month, _ := strconv.Atoi(c.Param("month"))
	if month < 0 || month > 11 {
		m := model.Message{
			Msg: "error",
		}

		return c.JSON(http.StatusNotFound, m)
	}
	var i int
	var customerCounter = 0
	for i = 1; i <= len(model.Users); i++ {
		var _, mon, _ = convertStringToDate(model.Users[i].RegisterDate)
		if month == mon-1 {
			customerCounter++
		}
	}

	m := model.ReportResponse{
		TotalCustomer: customerCounter,
		Period:        1,
		Msg:           "success",
	}
	return c.JSON(http.StatusCreated, m)
}

func convertDateToString(t time.Time) string {
	var strTime string
	strTime = strings.Split(t.Format(time.RFC3339), "T")[0]
	return strTime
}

func convertStringToDate(s string) (int, int, int) {
	date := strings.Split(s, "-")
	year, _ := strconv.ParseInt(date[0], 0, 32)
	month, _ := strconv.ParseInt(date[1], 0, 32)
	day, _ := strconv.ParseInt(date[2], 0, 32)

	return int(year), int(month), int(day)
}
