package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Printf("%#v", numbers)
	return sum
}

func SumAll(numsToSum ...[]int) []int {
	res := make([]int, len(numsToSum))
	for i, numbers := range numsToSum {
		res[i] = Sum(numbers)
	}

	return res
}

func SumAllTails(tailsToSum ...[]int) []int {
	var res []int
	for _, numbers := range tailsToSum {
		if len(numbers) == 0 {
			res = append(res, 0)
			continue
		}

		res = append(res, Sum(numbers[1:]))
	}

	return res
}
