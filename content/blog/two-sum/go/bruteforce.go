package main

func twoSumBruteforce(nums []int, target int) []int {
	for firstIdx, main := range nums {
		for secondIdx, complement := range nums {
			if firstIdx == secondIdx {
				continue
			}
			if main+complement == target {
				return []int{firstIdx, secondIdx}
			}
		}
	}

	return nil
}
