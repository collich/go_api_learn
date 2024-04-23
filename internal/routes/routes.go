package routes

import (
	"fmt"
	"log"
	"net/http"
)



func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page Welcome")
	fmt.Println("Status 200: Endpoint hit 'Home Page'")
}

func HandleReqs() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}