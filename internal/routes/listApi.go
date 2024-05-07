package routes

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
	"strconv"

	misc "github.com/collich/go_api_learn/internal/misc"
)

var Lists []List = []List{
	{ID: 1, BorrowerName: "James", Books: Books[0]},
	{ID: 2, BorrowerName: "Jasmine", Books: Books[1]},
	{ID: 3, BorrowerName: "Johnathon", Books: Books[2]},
}

func CRUDLists(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":
		var httpStatus string
		idParam := r.URL.Path[len("/lists/"):]

		if idParam != ""{
			idParamInt, err := strconv.Atoi(idParam)
			misc.ErrorHandling(err)

			if idParamInt < 0 || idParamInt > len(Lists){
				httpStatus = misc.SetApplicationJsonHeader(w, "notfound")

				ErrorResponse := misc.ErrorResponse{
					Message: "The person is not found",
				}

				json.NewEncoder(w).Encode(ErrorResponse)
				misc.StatusOutput(httpStatus, r.URL)
				return
			}

			httpStatus = misc.SetApplicationJsonHeader(w, "ok")
			
			var foundList List
			for _, i := range Lists{
				if i.ID == int32(idParamInt){
					foundList = i
				}
			}

			json.NewEncoder(w).Encode(foundList)

			misc.StatusOutput(httpStatus, r.URL)

		} else {
			httpStatus = misc.SetApplicationJsonHeader(w, "ok")
			misc.StatusOutput(httpStatus, r.URL)
			
			json.NewEncoder(w).Encode(Lists)
		}
	case "POST":
		id := len(Lists) + 1
		var newBody List

		body, err:= io.ReadAll(r.Body)
		misc.ErrorHandling(err)

		json.Unmarshal(body, &newBody)

		newList := IfBookEmpty(newBody, int32(id))

		Lists = append(Lists, newList)

		httpStatus := misc.SetApplicationJsonHeader(w, "ok")
		json.NewEncoder(w).Encode(newList)

		misc.StatusOutput(httpStatus, r.URL)
	}

}

func IfBookEmpty(body List, id int32) List {
	var newList List
	if body.Books == (Book{}) {
		newList = List{
			ID: int32(id),
			BorrowerName: body.BorrowerName,
			Books: Book{},
		}
		return newList
	}
	
	newList = List{
		ID: int32(id),
		BorrowerName: body.BorrowerName,
		Books: body.Books,
	}
	return newList
}