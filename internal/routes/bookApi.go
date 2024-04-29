package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/collich/go_api_learn/internal/misc"
)

type Book struct{
	ID int64 `json:"ID"`
	Name string `json:"Name"`
	Author string `json:"Author"`
	Desc string `json:"desc"`
	Summary string `json:"summary"`
	Price float64 `json:"Price"`
	BorrowStatus bool `json:"borrowStatus"`
}

type ErrorResponse struct {
	Message string `json:"Message"`
}

var Books []Book = []Book{
	{ID: 1, Name: "Pirates of the Caribbean", Author: "Johnny Deps", Desc: "American fantasy supernatural swashbuckler film series", Summary: "A pirate that lives their desired lifestyle", Price: 14.99, BorrowStatus: true},
	{ID: 2, Name: "Violent Prince", Author: "Mad Prince", Desc: "Korean Manhua", Summary: "A hero turn villian, will there be redemption for this poor soul who have lost their way of life?", Price: 49.99, BorrowStatus: false},
	{ID: 3, Name: "Harry Porter", Author: "Harry the Nerd", Desc: "American fantasy super natural best seller novel", Summary: "Will harry marry the ferry?", Price: 29.99, BorrowStatus: true},
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var httpStatus string
	idParam := r.URL.Path[len("/books/"):]
	
	if idParam != ""{
		idParam_int, err := strconv.Atoi(idParam)
		ErrorHandling(err)

		if idParam_int < 1 || idParam_int > len(Books){
			errorResponse := ErrorResponse{Message: "Book Not Found"}
			w.WriteHeader(http.StatusNotFound)

			httpStatus = strconv.Itoa(http.StatusNotFound)
			json.NewEncoder(w).Encode(errorResponse)
			// Misc Method to do status output
			misc.StatusOutput(httpStatus, r.URL)
			return
		}
		
		foundBook := Books[idParam_int - 1]
		
		json.NewEncoder(w).Encode(foundBook)		
		httpStatus = strconv.Itoa(http.StatusOK)
		
		// Misc Method to do status output
		misc.StatusOutput(httpStatus, r.URL)
	} else {
		httpStatus = strconv.Itoa(http.StatusOK)
		json.NewEncoder(w).Encode(Books)

		// Misc Method to do status output
		misc.StatusOutput(httpStatus, r.URL)
	}

}

func ErrorHandling(err error) {
	if err != nil {
		log.Fatalf("Got Error %v", err)
	}
}