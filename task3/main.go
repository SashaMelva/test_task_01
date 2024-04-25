package main

import "fmt"

func main() {
	fmt.Println(GetPeriod(0, 20))
	fmt.Println(GetPeriod(0, 1))
	fmt.Println(GetPeriod(11, 20))
	fmt.Println(GetPeriod(10, 20))
	fmt.Println(GetPeriod(10, 21))
}

func GetPeriod(start, end int) []int {
	res := []int{}
	if start%2 == 0 {
		start += 1
	}
	for start <= end {
		res = append(res, start)
		start += 2
	}
	return res
}
