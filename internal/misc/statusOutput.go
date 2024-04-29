package misc

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

var statusMap = map[string]int{
	"ok": http.StatusOK,
	"notfound": http.StatusNotFound,
	"accept": http.StatusAccepted,
}

func StatusOutput(httpStatus string, r *url.URL){
	fmt.Printf("Status %v: Endpoint hit on %v\n", httpStatus, r)
}

func SetApplicationJsonHeader(w http.ResponseWriter, statusPass string) string {
	httpStatus := statusMap[statusPass]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	
	return strconv.Itoa(httpStatus)
}