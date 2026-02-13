package funcs_and_points

import (
	"bufio"
	"os"
	"strconv"
)

func Filter(predicate func(int) bool, iterable []int) []int {
	res := make([]int, 0)

	for _, num := range iterable {
		if predicate(num) {
			res = append(res, num)
		}
	}

	return res
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func ReadInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
