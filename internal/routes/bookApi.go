package routes

// package bookroutes

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
	"strconv"

	misc "github.com/collich/go_api_learn/internal/misc"
)

var Books []Book = []Book{
	{ID: 1, Name: "Pirates of the Caribbean", Author: "Johnny Deps", Desc: "American fantasy supernatural swashbuckler film series", Summary: "A pirate that lives their desired lifestyle", Price: 14.99, BorrowStatus: true},
	{ID: 2, Name: "Violent Prince", Author: "Mad Prince", Desc: "Korean Manhua", Summary: "A hero turn villian, will there be redemption for this poor soul who have lost their way of life?", Price: 49.99, BorrowStatus: false},
	{ID: 3, Name: "Harry Porter", Author: "Harry the Nerd", Desc: "American fantasy super natural best seller novel", Summary: "Will harry marry the ferry?", Price: 29.99, BorrowStatus: true},
}

func CRUDBooks(w http.ResponseWriter, r *http.Request) {
	var httpStatus string
	idParam := r.URL.Path[len("/books/"):]
	

	switch r.Method {
	case "GET":
		if idParam != ""{
			idParamInt, err := strconv.Atoi(idParam)
			misc.ErrorHandling(err)
	
			if idParamInt < 1 || idParamInt > len(Books){
				errorResponse := misc.ErrorResponse{Message: "Book Not Found"}
				httpStatus = misc.SetApplicationJsonHeader(w, "notfound")
				json.NewEncoder(w).Encode(errorResponse)
	
				misc.StatusOutput(httpStatus, r.URL)
				return
			}
			
			// foundBook := Books[idParamInt - 1]
			var foundBook Book
			for _, i := range Books {
				if i.ID == int64(idParamInt){
					foundBook = i
				}
			}
			
			json.NewEncoder(w).Encode(foundBook)		
			httpStatus = misc.SetApplicationJsonHeader(w, "ok")
			
			// Misc Method to do status output
			misc.StatusOutput(httpStatus, r.URL)
		} else {
	
			httpStatus = misc.SetApplicationJsonHeader(w, "ok")
			json.NewEncoder(w).Encode(Books)
	
			// Misc Method to do status output
			misc.StatusOutput(httpStatus, r.URL)
		}	
	case "POST":
		var book Book
		body, err := io.ReadAll(r.Body)
		misc.ErrorHandling(err)
		
		err = json.Unmarshal(body, &book)
		misc.ErrorHandling(err)

		PostBookMutation(&book)
		Books = append(Books, book)

		httpStatus = misc.SetApplicationJsonHeader(w, "accept")

		json.NewEncoder(w).Encode(book)
		misc.StatusOutput(httpStatus, r.URL)
	case "DELETE":
		idParams := r.URL.Path[len("/books/"):]
		idParams_int, err := strconv.Atoi(idParams)
		misc.ErrorHandling(err)

		for i, k := range Books {
			if k.ID == int64(idParams_int) {
				Books = append(Books[:i], Books[i + 1: ]...)
			}
		}
		httpStatus := misc.SetApplicationJsonHeader(w, "ok")
		json.NewEncoder(w).Encode(Books)
		misc.StatusOutput(httpStatus, r.URL)
	}
}

func PostBookMutation(book *Book) {
	*book = Book{
		ID: int64(len(Books) + 1),
		Name: book.Name,
		Author: book.Author,
		Desc: book.Desc,
		Summary: book.Summary,
		Price: book.Price,
		BorrowStatus: false,
	}
}