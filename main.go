package main

import (
	"fmt"
	route2 "github.com/juliofilizzola/simulator/application/route"
)

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
