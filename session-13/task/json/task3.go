package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type students struct {
	Name  string `json: "name"`
	Age   int    `json: "age"`
	Grade int    `json: "grade"`
}

func parseJSON(fileName string) []students {
	var student []students
	file, _ := os.ReadFile(fileName)

	err := json.Unmarshal(file, &student)
	if err != nil {
		fmt.Println(err)
	}

	return student
}

func prettyPrint(msg string, winners []string) {
	fmt.Println(msg)
	for _, winner := range winners {
		fmt.Println("-", winner)
	}

}

func RunTask3() {
	var winners []string
	students := parseJSON("task/json/students.json")
	for _, student := range students {
		if student.Grade < 60 {
			winners = append(winners, student.Name)
		}
	}
	prettyPrint("Students with passing grades", winners)
}
