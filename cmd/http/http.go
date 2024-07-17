package main

import (
	"log"

	dateTimeHttp "github.com/codescalersinternships/DateTime-Server-Basem-Elgalfy/pkg/http"
)

func main() {

	log.Println("Starting server on port 8080...")
	log.Fatal(dateTimeHttp.StartServer())

}
