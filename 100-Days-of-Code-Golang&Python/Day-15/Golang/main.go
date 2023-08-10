package main

import (
	"fmt"
)

var resources = map[string]float64{
	"water":  300,
	"milk":   200,
	"coffee": 100,
}

var menu = map[string]interface{}{
	"espresso": map[string]interface{}{
		"ingredients": map[string]float64{
			"water":  50,
			"coffee": 18,
		},
		"cost": 1.5,
	},

	"latte": map[string]interface{}{
		"ingredients": map[string]float64{
			"water":  200,
			"milk":   150,
			"coffee": 24,
		},
		"cost": 2.5,
	},

	"cappuccino": map[string]interface{}{
		"ingredients": map[string]float64{
			"water":  250,
			"milk":   100,
			"coffee": 24,
		},
		"cost": 3.0,
	},
}

var profit float64

func main() {
	for {
		var choice string
		fmt.Println("What would you like? (espresso/latte/cappuccino):")
		fmt.Scanln(&choice)
		if choice == "off" {
			return
		} else if choice == "report" {
			report()
			continue
		} else if choice == "espresso" || choice == "latte" || choice == "cappuccino" {
			if checkResources(choice) {
				if isTransactionSuccessful(processCoin(), choice) {
					makeCoffee(choice)
				}
			}
			continue
		} else {
			fmt.Println("Please provide valid input.")
			continue
		}
	}
}

func makeCoffee(drink string) {
	for ingredient, quantity := range menu[drink].(map[string]interface{})["ingredients"].(map[string]float64) {
		if quantity < resources[ingredient] {
			resources[ingredient] -= quantity
		}
	}
	profit += menu[drink].(map[string]interface{})["cost"].(float64)
	fmt.Printf("Here is your %s. Enjoy!\n", drink)
}

func isTransactionSuccessful(moneyInserted float64, drink string) bool {
	if moneyInserted < menu[drink].(map[string]interface{})["cost"].(float64) {
		fmt.Println("Sorry that's not enough money. Money refunded.")
		return false
	} else if moneyInserted > menu[drink].(map[string]interface{})["cost"].(float64) {
		fmt.Printf("Here is $%.2f dollars in change.\n", (moneyInserted - menu[drink].(map[string]interface{})["cost"].(float64)))
		return true
	}
	return true
}

func processCoin() (total float64) {
	coins := []struct {
		name  string
		value float64
	}{
		{"quarters", 0.25},
		{"dimes", 0.10},
		{"nickles", 0.05},
		{"pennies", 0.01},
	}
	for _, coin := range coins {
		var count int
		fmt.Printf("how many %s?: ", coin.name)
		fmt.Scan(&count)
		total += (coin.value * float64(count))
	}
	return total
}

func checkResources(drink string) bool {
	for ingredient, quantity := range menu[drink].(map[string]interface{})["ingredients"].(map[string]float64) {
		if quantity > resources[ingredient] {
			fmt.Printf("Sorry there is not enough %s.\n", ingredient)
			return false
		}
	}
	return true
}

func report() {
	fmt.Printf("Water: %.2fml\n", resources["water"])
	fmt.Printf("Milk: %.2fml\n", resources["milk"])
	fmt.Printf("Coffee: %.2fgm\n", resources["coffee"])
	fmt.Printf("Money: $%.2f\n", profit)
}
