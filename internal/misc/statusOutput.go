package misc

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
)

var statusMap = map[string]int{
	"ok": http.StatusOK,
	"notfound": http.StatusNotFound,
	"accept": http.StatusAccepted,
}

// Status Output Method
func StatusOutput(httpStatus string, r *url.URL){
	fmt.Printf("Status %v: Endpoint hit on %v\n", httpStatus, r)
}

// Setting Site Header
func SetApplicationJsonHeader(w http.ResponseWriter, statusPass string) string {
	httpStatus := statusMap[statusPass]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	
	return strconv.Itoa(httpStatus)
}

// Error handling functions
func ErrorHandling(err error) {
	if err != nil {
		log.Fatalf("Got Error: %v", err)
	}
}

// One time use function that prints out the host when start
func URLCLIPrintOut(PORT string)  {

}

