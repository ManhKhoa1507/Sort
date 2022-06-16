package main

// Sort array using merge sort
func MergeSort(array []int) []int {

	// If lenArray < 1 don't split array
	lenArray := len(array)
	if lenArray <= 1 {
		return array
	}

	// Create right and left array with len is mid = lenArray / 2
	mid := lenArray / 2
	leftArray := array[:mid]
	rightArray := array[mid:]

	// Call MergeSort for the left and right array
	leftArray = MergeSort(leftArray)
	rightArray = MergeSort(rightArray)

	// Merge left and right array to main array
	return Merge(leftArray, rightArray)
}

// Merge sort
func Merge(leftArray, rightArray []int) []int {

	// Create result array
	result := make([]int, 0, len(leftArray)+len(rightArray))

	// i : index of leftArray
	// j : index of rightArray
	i, j := 0, 0
	for i < len(leftArray) && j < len(rightArray) {

		// If get the minimum between leftArray[i] , rightArray[j], push the min value to result
		// Increase index of left or right
		if leftArray[i] < rightArray[j] {
			result = append(result, leftArray[i])
			i++
		} else {
			result = append(result, rightArray[j])
			j++
		}
	}

	// Push remaining leftArray to result array
	for i < len(leftArray) {
		result = append(result, leftArray[i])
		i++
	}

	// Push remaining leftArray to result array
	for j < len(rightArray) {
		result = append(result, rightArray[j])
		j++
	}

	return result
}
