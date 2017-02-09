package main

type Item struct {
	Name            string  `json:"name"`
	Cost            float64 `json:"cost"`
	Calories        int     `json:"calories"`
	CaloriesFromFat int     `json:"caloriesFromFat"`
	TotalFat        int     `json:"totalFat"`
	SaturatedFat    int     `json:"saturatedFat"`
	TransFat        int     `json:"transFat"`
	Cholesterol     int     `json:"cholesterol"`
	Sodium          int     `json:"sodium"`
	Carbohydrates   int     `json:"carbohydrates"`
	DietaryFiber    int     `json:"dietaryFiber"`
	Sugars          int     `json:"sugars"`
	Protein         int     `json:"protein"`
	Updated         string  `json:"updated"`
}

type Items []Item
