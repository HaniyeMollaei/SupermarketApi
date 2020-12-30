package model

import "time"

type Customer struct {
	Name         string    `json:"cName"`
	Tel          int64     `json:"cTel"`
	Address      string    `json:"cAddress"`
	ID           int       `json:"cID"`
	RegisterDate time.Time `json:"cRegisterDate"`
}

var (
	Users = map[int]*Customer{}
	Seq   = 1
)
