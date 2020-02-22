package main

import "fmt"

func main(){
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	var res int
	res = removeElement(nums, val)
	fmt.Print(res)
}

func removeElement(nums []int, val int) int {
	i := len(nums)
	for j:=0; j<i; j++ {
		for nums[j] == val && j != i{
			i--
			nums[j] = nums[i]
		}
	}
	return i
}
