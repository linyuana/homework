package work1010

import (
	"testing"
)

func TestSearch(a *testing.T) {
	data := []struct {
		nums   []int
		target int
		answer int
	}{
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, answer: 4},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, answer: -1},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: -1, answer: 0},
		{nums: []int{-1, 0, 3, 5, 9, 12}, target: 12, answer: 5},
	}

	for _, v := range data {
		n, tg := v.nums, v.target
		if search(n, tg) != v.answer {
			a.Fatalf("expect:[%d] != result[%d]", search(n, tg), v.answer)
		}
	}
}

func Failed() {
	panic("unimplemented")
}

func search(nums []int, target int) int {
	i, j := 0, len(nums)-1
	mid := 0
	if nums[i] == target {
		return i
	}
	if nums[j] == target {
		return j
	}
	for i < j-1 {
		mid = (i + j) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			j = mid
		} else {
			i = mid
		}
	}
	return -1
}
