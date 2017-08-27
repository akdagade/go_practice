package main

import "fmt"

func main() {
	s := greet("Akshay", "Dagade")
	fmt.Println(s)
}

func greet(fname, lname string) (string) {

	return "Hello, " + fname + lname + "!!"
}