package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	menuitem   = make(map[string]float64)
	itemRecipe = make(map[string][]string)

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
		fmt.Println("6.) Delete ingredient for an item")
		fmt.Println("7.) List Menu Items")
		fmt.Println("8.) Update Ingredients")
		fmt.Println("9.) Save Session")
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
		} else if input == 8 {
			clear()
			UpdateIngredients()
			fmt.Printf("Press Enter.......")
			fmt.Scanln()
			clear()
		} else if input == 3 {
			clear()
			DeleteMenuItem()
			fmt.Printf("Press Enter.......")
			fmt.Scanln()
			clear()
		} else if input == 6 {
			clear()
			DeleteIngredient()
			fmt.Scanln()
			clear()
		} else if input == 2 {
			clear()
			UpdateMenuItem()
			fmt.Scanln()
			clear()
		} else if input == 9 {
			clear()
			WriteToFile()
			WriteRecipeToFile()
			fmt.Scanln()
			clear()
		}
	}
	clear()

}

func init() {
	LoadMenuItems()
	LoadRecipe()
	fmt.Scanln()
}

func LoadMenuItems() {
	//load text file if it exist
	file, err := os.Open("./files/Menuitem.txt")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	//Scan file line by line into a map
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		price, errr := strconv.ParseFloat(line[1], 64)
		if errr != nil {
			fmt.Println(errr.Error())
		} else {
			menuitem[line[0]] = price
		}

	}
	fmt.Println("Menu Items Loaded")
}

func LoadRecipe() {
	//open file if not creates
	file, err := os.Open("./files/Recipe.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := strings.Split(s.Text(), ",")
		for i := 1; i < len(line); i++ {
			itemRecipe[line[0]] = append(itemRecipe[line[0]], line[i])
		}
	}
	fmt.Println("Recipes for items have been loaded")
}

func WriteRecipeToFile() {
	//create a new file
	file, err := os.Create("./files/Recipe.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	//loop through recipes
	for item, ingredients := range itemRecipe {
		s := item + ","
		for i := 0; i < len(ingredients); i++ {
			s += ingredients[i] + ","
		}
		s = strings.TrimSuffix(s, ",")
		//write the items to a file
		_, err := fmt.Fprintln(file, s)
		if err != nil {
			fmt.Println(err.Error())
			file.Close()
		}
	}
	file.Close()
}

func WriteToFile() {
	//Create the file
	file, err := os.Create("./files/Menuitem.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	//write menu items to a file
	for item, price := range menuitem {
		s := item + "," + strconv.FormatFloat(float64(price), 'f', -1, 64)
		_, errr := fmt.Fprintln(file, s)
		if errr != nil {
			fmt.Println(errr.Error())
			file.Close()
		}

	}
	fmt.Print("written to file")
	file.Close()
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

func DeleteMenuItem() {
	fmt.Println()
	fmt.Println("Delete A Menu Item.....")
	fmt.Println("---------------------------")
	fmt.Println()
	fmt.Print("Enter the menu item you want to delete: ")
	item, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	} else {
		item = strings.Replace(item, "\n", "", -1)
		//check if menu item exists
		if _, ok := menuitem[item]; ok {
			delete(menuitem, item)
			fmt.Printf("%s has been deleted: \n", item)
			//check if item had a recipe
			if _, ok := itemRecipe[item]; ok {
				delete(itemRecipe, item)
				fmt.Printf("Recipe for %s has been deleted: \n", item)
			}

		} else {
			fmt.Println("Menu Item does not exist")
		}
	}
}

func UpdateMenuItem() {
	fmt.Println()
	fmt.Println("Updating A Menu Item.....")
	fmt.Println("---------------------------")
	fmt.Println()
	fmt.Print("Enter the menu item you wish to change: ")
	menuItemTochange, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err.Error())
		fmt.Scanln()
	}
	menuItemTochange = strings.Replace(menuItemTochange, "\n", "", -1)
	//check if menu item exists
	if _, ok := menuitem[menuItemTochange]; ok {
		fmt.Print("Enter the name you wish to change to: ")
		changedMenuItem, error := reader.ReadString('\n')
		if error != nil {
			fmt.Println("ERROR ", error.Error())
		} else {
			//Copy menu item values and delete old menu item
			v := menuitem[menuItemTochange]
			delete(menuitem, menuItemTochange)

			//add the changed menu item to map
			menuitem[changedMenuItem] = v
			fmt.Println("Menu item changed..")

			//Check if that menu item has a recipe
			if _, ok := itemRecipe[menuItemTochange]; ok {
				//Copy recipe values and delet old recipe menu item
				r := itemRecipe[menuItemTochange]
				delete(itemRecipe, menuItemTochange)

				//add changed menu item to recipe
				itemRecipe[changedMenuItem] = r
				fmt.Println("Recipe changes..")
			}

		}

	} else {
		fmt.Println("Menu item does not exist")
		fmt.Scanln()
	}
}

