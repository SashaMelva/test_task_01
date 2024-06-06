package main

import "fmt"

func main() {
	MultiplicationTable(5)
	MultiplicationTable(0)
	MultiplicationTable(1)
}

func MultiplicationTable(num int) {
	hash := make([][]int, num+1)
	ferst := make([]int, num+1)

	for i := 0; i <= num; i++ {
		ferst[i] = i
	}
	hash[0] = ferst

	for i := 1; i < len(hash); i++ {
		second := make([]int, num+1)
		second[0] = ferst[i]
		for j := 1; j < len(second); j++ {
			second[j] = second[0] * ferst[j]
		}

		hash[i] = second
	}

	for i := range hash {
		fmt.Println(hash[i])
	}
}
