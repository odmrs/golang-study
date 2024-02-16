// My weight loss program
package main

import "fmt"

// All things start here!
func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" libs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Println(" years old!")

	// Using the Prinf -> output formated

	fmt.Printf("My weight on the surface of Mars is %v lbs,", 75.0 * 0.3783)
	fmt.Printf(" and I would be %v years old!\n", 19 * 365 / 687)

	// More about prinft

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galatic", 100)
}