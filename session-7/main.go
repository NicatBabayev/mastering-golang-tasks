package main

import (
	"fmt"
	"session-7/model/embedding"
	"session-7/model/methods"
	"session-7/model/structs"
)

// #TODO Write required comments
func main() {
	// Task 1
	book := structs.Book{
		Title:  "The Go Programming Language",
		Author: "Alan A. A. Donovan",
		Pages:  380,
	}
	// Task 2
	rectangle := structs.Rectangle{
		Width:  10.5,
		Height: 5.0,
	}
	// Task 3
	circle := methods.Circle{
		Radius: 7.0,
	}
	// Task 4
	std1 := methods.Student{
		Name:   "Ali",
		Grades: []int{80, 85, 90, 87, 83},
	}
	std2 := methods.Student{
		Name:   "Vali",
		Grades: []int{88, 92, 91, 89, 90},
	}
	// Task 5
	employee := embedding.Employee{
		EmployeeID: 12345,
		Position:   "Software Engineer",
		Person: embedding.Person{
			FirstName: "Ali",
			LastName:  "Aliyev",
		},
	}
	// Task 6
	car := embedding.Car{
		Vehicle: embedding.Vehicle{
			Make:  "Kia",
			Model: "K5",
			Year:  2022,
		},
		FuelType: "Gasoline",
	}
	// Task 1
	fmt.Println(" // Task 1")
	fmt.Println(book)

	// Task 2
	fmt.Println(" // Task 2")
	fmt.Println(rectangle)

	// Task 3
	fmt.Println(" // Task 3")
	fmt.Printf(" Circle Radius: %.1f\n Area: %.3f\n", circle.Radius, circle.Area())

	fmt.Println()

	// Task 4
	fmt.Println(" // Task 4")
	printHigherAverage(std1, std2)

	fmt.Println()

	// Task 5
	fmt.Println(" // Task 5")
	fmt.Printf(" Name: %s %s\n", employee.Person.FirstName, employee.Person.LastName)
	fmt.Printf(" Employee ID: %d\n", employee.EmployeeID)
	fmt.Printf(" Position: %s\n", employee.Position)

	fmt.Println()

	// Task 6
	fmt.Println(" // Task 6")
	fmt.Printf(" Make: %s\n", car.Vehicle.Make)
	fmt.Printf(" Model: %s\n", car.Vehicle.Model)
	fmt.Printf(" Year: %d\n", car.Vehicle.Year)
	fmt.Printf(" Fuel Type: %s\n", car.FuelType)
}

// Task 4
func printHigherAverage(std1, std2 methods.Student) {
	var res string

	if std1.Average() > std2.Average() {
		res = std1.Name

	} else {
		res = std2.Name
	}

	fmt.Printf(" Student 1: %s, Average Grade: %.0f\n", std1.Name, std1.Average())
	fmt.Printf(" Student 2: %s, Average Grade: %.0f\n", std2.Name, std2.Average())
	fmt.Printf(" %s has a higher average grade.\n", res)

}
