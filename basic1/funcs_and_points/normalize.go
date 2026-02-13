package funcs_and_points

func Normalize(nums ...*float64) {
	divider := func(x []*float64) float64 {
		var res float64
		for _, num := range x {
			res += *num
		}

		return res
	}(nums)

	for _, num := range nums {
		*num = *num / divider
	}
}
  