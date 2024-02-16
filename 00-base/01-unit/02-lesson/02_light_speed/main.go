// how long take to travel from Earth to Mars
package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km's
	var distance = 56000000 // km

	fmt.Println(distance / lightSpeed, " distance")

	distance = 401000000 // km

	fmt.Println(distance / lightSpeed, " distance")

	// Quick answer

	const speedOfSpaceX = 100800 // km's
	const oneDayInHours = 24
	var distanceOfMars = 96000000 // km
	var daysOfTravel = distanceOfMars / speedOfSpaceX / oneDayInHours 
	fmt.Printf("-------%v-------\n", "SpaceX")
	fmt.Printf("The SpaceX will travel to Mars in %v days.\n", daysOfTravel)

	// Answer -> hours, minutes = 24, 60
	// Answer -> weight -= 2
}