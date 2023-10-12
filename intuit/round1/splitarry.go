package main

import "fmt"
//time o(n) space o(n)
func solve(ar []int) string {
	sum := 0
	for _, num := range ar {
		sum += num
	}

	selectSum := 0
	for _, num := range ar {
		selectSum += num
		sum -= num
		if selectSum == sum {
			return "yes"
		} else if selectSum > sum {
			return "no"
		}
	}
	return "no"
}

func main() {
	ar := []int{1, 2, 3, 4, 2}
	fmt.Println(solve(ar))
	ar2 := []int{1, 1, 1}
	fmt.Println(solve(ar2))
}
