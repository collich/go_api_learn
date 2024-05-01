package routes

type Book struct{
	ID int64 `json:"ID"`
	Name string `json:"Name"`
	Author string `json:"Author"`
	Desc string `json:"desc"`
	Summary string `json:"summary"`
	Price float64 `json:"Price"`
	BorrowStatus bool `json:"borrowStatus"`
}