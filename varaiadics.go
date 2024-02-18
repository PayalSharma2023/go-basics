package main

func addnumbers(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
func main() {
	arrOne := []int{1, 2, 3}
	arrTwo := []int{4, 5, 6}

	all := append(arrOne, arrTwo...)
	println(addnumbers(all...))
}
