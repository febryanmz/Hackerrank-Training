package main

func BirthdayCakeCandles(candles []int32) int32 {
	count, max := int32(0), int32(0)

	for _, val := range candles {
		if val > max {
			max = val
			count = 0
		}
		if val == max {
			count++
		}
	}
	return count
}
