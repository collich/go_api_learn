package controller

import (
	"log"
	"net/http"

	misc "github.com/collich/go_api_learn/internal/misc"
	routes "github.com/collich/go_api_learn/internal/routes"
)

func HandleReqs()  {
	PORT := ":8080"

	misc.URLCLIPrintOut(PORT)
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/books/", routes.CRUDBooks)
	http.HandleFunc("/lists/", routes.CRUDLists)
	log.Fatal(http.ListenAndServe(PORT, nil))
}