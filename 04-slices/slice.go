package slice

func SumArr(numbers [5]int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numToSum {
		sum = append(sum, SumSlice(numbers))
	}
	return sum
}

func SumAllTails(numToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, SumSlice(tail))
		}
	}
	return sum
}
