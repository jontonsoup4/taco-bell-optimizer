# taco-bell-optimizer
An API server that returns information on items from the Taco Bell menu.

# Usage:
```
make install
make build
make
```

# Available Endpoints:
```
api/menu/{type}                returns all of type
api/sort/{type}/{property}     returns type sorted by property
api/value/{type}/{property}    returns type sorted by (cost / property)
    
?reverse=true                  reverses order
    
type = ["food", "drinks", "all"]
property = ["Cost", "Calories", "CaloriesFromFat", "TotalFat", "SaturatedFat",
          "TransFat", "Cholesterol", "Sodium", "Carbohydrates", "DietaryFiber",
          "Sugars", "Protein"]
```

# Example Calls:
```
api/menu/drinks
api/menu/combos
api/sort/all/sodium
api/value/food/calories?reverse=true
```

# Example Output:
```
// url: localhost:8000/api/value/food/calories
[
    {
        name: "Spicy Tostada",
        cost: 1,
        calories: 440,
        caloriesFromFat: 210,
        totalFat: 24,
        saturatedFat: 4,
        transFat: 0,
        cholesterol: 10,
        sodium: 640,
        carbohydrates: 46,
        dietaryFiber: 5,
        sugars: 2,
        protein: 10
    },
    {
        name: "Beefy Fritos® Burrito",
        cost: 1,
        calories: 430,
        caloriesFromFat: 160,
        totalFat: 18,
        saturatedFat: 4,
        transFat: 0,
        cholesterol: 20,
        sodium: 1,
        carbohydrates: 55,
        dietaryFiber: 4,
        sugars: 3,
        protein: 13
    },
    {
        name: "Cheesy Bean & Rice Burrito",
        cost: 1,
        calories: 420,
        caloriesFromFat: 160,
        totalFat: 17,
        saturatedFat: 3,
        transFat: 0,
        cholesterol: 5,
        sodium: 920,
        carbohydrates: 55,
        dietaryFiber: 6,
        sugars: 4,
        protein: 11
    }
]
```

## ToDo:
* Add more combos
* Limit returned items
* Return list of items with the best value based on attribute given a dollar amount (float)
