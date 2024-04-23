package routes

import (
	"fmt"
	"net/http"
)
	
var Books []Book

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page Welcome")
	fmt.Printf("Status %v: Endpoint hit %v\n", http.StatusAccepted, r.URL)
}
