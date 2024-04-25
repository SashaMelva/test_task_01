package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(ComputerStr(0))
	fmt.Println(ComputerStr(1))
	fmt.Println(ComputerStr(3))
	fmt.Println(ComputerStr(5))
	fmt.Println(ComputerStr(10))

	fmt.Println(ComputerStr(-1))
	fmt.Println(ComputerStr(-3))
	fmt.Println(ComputerStr(-5))
	fmt.Println(ComputerStr(-10))
}

func ComputerStr(num int) string {
	res := "компьютер"
	if num == 0 || math.Abs(float64(num)) >= 5 {
		res += "ов"
	} else if math.Abs(float64(num)) >= 2 && math.Abs(float64(num)) < 5 {
		res += "a"
	}
	return strconv.Itoa(num) + " " + res
}
