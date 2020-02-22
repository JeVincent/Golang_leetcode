package main

import "fmt"

func main(){
	nums := []int{2,7,11,15}
	target := 9
	var res []int

	res = twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	var res []int
	var m map[int]int
	m = make(map[int]int)
	for i,v := range(nums){
		_, ok := m[target-v]
		if ok {
			res = append(res, m[target-v],i)
			break
		}
		m[v] = i
	}
	return res
}

