package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := make([]int, len(numbersToSum))
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
