func search(nums []int, target int) int {
    i, j := 0, len(nums) - 1
    mid := 0
    if nums[i] == target {
            return i
    }
    if nums[j] == target {
            return j
    }
    for i < j-1{
        mid = (i + j) / 2
        if nums[mid] == target {
            return mid
        }else if nums[mid] > target{
            j = mid
        }else {
            i = mid
        }
    }
    return -1
}