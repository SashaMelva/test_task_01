package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(commonDivisor([]int{9, 27, 18}))
	fmt.Println(commonDivisor([]int{}))
	fmt.Println(commonDivisor([]int{0}))
	fmt.Println(commonDivisor([]int{1, 1, 1}))
	fmt.Println(commonDivisor([]int{27}))
}

func commonDivisor(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	numMin := nums[0]
	res := make(map[int]int)

	divider := numMin
	div := 1
	for divider > 0 {
		if numMin%divider == 0 {
			if res[divider] == 0 {
				res[divider] = 1
			}
			div = numMin / divider
			if res[div] == 0 {
				res[div] = 1
			}
		}
		divider -= 1
	}

	for i := 1; i < len(nums); i++ {
		for key := range res {
			if nums[i]%key > 0 {
				res[key] = 0
			}
		}
	}

	resOst := make([]int, 0)
	for key, val := range res {
		if val > 0 {
			resOst = append(resOst, key)
		}
	}

	return resOst
}
