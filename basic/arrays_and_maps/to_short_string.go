package arrays_and_maps

import "fmt"

func ToShortString(text string, width int) {
	// if len(text) <= width {
	//     fmt.Println(text)
	//     return
	// }

	// text = text[:width] + "..."
	// fmt.Println(text)

	switch len(text) <= width {
	case true:
		fmt.Println(text)
	default:
		res := text[:width] + "..."
		fmt.Println(res)
	}
}
