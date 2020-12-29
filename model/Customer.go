package model

type Customer struct {
	CName    string `json:"name"`
	CTel     int64  `json:"tel"`
	CAddress string `json:"address"`
	CID      int64  `json:"id"`
	//cRegisterDate time.Time `json:"date"`
}
