package main

import (
	"math"
)

func RadixSort(a []int) []int {

	// Define variable
	// Count number of rounds
	var round float64 = 0

	for {

		// Create the filterArray,
		// check_elements for check if still sorting
		// tmp_list to contain lastest value : 101 -> 10
		filterArray := createFilterArray()
		check_element := false
		tmp_list := make([]int, 0, TABLE_SIZE)

		// Delete the lastest elements in number
		for i := 0; i < TABLE_SIZE; i++ {

			tmp_value := a[i] / int(math.Pow(10, round))
			tmp_list = append(tmp_list, tmp_value)

			if tmp_value == 0 {
				check_element = true
			}
		}

		// If no need to sort, all the elements in tmp_list = 0
		if check_element {
			break
		}

		// Get the lastest value by % 10 and insert it into filterArray
		for i := 0; i < TABLE_SIZE; i++ {
			last_value := tmp_list[i] % 10
			filterArray[last_value].Insert(a[i])
		}

		// Call mergeLinkedList and increase round
		a = mergeLinkedList(filterArray)
		round++
	}

	return a
}

// Merge from linked list array to array
func mergeLinkedList(l [10]List) []int {

	// Create result array
	result := make([]int, 0, TABLE_SIZE)

	//Create 10 loop in filterArray list
	for i := 0; i < 10; i++ {

		node := l[i].Head

		// If list[i] != nil (cotainned node, append to result list )
		for node != nil {
			tmp_value := node.Value
			result = append(result, tmp_value)
			node = node.Next
		}

		// l[i].EmptyLinkedList()
	}
	return result
}

// Create filterArray with contains 10 elements from 0 -> 10
// Each elements in filterArray work as a linkedlist, contain the head node
// Init the filter Array to avoid error
func createFilterArray() [10]List {
	var filterArray [10]List

	list := List{
		Head: nil,
		Tail: nil,
	}

	// Add list to filterArray
	for i := range filterArray {
		filterArray[i] = list
	}
	return filterArray
}
