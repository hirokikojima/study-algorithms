package sort

type Sort interface {
	SelectionSort() ([]int, error)
	BubbleSort() ([]int, error)
	QuickSort() ([]int, error)
}

type sort struct {
	List []int
}

func NewSort(list []int) Sort {
	return &sort{
		List: list,
	}
}

func (s *sort) SelectionSort() ([]int, error) {
	return s.List, nil
}

func (s *sort) BubbleSort() ([]int, error) {
	return s.List, nil
}

func (s *sort) QuickSort() ([]int, error) {
	return s.List, nil
}