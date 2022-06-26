/**
 * @Author raven
 * @Description
 * @Date 2022/6/13
 **/
package main

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxArea := 0
	for i < j {
		weight := j - i
		h := min(height[i], height[j])
		area := weight * h
		if maxArea < area {
			maxArea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
