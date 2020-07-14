package main

import (
	"fmt"
	"github.com/hirokikojima/study-algorithms/search"
	"github.com/hirokikojima/study-algorithms/sort"
)

func main() {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, i)
	}

	search := search.NewSearch(list)

	linear_result, err := search.LinearSearch(6543); 
	if err == nil {
		fmt.Println(linear_result)	
	}

	binary_result, err := search.BinarySearch(6543);
	if err == nil {
		fmt.Println(binary_result)
	}

	hash_result, err := search.HashSearch(6543);
	if err == nil {
		fmt.Println(hash_result)
	}

	

	sort := sort.NewSort(list)

	sort.SelectionSort()

	sort.BubbleSort()

	sort.QuickSort()
}