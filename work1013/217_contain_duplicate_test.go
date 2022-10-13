package work1013

import (
	"testing"
)

func TestContainDuplicate(t *testing.T) {
	data := []struct {
		val    []int
		answer bool
	}{
		{val: []int{1, 2, 3, 1}, answer: true},
		{val: []int{1, 2, 3, 4}, answer: false},
		{val: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, answer: true},
	}

	for _, v := range data {
		if containsDuplicate(v.val) != v.answer {
			t.Fatalf("expect:[%t] != result[%t]", containsDuplicate(v.val), v.answer)
		}
	}
}

// containDuplicate判断数组是否存在重复元素
func containsDuplicate(nums []int) bool {
	//双重循环遍历
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
