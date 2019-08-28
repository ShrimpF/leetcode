package main

import "fmt"

func main() {
	ar := []int{1, 2, 2, 1}
	fmt.Println(plusOne(ar))
}

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i > -1; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			} else {
				digits[i-1]++
			}
			continue
		}
		break
	}
	return digits
}
