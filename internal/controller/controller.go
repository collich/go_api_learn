package controller

import (
	"log"
	"net/http"

	"github.com/collich/go_api_learn/internal/routes"
)

func HandleReqs()  {
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/books/", routes.CRUDBooks)
	http.HandleFunc("/name/", routes.GetLists)
	log.Fatal(http.ListenAndServe(":8080", nil))
}