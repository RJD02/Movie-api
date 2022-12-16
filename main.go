package main

import (
	"fmt"
	"net/http"

	router "github.com/RJD02/mongoapi/routes"
)

func main() {
	r := router.Router()
	fmt.Println("MongoDB API")
	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", r)
}
