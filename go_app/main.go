// Created by: Charlotte Jhu
// Created on: March 2023
//
// This program shows an adress

package main

import "fmt"

func main() {
	// This function shows an adress
	// variables
	var streetNumber string
	var streetName string

	// input
	fmt.Println("This program shows an adress")
	fmt.Println()
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetNumber)
	fmt.Print("Enter your street name: ")
	fmt.Scanln(&streetName)

	// output
	fmt.Println("Your address is:", streetNumber, streetName)

	fmt.Println("\nDone.")
}
