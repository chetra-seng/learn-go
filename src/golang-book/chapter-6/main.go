package main

import "fmt"

func main() {
	fmt.Println(half(1))
}

func half(num int) (int, bool) {
	return int(num / 2), num%2 == 0
}

func max(nums ...int) int {
	var max int

	for i, num := range nums {
		if i == 0 || num > max {
			max = num
		}
	}

	return max
}

func makeOddGenerator() func() uint {
	i := uint(1)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
