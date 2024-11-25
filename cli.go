package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "Muteeb", "My name ")
	age := flag.Int("age", 0, "My age ")

	//	Parse the the command line arguments
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: Hi <name> <age>")
		return
	}

	if *age > 0 {
		fmt.Printf("Hello %s! Your age is %d\n", *name, *age)
	} else {
		fmt.Printf("Hello %s\n", *name)
	}
}
