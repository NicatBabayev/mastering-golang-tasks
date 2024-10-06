package main

import "fmt"

func main() {
	countries := map[string]string{
		"Azerbaijan": "Baku",
		"USA":        "Washington",
		"Germany":    "Berlin",
		"Japan":      "Tokyo",
	}
	countries["Turkey"] = "Ankara"
	countries["Italy"] = "Rome"
	for k := range countries {
		fmt.Printf("Capital of %s is %s\n", k, getCountryName(k, countries))
	}
	fmt.Println(getCountryName("France", countries))
}

func getCountryName(s string, countries map[string]string) string {
	var Red = "\033[31m"
	var Reset = "\033[0m"
	value, ok := countries[s]
	if ok {
		return value
	} else {
		return Red + "`Capital of " + s + " is not found`" + Reset
	}
}
