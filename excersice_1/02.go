package excersice_1

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Print("Enter number 1: ")
	fmt.Scan(&n1)
	fmt.Print("Enter number 2: ")
	fmt.Scan(&n2)
	fmt.Println()
	if n1 > n2 {
		fmt.Printf("%d", n1%n2)
	} else {
		fmt.Printf("%d", n2%n1)
	}
}
