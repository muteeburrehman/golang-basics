package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}
func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

// Functions
func add(a int, b int) int {
	return a + b
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Structures
type Person struct {
	Name string
	Age  int
}

func main() {

	var age int = 25
	name := "John"
	const PI = 3.14
	fmt.Printf("Name: %s, Age: %d, PI: %.2f\n", name, age, PI)
	// Conditional
	num := 10
	if num%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("Iteration", i)
	}
	// Switch
	day := "Thursday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("End of workweek")
	default:
		fmt.Println("Middle of the week")
	}
	fmt.Println("Hello, World")

	result := add(5, 3)
	fmt.Println("Sum:", result)

	// Arrays

	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20

	fmt.Println(numbers)
	// Slicing
	Num := []int{1, 2, 3}
	Num = append(Num, 4)

	fmt.Println("Value after slicing is ", Num)

	//  Maps
	fruits := map[string]int{
		"Apple":  400,
		"Banana": 100,
	}
	fruits["Orange"] = 150

	fmt.Println(fruits)

	// USING our defined structure
	p := Person{Name: "Muteeb Ur Rehman", Age: 20}
	fmt.Printf("Person name is %s and age is %d\n", p.Name, p.Age)

	x := 5
	pointer := &x

	fmt.Println("Value of x is: ", x)
	fmt.Println("Pointer address is: ", pointer)
	fmt.Println("Value through pointer is: ", *pointer)

	c := Circle{Radius: 5}
	var s Shape = c
	fmt.Println("Area of the circle is: ", s.Area())

	go printMessage("Hello from Goroutine")
	printMessage("Hello from Main")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello World, From Muteeb!")
	fmt.Println("File contents written to test.txt")

	input := "hello world"
	fmt.Println(strings.ToUpper(input))
}
