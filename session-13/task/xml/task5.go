package xml

import (
	"encoding/xml"
	"fmt"
	"os"
)

type employees struct {
	Employees []Employee `xml:"employee"`
}

type Employee struct {
	Name     string `xml:"name"`
	Position string `xml:"position"`
	Salary   int    `xml:"salary"`
}

func parseXML(fileName string) employees {
	var employees employees
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	err = xml.Unmarshal(file, &employees)
	if err != nil {
		fmt.Println(err)
	}
	return employees
}

func findWinner(employees employees) [][]string {
	var winners [][]string
	for _, employee := range employees.Employees {

		if employee.Salary > 50000 {
			winners = append(winners, []string{employee.Name, employee.Position})
		}
	}
	return winners
}
func prettyPrint(msg string, winners [][]string) {
	fmt.Println(msg)
	for _, winner := range winners {
		fmt.Printf("- %s, %s\n", string(winner[0]), string(winner[1]))
	}

}

func RunTask5() {
	employees := parseXML("task/xml/employees.xml")
	winners := findWinner(employees)
	prettyPrint("Employees with Salary above 50000:", winners)

}