func DeleteIngredient() {
	fmt.Println()
	fmt.Println("Delete A Menu Item Ingredient.....")
	fmt.Println("---------------------------")
	fmt.Println()
	fmt.Printf("Enter a menu item: ")
	item, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	} else {
		item = strings.Replace(item, "\n", "", -1)
		//checking if the item exists
		if _, ok := itemRecipe[item]; ok {
			fmt.Printf("Enter the ingredient you want to delete: ")
			ingredientTodelete, error := reader.ReadString('\n')
			if error != nil {
				fmt.Println(err.Error())
			}
			ingredientTodelete = strings.Replace(ingredientTodelete, "\n", "", -1)
			for index, ingredient := range itemRecipe[item] {
				//Deleting from a slice
				if ingredientTodelete == ingredient {
					itemRecipe[item] = append(itemRecipe[item][:index], itemRecipe[item][index+1:]...)
				}
			}
		} else {
			fmt.Printf("%s does not have a recipe or does not exist \n", item)
		}
	}

}

func UpdateIngredients() {
	fmt.Println()
	fmt.Println("Updating Menu Item Ingredient.....")
	fmt.Println("---------------------------")
	fmt.Println()
	fmt.Print("Enter the Menu Item: ")
	item, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	item = strings.Replace(item, "\n", "", -1)
	changed := false
	//Find if the menu item exists
	if _, ok := itemRecipe[item]; ok {
		fmt.Print("Enter the ingredient you wish to change: ")
		oldIngredient, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		} else {

			oldIngredient = strings.Replace(oldIngredient, "\n", "", -1)
			//Finding the ingredients
			for index, ingredient := range itemRecipe[item] {
				if ingredient == oldIngredient {
					fmt.Print("Enter new ingredient: ")
					newIngredient, error := reader.ReadString('\n')
					if error != nil {
						fmt.Println(error.Error())
					}
					//Changing the ingredient
					fmt.Println("Changed From", itemRecipe[item][index])
					itemRecipe[item][index] = newIngredient
					fmt.Println("TO ..", itemRecipe[item][index])
					changed = true
					break
				}
			}
			if !changed {
				fmt.Printf("%s does not have ingredient %s \n", item, oldIngredient)
			}
		}

	} else {
		fmt.Println("This menu item does not have a recipe or does not exist")
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
	var itemPrice float64
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
	//Check if the item exists
	if _, ok := menuitem[Item]; ok {
		fmt.Print("How many Ingredients do you want to enter: ")
		_, err := fmt.Scanln(&noOfIngredients)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			//loop through ingredients and add them to map
			for i := 0; i < noOfIngredients; i++ {
				fmt.Printf("Enter Ingredient %d: ", i+1)
				ingredient, error := reader.ReadString('\n')
				if error != nil {
					fmt.Println(error.Error())
					break
				}
				//add ingredients to map
				ingredient = strings.Replace(ingredient, "\n", "", -1)
				itemRecipe[Item] = append(itemRecipe[Item], ingredient)
			}

		}

	} else {
		fmt.Println("There's no such menu item: ")
	}

}
