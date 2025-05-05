package main

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		secondIndex, ok := indexMap[complement]
		if ok {
			return []int{i, secondIndex}
		}
		indexMap[num] = i
	}

	return nil
}
