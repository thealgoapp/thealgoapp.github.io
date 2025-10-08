package main

func maxAreaBruteforce(height []int) int {
	maxVolume := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			volume := (j - i) * min(height[i], height[j])
			if volume > maxVolume {
				maxVolume = volume
			}
		}
	}

	return maxVolume
}
