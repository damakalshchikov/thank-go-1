package arrays_and_maps

import (
	"fmt"
	"sort"
)

func NumCounter(number int) {
	counter := make(map[int]int)

	for number > 0 {
		num := number % 10
		counter[num]++
		number /= 10
	}

	printCounter(counter)
}

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
