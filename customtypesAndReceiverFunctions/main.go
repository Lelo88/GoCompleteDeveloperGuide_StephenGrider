package main

import "fmt"

type color string

func (c color) describe(description string) string {
	return string(c) + " " + description
}

func (c color) count(times  int) string {
	 return string(c) + " appears " + fmt.Sprint(times) + " times"
}

func main() {
	newColor := color("red")
	newColor2 := color("blue")
	fmt.Println(newColor.describe("is an awesome color"))
	fmt.Println(newColor.count(3))

	fmt.Println(newColor2.describe("is a nice color"))

	fmt.Println(newColor2.count(2))
}

