package excersice_1

import "fmt"

func main() {

	var fname string
	var lname string

	fmt.Printf("Enter your fname: ")
	fmt.Scan(&fname)
	fmt.Printf("Enter your lname: ")
	fmt.Scan(&lname)
	fmt.Printf("Hello, " + fname + " " + lname + "!!")

}
