package controller

import (
	"log"
	"net/http"

	"github.com/collich/go_api_learn/internal/routes"
)

func HandleReqs()  {
	PORT := ":8080"
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/books/", routes.CRUDBooks)
	http.HandleFunc("/lists/", routes.CRUDLists)
	log.Fatal(http.ListenAndServe(PORT, nil))
}