package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	menuitem = map[string]float32{"Cheesy Broccoli Soup in a Bread Bowl": 50.00, "Bacon Cheddar Potato Skins": 60.00,
		"Sour Cream-Lemon Pie": 45.00}

	itemRecipe = map[string][]string{"Cheesy Broccoli Soup in a Bread Bowl": {"1/4 cup butter, cubed",
		"1/2 cup chopped onion", "2 garlic cloves, minced", "4 cups fresh broccoli florets (about 8 ounces)", "1 large carrot, finely chopped",
		"3 cups chicken stock", "2 cups half-and-half cream", "2 bay leaves", "1/2 teaspoon salt", "1/4 teaspoon ground nutmeg",
		"1/4 teaspoon pepper", "1/4 cup cornstarch", "1/4 cup water or additional chicken stock", "2-1/2 cups shredded cheddar cheese",
		"6 small round bread loaves (about 8 ounces each), optional", "Optional toppings: Crumbled cooked bacon, additional shredded cheddar cheese, ground nutmeg and pepper"},
		"Bacon Cheddar Potato Skins": {"4 large baking potatoes, baked", "3 tablespoons canola oil", "1 tablespoon grated Parmesan cheese", "1/2 teaspoon salt", "1/4 teaspoon garlic powder",
			"1/4 teaspoon paprika", "1/8 teaspoon peppe", "8 bacon strips, cooked and crumbled", "1-1/2 cups shredded cheddar cheese",
			"1/2 cup sour cream", "4 green onions, sliced"}}

	reader = bufio.NewReader(os.Stdout)
)

func main() {
	fmt.Println("==================================")
	fmt.Println("Start Building Your Restuarant Menu")
	fmt.Println("===================================")
	fmt.Println()

	loop := true
	var input int

	for loop {
		fmt.Println("Select an item below by select the number of the item")
		fmt.Println("-----------------------------------------------------")
		fmt.Println()

		fmt.Println("1.) Add a menu item")
		fmt.Println("2.) Update a menu item")
		fmt.Println("3.) delete a menu item")
		fmt.Println("4.) Add ingredients for an item ")
		fmt.Println("5.) List ingredients for an item ")
		fmt.Println("6.) Delete ingredients for an item")
		fmt.Println("7.) List Menu Items")
		fmt.Println("11.) Exit")
		fmt.Println()
		fmt.Print("Enter a selection: ")

		_, err := fmt.Scan(&input)
		if err != nil {
			clear()
			fmt.Printf("Error %v \n", err.Error())
			fmt.Println()

		} else if input == 11 {
			loop = false
		} else if input == 1 {
			clear()
			AddMenuItem()
			ListMenuItems()
			fmt.Printf("Press Enter.......")
			fmt.Scanln()
			clear()
		} else if input == 5 {
			clear()
			ListIngredients()
			fmt.Printf("Press Enter.......")
			fmt.Scanln()
			clear()
		} else if input == 4 {
			clear()
			AddIngredient()
			ListIngredients()
			fmt.Printf("Press Enter.......")
			fmt.Scanln()
			clear()
		} else if input == 7 {
			clear()
			ListMenuItems()
			fmt.Printf("Press Enter.......")
			fmt.Scanln()
			clear()
		}
	}
	clear()

}

func init() {
	//Loads Data from a File

}

func ListIngredients() {
	fmt.Println()
	fmt.Println("Listing Ingredients.....")
	fmt.Println("---------------------------")
	fmt.Println()

	for item, ingredient := range itemRecipe {
		fmt.Println("Recipe For: ", item)
		fmt.Println("------------------")
		for i := 0; i < len(ingredient); i++ {
			fmt.Printf("%d) %s \n", i+1, ingredient[i])
		}
		fmt.Println()
	}
}

func clear() {
	//Clears the command terminal
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ListMenuItems() {
	fmt.Println()
	fmt.Println("Listing Menu.....")
	fmt.Println("---------------------------")
	fmt.Println()
	i := 1
	for item, price := range menuitem {

		fmt.Printf("%d)%s\t R%v \n", i, item, price)
		i++
	}
	fmt.Println()
}

func AddMenuItem() {
	fmt.Println()
	fmt.Println("Adding a new menu item")
	fmt.Println("----------------------")
	fmt.Println()
	fmt.Print("Enter item name: ")
	//_, err := fmt.Scanf("Mac and cheese", &itemName)

	itemName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	itemName = strings.Replace(itemName, "\n", "", -1)
	var itemPrice float32
	fmt.Print("Enter item Price: ")
	_, err2 := fmt.Scan(&itemPrice)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	//Add item to menuItems
	menuitem[itemName] = itemPrice
	fmt.Println()
}

func AddIngredient() {
	fmt.Println()
	fmt.Println("Adding an ingredient to menu item")
	fmt.Println("---------------------------------")
	fmt.Println()

	fmt.Print("Enter menu item: ")
	Item, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	Item = strings.Replace(Item, "\n", "", -1)

	var noOfIngredients int
	if _, ok := menuitem[Item]; ok {
		fmt.Print("How many Ingredients do you want to enter: ")
		_, err := fmt.Scanln(&noOfIngredients)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			for i := 0; i < noOfIngredients; i++ {
				fmt.Printf("Enter Ingredient %d: ", i)
				ingredient, error := reader.ReadString('\n')
				if error != nil {
					fmt.Println(error.Error())
					break
				}
				ingredient = strings.Replace(ingredient, "\n", "", -1)
				itemRecipe[Item] = append(itemRecipe[Item], ingredient)
			}

		}

	} else {
		fmt.Println("There's no such menu item: ")
	}

}
