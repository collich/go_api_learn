package misc

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
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
	var ListOfIP []string

	CMD := exec.Command("hostname", "-i")
	Output, err := CMD.Output()
	ErrorHandling(err)
	
	r, err := regexp.Compile("")
	ErrorHandling(err)
	
	Results := strings.Split(string(Output), " ")

	for _, i := range Results{
		if match := r.Match([]byte(i)); match {
			// fmt.Printf("Site is up:", i, PORT)
			ListOfIP = append(ListOfIP, i)
		}
	}

	if len(ListOfIP) != 0{
		fmt.Println("Your site is up: ")
		for _, i := range ListOfIP{
			fmt.Printf("http://%v%v\n", i,PORT)
		}
	} else {
		fmt.Println("Couldn't find any server")
		return
	}
}

