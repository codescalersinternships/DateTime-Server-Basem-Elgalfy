package main

import (
	"log"

	dateTime "github.com/codescalersinternships/DateTime-Server-Basem-Elgalfy/pkg"
)

func main() {

	log.Println("Starting server on port 8080...")
	log.Fatal(dateTime.StartServer())

}
