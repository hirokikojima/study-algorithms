package sort

type Sort interface {
	SelectionSort() ([]int, error)
	BubbleSort() ([]int, error)
	InsertionSort() ([]int, error)
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
	for i := 0; i < len(s.List); i++ {
		m := i
		for k := i + 1; k < len(s.List); k++ {
			if s.List[i] > s.List[k] {
				m = k
			}
		}

		if i != m {
			w := s.List[i]
			s.List[i] = s.List[m]
			s.List[m] = w
		}
	}
	return s.List, nil
}

func (s *sort) BubbleSort() ([]int, error) {
	for i := 0; i < len(s.List); i++ {
		for k := len(s.List) - 1; k > i; k-- {
			if s.List[k] < s.List[k - 1] {
				w := s.List[k]
				s.List[k] = s.List[k - 1]
				s.List[k-1] = w
			} 
		}
	}
	return s.List, nil
}

func (s *sort) InsertionSort() ([]int, error) {
	for i := 1; i < len(s.List); i++ {
		w := s.List[i]
		for k := i; k >= 0; k-- {
			if k == 0 || s.List[k - 1] < w {
				s.List[k] = w
				break
			} else {
				s.List[k] = s.List[k - 1]
			}
		}
	}
	return s.List, nil
}

func (s *sort) QuickSort() ([]int, error) {
	left := 0
	right := len(s.List) - 1

	s.subQuickSort(left, right)

	return s.List, nil
}

func (s *sort) subQuickSort(left int, right int) error {	
	i := left + 1
	k := right

	for i < k {
		for i < right {
			if s.List[i] > s.List[left] {
				break
			}
			i++
		}
		for k > left {
			if s.List[k] <= s.List[left] {
				break
			}
			k--
		}

		if i < k {
			w := s.List[i]
			s.List[i] = s.List[k]
			s.List[k] = w
		}
	}

	if s.List[left] > s.List[k] {
		w := s.List[left]
		s.List[left] = s.List[k]
		s.List[k] = w
	}

	if left + 1 < k {
		s.subQuickSort(left, k - 1)
	}
	if right -1 > k {
		s.subQuickSort(k + 1, right)
	}

	return nil
}