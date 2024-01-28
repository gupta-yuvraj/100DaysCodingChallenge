class CofeeMaker():
    """Coffee Maker Class"""
    def __init__(self):
        self.water = 300
        self.milk = 200
        self.coffee = 100
    
    def report(self):
        print("Resources available:")
        print(f"Water: {self.water}")
        print(f"Milk: {self.milk}")
        print(f"Coffee: {self.coffee}")
    
    def is_resource_sufficient(self, drink):
        return drink.ingredients["water"] > self.water and drink.ingredients["milk"] > self.milk and drink.ingredients["coffee"] > self.coffee
    
    def make_coffee(self, order):
        """
        Parameter order: (MenuItem) The MenuItem object to make. 
        Deducts the required ingredients from the resources.
        """
        self.water -= order.ingredients["water"]
        self.coffee -= order.ingredients["coffee"]
        self.milk -= order.ingredients["milk"]

