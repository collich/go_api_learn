package misc

import (
	"fmt"
	"net/url"
)

func StatusOutput(httpStatus string, r *url.URL){
	fmt.Printf("Status %v: Endpoint hit on %v\n", httpStatus, r)
}