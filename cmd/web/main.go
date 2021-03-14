package main

import (
	"fmt"
	"net/http"

	"github.com/rosaCarolSampaio/qaAnalytics/pkg/handlers"
)

const portWebServer = ":8090"


func main(){
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Stating application on port %s", portWebServer )
	_ = http.ListenAndServe(portWebServer, nil)
}

