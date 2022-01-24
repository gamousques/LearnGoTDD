package main

func Sum(numbers []int) int {
	sum :=0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int ) []int {
	numberOfSlices := len(slices);
	sums := make([]int, numberOfSlices)

	for i, numbers := range slices {
		sums[i] = Sum(numbers)
	}

	return sums;
}