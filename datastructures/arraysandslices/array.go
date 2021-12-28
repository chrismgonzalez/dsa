package arraysandslices


// Sum adds the numbers in the input array and returns the sum
func Sum(numbers [4]int) (int, error){
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum, nil
}