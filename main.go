package main

import (
	"fmt"
	"github.com/joho/godotenv"
	route2 "github.com/juliofilizzola/simulator/application/route"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, err := route.ExportJsonPositions()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stringJson[2])

}
