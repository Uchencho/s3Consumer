package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Uchencho/s3Consumer/internal/app"
)

func main() {

	port := os.Getenv("PORT")

	a := app.New()

	log.Println(fmt.Sprintf("Starting server on address: %s", port))
	log.Println(http.ListenAndServe(fmt.Sprintf(":%s", port), a.Handler()))
}
