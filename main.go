package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arturgumerob/postgresdbAPI/router"
	_ "github.com/lib/pq"
)

func main() {
	r := router.Router()
	fmt.Println("Start server on the port 8080..")

	log.Fatal(http.ListenAndServe(":8080", r))
}
