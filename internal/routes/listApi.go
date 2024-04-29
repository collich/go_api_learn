package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type List struct {
	ID int32 `json:"ID"`
	BorrowerName string `json:"BorrowerName"`
	Books Book `json:"Books"`
}

var Lists []List

func GetLists(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var httpStatus string

	Lists = []List{
		{ID: 1, BorrowerName: "James", Books: Books[0]},
		{ID: 2, BorrowerName: "Jasmine", Books: Books[1]},
		{ID: 3, BorrowerName: "Johnathon", Books: Books[2]},
	}

	httpStatus = strconv.Itoa(http.StatusAccepted)
	fmt.Println(httpStatus)
	json.NewEncoder(w).Encode(Lists)
}