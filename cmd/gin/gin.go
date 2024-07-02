package main

import (
	"log"

	dateTimeGin "github.com/codescalersinternships/DateTime-Server-Basem-Elgalfy/pkg/gin"
)

func main() {

	log.Println("Starting Gin server on port 8080...")
	log.Fatal(dateTimeGin.StartGinServer())

}
