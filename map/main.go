package main

import (
	"fmt"

)

func main() {


	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

//  map implement 
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white" : "#ffffff",
	}

	// delete(colors,"white")
	fmt.Println((colors))
}