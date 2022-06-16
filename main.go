package main

import (
	"flag"
	"fmt"
)

var (
	TABLE_SIZE = 0
	sort       = flag.String("sort", "MergeSort", "Choose sort options")
)

func main() {
	flag.Parse()
	fmt.Println("-----Sorting-------")

	unsortedArray := getArrayInput()
	sortedArray := make([]int, 0, TABLE_SIZE)

	if *sort == "MergeSort" {

		fmt.Println("-----Using Merge Sort----------")
		sortedArray = MergeSort(unsortedArray)
	} else if *sort == "RadixSort" {

		fmt.Println("------Using Radix Sort----------")
		sortedArray = RadixSort(unsortedArray)
	} else {
		fmt.Println("Not such a option")
	}

	// Print sorted array
	fmt.Println("Sorted array: ")
	printArray(sortedArray)
}

// Get numberOfElements and array from screen
func getArrayInput() []int {
	fmt.Println("Number of elements: ")
	fmt.Scanf("%d", &TABLE_SIZE)

	// Create table
	a := make([]int, 0, TABLE_SIZE)
	fmt.Println(a)

	// Get each elements value
	for i := 0; i < TABLE_SIZE; i++ {
		var value int

		fmt.Println("Element ", i, ": ")
		fmt.Scanf("%d", &value)

		a = append(a, value)
	}

	return a
}

// Print Array
func printArray(array []int) {
	fmt.Println(array)
}
