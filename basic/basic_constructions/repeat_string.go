package basic_constructions

import "fmt"

func RepeatString(sours string, times int) {
	var result string

	for range times {
		result += sours
	}

	fmt.Println(result)
}
