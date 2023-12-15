// maps is similar to dicts in python
// all the keys and values in the map must be of same type but not necessarily together
package main

import "fmt"

func main() {
	colors := map[string]string{ //i.e colrs is of type map, [string] inside indicates all the keys are of type string and string outside indicates all the values here are string type
		"red": "xyz",
		"blue": "xyz",
	}
	//method:2 to intialise 
	// var colors map[string]string //initialises to empty strings
	//or
	// colors := make(map[string]string)
	// // assign values after intialiseing
	// colors["white"] ="hbhj"
	// // delete key
	// delete(colors, "white")
	// fmt.Println(colors)
	printMap(colors)
}


func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}