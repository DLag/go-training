package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int, c chan struct{}) {
	if c == nil {
		return
	}
	defer func() {
		c <- struct{}{}
	}()
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(arr)

	// Move the pivot to the right
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	arr[left], arr[right] = arr[right], arr[left]

	// Go down the rabbit hole
	end := make(chan struct{})
	go QuickSort(arr[:left], end)   // Left side
	go QuickSort(arr[left+1:], end) // Right side
	<-end
	<-end
}

func main() {
	arr := []int{9, 2, 10, 12, 3, 7, 8, 4, 6, 5, 1}
	fmt.Println("Initial array is:", arr)
	end := make(chan struct{})
	go QuickSort(arr, end)
	<-end
	fmt.Println("Sorted array is:", arr)
}
