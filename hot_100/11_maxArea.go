package main

func maxArea(height []int) int {
	if len(height) < 1 {
		return 0
	}
	if len(height) == 2 {
		return calArea(0, height[0], 1, height[1])
	}
	i := 0
	j := len(height) - 1
	area := calArea(i, height[i], j, height[j])
	for i < j {
		ih := height[i]
		jh := height[j]
		curArea := calArea(i, ih, j, jh)
		if area < curArea {
			area = curArea
		}
		if ih < jh {
			i++
		} else {
			j--
		}
	}
	return area
}
func calArea(fx, fy, sx, sy int) int {
	h := 0
	if fy > sy {
		h = sy
	} else {
		h = fy
	}
	return (sx - fx) * h
}
