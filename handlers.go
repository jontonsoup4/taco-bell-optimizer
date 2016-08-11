package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oleiade/reflections"
)

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.Header().Set("Access-Control-Allow-Methods", "GET");
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	filename := mux.Vars(r)["type"];
	output, err := LoadJSON(w, filename);
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(output);
}

func SortHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.Header().Set("Access-Control-Allow-Methods", "GET");
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	vars := mux.Vars(r);
	items, err := LoadJSON(w, vars["type"]);
	if err != nil {
		return
	}
	propertyName := getPropertyName(strings.ToLower(vars["property"]));
	if propertyName == ""{
		json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: fmt.Sprintf("No filter by that name. Try: %q", SortByOptions),
		});
		return
	}
	stringQueries := r.URL.Query();
	reverse := stringQueries.Get("reverse");
	var sorted Items;
	if reverse == "true" {
		sorted = reverseSortProperties(propertyName, items);
	} else {
		sorted = sortProperties(propertyName, items);
	}
	json.NewEncoder(w).Encode(sorted);
}

func ValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.Header().Set("Access-Control-Allow-Methods", "GET");
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	vars := mux.Vars(r);
	items, err := LoadJSON(w, vars["type"]);
	if err != nil {
		return
	}
	itemsMap := make(map[float64]Item);
	keys := make([]float64, len(items));
	propertyName := getPropertyName(strings.ToLower(vars["property"]));
	if propertyName == ""{
		json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: fmt.Sprintf("No filter by that name. Try: %q", SortByOptions),
		});
		return
	}
	for index, item := range items {
		propVal, _ := reflections.GetField(item, propertyName);
		var value float64;
		switch propVal.(type) {
		case int:
			value = item.Cost / float64(propVal.(int));
		case float64:
			value = item.Cost / propVal.(float64);
		}
		itemsMap[value] = item;
		keys[index] = value;
	}
	stringQueries := r.URL.Query();
	reverse := stringQueries.Get("reverse");
	if reverse == "true" {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)));
	} else {
		sort.Float64s(keys);
	}
	returnItems := make([]Item, len(keys))
	for index, item := range keys {
		returnItems[index] = itemsMap[item];
	}
	json.NewEncoder(w).Encode(returnItems);
}


func MoneyValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.Header().Set("Access-Control-Allow-Methods", "GET");
	w.Header().Set("Content-Type", "application/json; charset=UTF-8");
	vars := mux.Vars(r);
	
	propertyName := getPropertyName(strings.ToLower(vars["property"]));
	if propertyName == ""{
		json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: fmt.Sprintf("No filter by that name. Try: %q", SortByOptions),
		});
		return
	}

	money := vars["money"];
	if money == ""{
		json.NewEncoder(w).Encode(jsonErr{
			Code: http.StatusNotFound,
			Text: fmt.Sprintf("Money amount missing. Try: '5' for 5 dollars"),
		});
		return
	}
	money_parsed, err := strconv.ParseFloat(money, 64)

	items, err := LoadJSON(w, vars["type"]);
	if err != nil {
		return
	}
	itemsMap := make(map[float64]Item);
	keys := make([]float64, len(items));

	skipped_keys := len(items)
	for index, item := range items {
		propVal, _ := reflections.GetField(item, propertyName);
		var value float64;
		switch propVal.(type) {
			case int:
				value = (item.Cost / float64(propVal.(int))) / money_parsed;
			case float64:
				value = (item.Cost / propVal.(float64)) / money_parsed;
		}

		if (money_parsed >= item.Cost){
			itemsMap[value] = item;
			keys[index] = value;
		} else {
			keys[index] = 0;
			skipped_keys--;
		}
	}


	stringQueries := r.URL.Query();
	reverse := stringQueries.Get("reverse");
	if reverse == "true" {
		sort.Sort(sort.Reverse(sort.Float64Slice(keys)));
	} else {
		sort.Float64s(keys);
	}
	returnItems := make([]Item, skipped_keys)
	current_index := 0
	for index, item := range keys {
		if (keys[index] != 0){
			returnItems[current_index] = itemsMap[item];
			current_index++;
		}
	}
	
	json.NewEncoder(w).Encode(returnItems);
}