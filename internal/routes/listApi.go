package routes

import (
	"encoding/json"
	"net/http"

	"github.com/collich/go_api_learn/internal/misc"
)

var Lists []List

func GetLists(w http.ResponseWriter, r *http.Request)  {

	httpStatus := misc.SetApplicationJsonHeader(w, "accept")

	Lists = []List{
		{ID: 1, BorrowerName: "James", Books: Books[0]},
		{ID: 2, BorrowerName: "Jasmine", Books: Books[1]},
		{ID: 3, BorrowerName: "Johnathon", Books: Books[2]},
	}

	misc.StatusOutput(httpStatus, r.URL)
	
	json.NewEncoder(w).Encode(Lists)
}
