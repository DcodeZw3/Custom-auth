package main

import (
	"fmt"
	"net/http"
	"github.com/DibyashaktiMoharana/Custom-Auth/pkg/handlers"
)



const portNumber = ":3000"


//main
func main() {
http.HandleFunc("/", handlers.Home)
http.HandleFunc("/about", handlers.About)
fmt.Printf("Starting server %v", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}


