package array

func SumAllRest(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			final := nums[1:]
			sums = append(sums, Sum(final))
		}
	}

	return sums
}