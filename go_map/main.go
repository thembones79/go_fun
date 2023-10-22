package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#jkjkjkjkjkjj",
		"green": "#jkjkjkjkjkjj",
		"blue":  "#jkjkjkj",
	}
	colors["white"] = "#ffffff"
	delete(colors, "red")
	fmt.Print(colors)
    printMap(colors)

}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
