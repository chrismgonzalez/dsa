package arraysandslices

// SumSlice adds the numbers in the input slice and returns the sum
func SumSlice(numbers []int) (int, error){
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum, nil
}