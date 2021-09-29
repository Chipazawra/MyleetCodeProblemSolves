package containerWithMostWater

func MaxArea(height []int) int {
	return maxArea(height)
}

func maxArea(height []int) int {

	l, r, maxarea := 0, len(height)-1, 0

	for l != r {

		if (r-l)*min(height[r], height[l]) > maxarea {
			maxarea = (r - l) * min(height[l], height[r])
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}

	}

	return maxarea
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
