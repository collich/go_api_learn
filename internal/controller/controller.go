package controller

import (
	"log"
	"net/http"

	"github.com/collich/go_api_learn/internal/routes"
)

func HandleReqs()  {
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/books", routes.GetAllBooks)
	log.Fatal(http.ListenAndServe(":8080", nil))
}