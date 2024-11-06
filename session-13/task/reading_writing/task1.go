package reading_writing

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type person struct {
	name  string
	age   int
	grade int
}

func parseCSV(fileName string) []person {
	people := make([]person, 0)
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error ocurred when reading file:", err)
	}
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines[1:] {
		words := bytes.Split(line, []byte(","))
		age, _ := strconv.Atoi(string(words[1]))
		grade, _ := strconv.Atoi(string(words[2]))
		people = append(people, person{
			name:  string(words[0]),
			age:   age,
			grade: grade,
		})

	}
	return people
}

func prettyPrint(msg string, winners []string) {
	fmt.Println(msg)
	for _, winner := range winners {
		fmt.Println("-", winner)
	}

}

func RunTask1() {
	var winners []string
	people := parseCSV("task/reading_writing/data.csv")
	for _, person := range people {
		if person.grade > 60 {
			winners = append(winners, person.name)
		}
	}
	prettyPrint("Students with passing grades", winners)
}
