package main

import "fmt"

func main()  {
	nums := []int{1,3, 5, 6}
	target := 7
	var res int
	res = searchInsert(nums, target)
	fmt.Print(res)
}

func searchInsert(nums []int, target int) int {

	high := len(nums)-1
	low := 0
	mid := 0
	for low <= high{
		mid = low + (high-low)/2
		if nums[mid] == target{
			return mid
		}
		if nums[mid] < target{
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	if nums[mid] < target{
		return mid+1
	}
	return mid
}
