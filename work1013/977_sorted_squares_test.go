package work1013

import (
	"testing"
)

func TestSortedSquares(t *testing.T) {
	data := []struct {
		val    []int
		answer []int
	}{
		{val: []int{-4, -1, 0, 3, 10}, answer: []int{0, 1, 9, 16, 100}},
		{val: []int{-7, -3, 2, 3, 11}, answer: []int{4, 9, 9, 49, 121}},
	}

	for _, v := range data {
		flag := true
		//比较排序后是否和答案一致
		for i := 0; i < len(v.answer); i++ {
			//发现不一致转换flag标识为false
			if sortedSquares(v.val)[i] != v.answer[i] {
				flag = false
			}
		}
		if !flag {
			t.Fatalf("expect:[%v] != result[%v]", sortedSquares(v.val), v.answer)
		}
	}
}

// sortdeSquares返回非递减数组每个数字平方后的非递减数组
// 数组本来就非递减，故两边向中的平方逐渐减小，所以原数组两边向中可从后往前构建一个正序非递减数组
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i, j, index := 0, n-1, n-1; index >= 0; index-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			ans[index] = nums[i] * nums[i]
			i++
		} else {
			ans[index] = nums[j] * nums[j]
			j--
		}
	}
	return ans
}
