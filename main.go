package main

func SumExercise(nums []int64) int64 {
	// Write code for the exerxise here
	// Tip: return -1 for avoiding compiler errors, you can safely remove it
	sum := int64(0)

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {

}
