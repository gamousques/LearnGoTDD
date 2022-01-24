package main

func Sum(numbers []int) int {
	sum :=0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int ) []int {

	var sums []int
	for _, numbers := range slices {
		sums = append(sums, Sum(numbers))
	}

	return sums;
}

func SumAllTails(slices...[]int) []int {
	var sums[]int
	for _, slice := range slices {
		if len(slice) != 0 {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}