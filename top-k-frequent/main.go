package main

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	Frequency int
	Value     int
}

type MinHeap []Pair

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].Frequency < h[j].Frequency
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	h := &MinHeap{}
	heap.Init(h)
	for num, freq := range frequency {
		heap.Push(h, Pair{Frequency: freq, Value: num})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, h.Len())
	for i := range k {
		result[i] = heap.Pop(h).(Pair).Value
	}
	return result
}

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	nums2 := []int{1}
	nums3 := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}

	fmt.Println(topKFrequent(nums1, 2)) // Output: [1, 2]
	fmt.Println(topKFrequent(nums2, 1)) // Output: [1]
	fmt.Println(topKFrequent(nums3, 2)) // Output: [1, 2]
}
