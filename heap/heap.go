package heap

import (
	"math"
	"errors"
)

// Inplace modification of A

func Parent(i int) *int {
	p := int(math.Floor(float64(i-1)/2.0))
	return &p
}

func Left(i *int) *int {
	l := 2*(*i)
	l = l + 1
	return &l
}

func Right(i *int) *int {
	r := 2*(*i)
	r = r + 2
	return &r
}

func HeapMaximum(A *[]int) *int {
	return &(*A)[0]
}

func MaxHeapify(A *[]int, i *int) {
	l := Left(i)
	r := Right(i)
	largest := 0

	if *l < len(*A) && (*A)[*l] > (*A)[*i] {
		largest = *l
	} else {
		largest = *i
	}

	if *r < len(*A) && (*A)[*r] > (*A)[largest] {
		largest = *r
	}

	if largest != *i {
		// Swap indexes
		// a, b = b, a
		(*A)[largest], (*A)[*i] = (*A)[*i], (*A)[largest]
		MaxHeapify(A, &largest)
	}
}

func BuildMaxHeap(A *[]int) {
	A_floor := int(math.Floor(float64(len(*A)))/2.0)-1
	for i := A_floor; i >= 0; i-- {
		MaxHeapify(A, &i)
	}
}

func HeapExtractMax(A *[]int) (int, error) {
	if len(*A) < 1 {
		return 0, errors.New("heap underflow")
	}
	max := (*A)[0]
	i := 0

	// Set first item to last item
	(*A)[i] = (*A)[len(*A)-1]

	// Remove last index 
	(*A) = (*A)[:len(*A)-1] 
	MaxHeapify(A, &i)
	return max, nil
}

func HeapIncreaseKey(A *[]int, i *int, key *int) error {
	j := *i
	if *key < (*A)[j-1] {
		return errors.New("new key is smaller than current key")
	}

	(*A)[j] = *key
	for j > 0 && (*A)[*Parent(j)] < (*A)[j] {
		// Swap indexes
		// a, b = b, a
		(*A)[*Parent(j)], (*A)[j] = (*A)[j], (*A)[*Parent(j)]
		j = *Parent(j)
	}
	return nil
}

func MaxHeapInsert(A *[]int, key *int) error {
	*A = append((*A), -2147483648)
	n := len(*A)-1 // Index starts at one
	return HeapIncreaseKey(A, &n, key)
}