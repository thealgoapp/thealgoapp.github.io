package main

func maxAreaWindow(height []int) int {
	maxVolume, i, j := 0, 0, len(height)-1
	for i < j {
		width := j - i
		maxVolume = max(maxVolume, width*min(height[i], height[j]))

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxVolume
}
