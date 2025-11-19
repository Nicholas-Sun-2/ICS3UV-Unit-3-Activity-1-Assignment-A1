/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-18
 * @fileoverview This program calculates the area of a circle.
 */

package main

import "fmt"

func main() {
	// Declaring variables
	var radius float32 = 5.6
	const PI float32 = 3.14

	var area float32 = PI * radius * radius

	// Printing
	fmt.Println("A circle with radius of " + fmt.Sprint(radius) + " cm has an area of " + fmt.Sprint(area) + " cm^2.")
}