package main

import (
	"fmt"

	"github.com/vzx7/go-head-first/pkg/magazine"
)

func main() {
	var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)

	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	employee.City = "Moscow"
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
	fmt.Println(employee.City)
}
