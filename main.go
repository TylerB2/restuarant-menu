package main

import "fmt"

func main() {
	fmt.Println("==================================")
	fmt.Println("Start Building Your Restuarant Menu")
	fmt.Println("===================================")
	fmt.Println()

	fmt.Println("Select an item below by select the number of the item")
	fmt.Println("-----------------------------------------------------")
	fmt.Println()
	loop := true
	var input int
	for loop {

		fmt.Println("1.) Add a menu item")
		fmt.Println("2.) Update a menu item")
		fmt.Println("3.) delete a menu item")
		fmt.Println("4.) Add ingredients for an item ")
		fmt.Println("5.) List ingredients for an item ")
		fmt.Println("6.) Delete ingredients for an item")
		fmt.Println()
		fmt.Print("Enter a selection: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			//fmt.Println("\033[H\033[2j")
			fmt.Printf("Error %v \n", err.Error())
			fmt.Println()

		} else {
			loop = false
		}
	}
}
