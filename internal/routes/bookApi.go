package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct{
	Name string `json:"Name"`
	Author string `json:"Author"`
	Desc string `json:"desc"`
	Summary string `json:"summary"`
	Price float64 `json:"Price"`
	BorrowStatus bool `json:"borrowStatus"`
} 

func GetAllBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	Books := []Book{
		{Name: "Pirates of the Caribbean", Author: "Johnny Deps", Desc: "American fantasy supernatural swashbuckler film series", Summary: "A pirate that lives their desired lifestyle", Price: 14.99, BorrowStatus: true},
		{Name: "Violent Prince", Author: "Mad Prince", Desc: "Korean Manhua", Summary: "A hero turn villian, will there be redemption for this poor soul who have lost their way of life?", Price: 49.99, BorrowStatus: false},
		{Name: "Harry Porter", Author: "Harry the Nerd", Desc: "American fantasy super natural best seller novel", Summary: "Will harry marry the ferry?", Price: 29.99, BorrowStatus: true},
	}
	httpStatus := strconv.Itoa(http.StatusAccepted)
	fmt.Printf("Status %v: Endpoint hit on %v\n", httpStatus, r.URL)
	json.NewEncoder(w).Encode(Books)
}