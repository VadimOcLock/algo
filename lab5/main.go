package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	const n = 100000
	bubbleArray := make([]int, n)
	conditionFunc := func(a int, b int) bool {
		return a <= b
	}

	for i := 0; i < n; i++ {
		bubbleArray[i] = rand.Intn(1000)
	}

	shakerArray := make([]int, n)
	copy(shakerArray, bubbleArray)
	oddEvenArray := make([]int, n)
	copy(oddEvenArray, bubbleArray)
	combArray := make([]int, n)
	copy(combArray, bubbleArray)
	insertArray := make([]int, n)
	copy(insertArray, bubbleArray)
	insertBinaryArray := make([]int, n)
	copy(insertBinaryArray, bubbleArray)
	insertPairArray := make([]int, n)
	copy(insertPairArray, bubbleArray)

	bubbleSortStart := time.Now()
	BubbleSort(bubbleArray, conditionFunc)
	bubbleSortTime := time.Since(bubbleSortStart)
	fmt.Printf("bubbleSortTime: %s\n", bubbleSortTime)

	shakerSortStart := time.Now()
	ShakerSort(shakerArray, conditionFunc)
	shakerSortTime := time.Since(shakerSortStart)
	fmt.Printf("shakerSortTime: %s\n", shakerSortTime)

	oddEvenSortStart := time.Now()
	OddEvenSort(oddEvenArray, conditionFunc)
	oddEvenSortTime := time.Since(oddEvenSortStart)
	fmt.Printf("oddEvenSortTime: %s\n", oddEvenSortTime)

	combSortStart := time.Now()
	CombSort(combArray, conditionFunc)
	combSortTime := time.Since(combSortStart)
	fmt.Printf("combSortTime: %s\n", combSortTime)

	insertSortStart := time.Now()
	InsertSort(insertArray, conditionFunc)
	insertSortTime := time.Since(insertSortStart)
	fmt.Printf("insertSortTime: %s\n", insertSortTime)

	insertBinarySortStart := time.Now()
	InsertBinarySort(insertBinaryArray, conditionFunc)
	insertBinarySortTime := time.Since(insertBinarySortStart)
	fmt.Printf("insertBinarySortTime: %s\n", insertBinarySortTime)

	insertPairSortStart := time.Now()
	InsertBinarySort(insertPairArray, conditionFunc)
	insertPairSortTime := time.Since(insertPairSortStart)
	fmt.Printf("insertPairSortTime: %s\n", insertPairSortTime)

}

func BubbleSort(arr []int, conditionFunc func(a int, b int) bool) []int {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if conditionFunc(arr[j+1], arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func ShakerSort(arr []int, conditionFunc func(a int, b int) bool) []int {

	left := 0
	right := len(arr) - 1
	swapped := true

	for swapped {
		swapped = false
		for i := left; i < right; i++ {
			if !conditionFunc(arr[i], arr[i+1]) {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		right--

		if !swapped {
			break
		}

		swapped = false
		for i := right; i > left; i-- {
			if conditionFunc(arr[i], arr[i-1]) {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
		left++
	}

	return arr
}

func OddEvenSort(arr []int, conditionFunc func(a int, b int) bool) []int {

	sorted := false

	for !sorted {

		sorted = true

		var wg sync.WaitGroup
		swappedChanel := make(chan bool, 2)

		wg.Add(2)

		go func() {
			defer wg.Done()
			localSwapped := false
			for i := 0; i < len(arr)-1; i += 2 {
				if !conditionFunc(arr[i], arr[i+1]) {
					arr[i], arr[i+1] = arr[i+1], arr[i]
					localSwapped = true
				}
			}
			swappedChanel <- localSwapped
		}()

		go func() {
			defer wg.Done()
			localSwapped := false
			for i := 1; i < len(arr)-1; i += 2 {
				if !conditionFunc(arr[i], arr[i+1]) {
					arr[i], arr[i+1] = arr[i+1], arr[i]
					localSwapped = true
				}
			}
			swappedChanel <- localSwapped
		}()

		wg.Wait()
		close(swappedChanel)

		for s := range swappedChanel {
			if s {
				sorted = false
			}
		}
	}

	return arr
}

func CombSort(arr []int, conditionFunc func(a int, b int) bool) []int {
	gap := len(arr)
	factor := 1.247
	swapped := true

	for gap > 1 || swapped {
		gap = int(float64(gap) / factor)
		if gap < 1 {
			gap = 1
		}

		swapped = false
		for i := 0; i+gap < len(arr); i++ {
			if !conditionFunc(arr[i], arr[i+gap]) {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
		}
	}

	return arr
}

func InsertSort(arr []int, conditionFunc func(a int, b int) bool) []int {

	for i := 1; i < len(arr); i++ {

		k := i
		for k > 0 {
			if conditionFunc(arr[k], arr[k-1]) {
				arr[k-1], arr[k] = arr[k], arr[k-1]
			}
			k--
		}
	}

	return arr
}

func InsertBinarySort(arr []int, conditionFunc func(a int, b int) bool) []int {

	for i := 1; i < len(arr); i++ {
		item := arr[i]
		j := i - 1

		pos, err := BinarySearch(arr, item, 0, j, conditionFunc)
		if err != nil {
			return nil
		}

		for pos <= j {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = item
	}
	return arr
}

var ErrIncorrectBinarySearchParams = fmt.Errorf("incorrect input param")

func BinarySearch(arr []int, item int, low int, high int, conditionFunc func(a int, b int) bool) (int, error) {

	if len(arr) == 0 || low > high {
		return -1, ErrIncorrectBinarySearchParams
	}

	for low <= high {
		mid := (low + high) / 2
		if conditionFunc(arr[mid], item) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low, nil
}

func PairInsertSort(arr []int, conditionFunc func(a int, b int) bool) []int {

	i := 0
	for i < len(arr)-1 {

		maxItem := arr[i]
		minItem := arr[i+1]

		if minItem > maxItem {
			minItem, maxItem = maxItem, minItem
		}

		j := i - 1
		for j >= 0 && arr[j] > maxItem {
			arr[j+2] = arr[j]
			j--
		}
		arr[j+2] = maxItem

		for j >= 0 && arr[j] > minItem {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = minItem

		i += 2
	}

	if i == len(arr)-1 {

		y := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > y {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = y
	}

	return arr
}
