package arraysslices

func Sum(numbers []int) int {
	sum := 0
	/* 
	Sans range : on parcourt chaque élément en accédant directement à l'index
		for i := 0; i < 5; i++ {
			sum += numbers[i]
	 	}
	*/

	// Avec range : on ignore l'index (_) et on récupère une copie de chaque élément dans "number"
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// func SumAll(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		sums = append(sums, Sum(numbers))
// 	}

// 	return sums
// }

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {

		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}