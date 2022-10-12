package work1010

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	data := []struct {
		nums   []int
		target int
		answer int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, answer: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, answer: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, answer: 4},
	}

	for _, v := range data {
		n, tg := v.nums, v.target
		if searchInsert(n, tg) != v.answer {
			t.Fatalf("expect:[%d] != result[%d]", searchInsert(n, tg), v.answer)
		}
	}
}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		} else if nums[i] > target {
			return i
		}
	}
	return len(nums)
}
