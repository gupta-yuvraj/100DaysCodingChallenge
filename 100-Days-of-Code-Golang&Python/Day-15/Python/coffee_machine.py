resources = {
    "water": 300,
    "milk": 200,
    "coffee": 100,
    "money": 0,
}

menu = {
    "espresso": {
        "ingredients": {
            "water": 50,
            "coffee": 18,
        },
        "cost": 1.5,
    },
    "latte": {
        "ingredients": {
            "water": 200,
            "milk": 150,
            "coffee": 24,
        },
        "cost": 2.5,
    },
    "cappuccino": {
        "ingredients": {
            "water": 250,
            "milk": 100,
            "coffee": 24,
        },
        "cost": 3.0,
    }
}


def main():
    while True:
        input_value = input("What would you like? (espresso/latte/cappuccino):")
        if input_value == "espresso" or input_value == "latte" or input_value == "cappuccino":
            make_coffee(input_value)
            print("Here is your "+ input_value +". Enjoy!")
            continue
        elif input_value == "report":
            print("Coffee Machine Report:")
            get_report()
            continue
        elif input_value == "off":
            print("Turning off the coffee machine.")
            return 
        else:
            print("Please provide valid input.")
            continue


def get_report():
    print("Water: "+ str(resources["water"])+ "ml")
    print("Milk: "+ str(resources["milk"])+ "ml")
    print("Coffee: "+ str(resources["coffee"])+ "gm")
    print("Money: $"+ str(resources["money"]))


def make_coffee(name):
    if not check_resources(name):
        return
    
    if not process_coins(name):
        return 
    elif name == "espresso":
        resources['coffee'] -= menu["espresso"]["ingredients"]["coffee"]
        resources['water'] -=  menu["espresso"]["ingredients"]["water"]
    elif name == "latte":
        resources['coffee']-=  menu["latte"]["ingredients"]["coffee"]
        resources['water']-=  menu["latte"]["ingredients"]["water"]
        resources['milk']-=  menu["latte"]["ingredients"]["milk"]
    else:
        resources['coffee']-=  menu["cappuccino"]["ingredients"]["coffee"]
        resources['water']-=  menu["cappuccino"]["ingredients"]["water"]
        resources['milk']-=  menu["cappuccino"]["ingredients"]["milk"]

def check_resources(name):
    if not ((menu[name]["ingredients"]["coffee"] <= resources["coffee"])) :
        print("Sorry there is not enough coffee.​")
        return False
    if not ((menu[name]["ingredients"]["water"] <= resources["water"])) :
        print("Sorry there is not enough water.​")
        return False
    if name != "espresso" and not ((menu[name]["ingredients"]["milk"] <= resources["milk"])):
        print("Sorry there is not enough milk.​")
        return False
    return True

def process_coins(name):
    print("Please insert coins.")
    quarters = float(input("how many quarters?:"))
    dimes = float(input("how many dimes?:"))
    nickles = float(input("how many nickles?:"))
    pennies = float(input("how many pennies?:"))

    total_money_inserted = 0.25 * quarters + 0.1 * dimes + 0.05 * nickles + 0.01 * pennies
    
    if (total_money_inserted < menu[name]['cost']):
        print("Sorry that's not enough money. Money refunded.")
        return False
    elif (total_money_inserted > menu[name]['cost']):
        resources["money"] +=  menu[name]["cost"]
        print("Here is $"+str(round(total_money_inserted - menu[name]['cost']))+" dollars in change.")
        # We can also use print(f"{value_to_be_rounded: .2f}") instead of round funtion
        return True
    else:
        resources["money"] +=  menu[name]["cost"]
        return True

main()