package model

type Customer struct {
	Name         string `json:"cName"`
	Tel          int64  `json:"cTel"`
	Address      string `json:"cAddress"`
	ID           int    `json:"cID"`
	RegisterDate string `json:"cRegisterDate"`
	Msg          string `json:"msg"`
}

type ReportResponse struct {
	TotalCustomer int    `json:"totalCustomers"`
	Period        int    `json:"period"`
	Msg           string `json:"msg"`
}
type GeneralReport struct {
	January   int `json:"january"`
	February  int `json:"february"`
	March     int `json:"march"`
	April     int `json:"april"`
	May       int `json:"may"`
	June      int `json:"june"`
	July      int `json:"july"`
	August    int `json:"august"`
	September int `json:"september"`
	October   int `json:"october"`
	November  int `json:"november"`
	December  int `json:"december"`
}

type Message struct {
	Msg string `json:"msg"`
}

type CustomersListCell struct {
	Name         string `json:"cName"`
	Tel          int64  `json:"cTel"`
	Address      string `json:"cAddress"`
	ID           int    `json:"cID"`
	RegisterDate string `json:"cRegisterDate"`
}

type GetUsers struct {
	Size      int                        `json:"size"`
	Customers map[int]*CustomersListCell `json:"customers"`
	Msg       string                     `json:"msg"`
}

var (
	Users = map[int]*Customer{}
	Seq   = 1
)
