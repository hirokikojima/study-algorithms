package search

import (
	"errors"
)

type Search interface {
	LinearSearch(target int) (int, error)
	BinarySearch(target int) (int, error)
	HashSearch(target int) (int, error)
}

type search struct {
	List []int
	Sorted []int
	Hashed [10][]int 
}

func NewSearch(list []int) Search {
	s := &search{
		List: list,
	}

	// TODO sort list
	s.Sorted = list

	// TODO hash list
	for _, v := range s.List {
		s.Hashed[v % 10] = append(s.Hashed[v % 10], v)
	}
	return s
}

func (s *search) LinearSearch(target int) (int, error) {
	for _, v := range s.List {
		if v == target {
			return v, nil
		}
	}
	return -1, errors.New("Target is not found.")
}

func (s *search) BinarySearch(target int) (int, error) {
	head := 0
	tail := len(s.Sorted)
	center := head + tail / 2

	for {
		if (head > tail) {
			break
		}

		if (s.Sorted[center] == target) {
			return s.Sorted[center], nil
		}

		if (s.Sorted[center] < target) {
			head = center + 1
		} else {
			tail = center - 1
		}

		center = (head + tail) / 2
	}
	return -1, errors.New("Target is not found.")
}

func (s *search) HashSearch(target int) (int, error) {
	for _, v := range s.Hashed[target % 10] {
		if (target == v) {
			return v, nil
		}
	}
	return -1, errors.New("Target is not found.")
}