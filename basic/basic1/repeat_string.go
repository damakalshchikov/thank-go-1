package basic1

import "fmt"

func RepeatString(sours string, times int) {
	var result string

	for range times {
		result += sours
	}

	fmt.Print(result)
}
