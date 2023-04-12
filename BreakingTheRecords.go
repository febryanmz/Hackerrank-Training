package main

func Records(Scores []int) []int {
	result := []int32{0, 0}
	max, min := Scores[0], Scores[0]

	for i := 1; i < len(Scores); i++ {
		if Scores[i] > max {
			max = Scores[i]
			result[0] += 1
		}

		if Scores[i] < min {
			min = Scores[i]
			result[1] += 1
		}

	}
	return result
}
