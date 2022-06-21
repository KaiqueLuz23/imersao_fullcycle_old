package main

import (
	"fmt"
	route2 "kaiqueluz23/go-simulator/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringsjon, _ := route.ExportJsonPositions()
	fmt.Println(stringsjon[0])
}
