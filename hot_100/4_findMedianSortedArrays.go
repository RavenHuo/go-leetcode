package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums2) == 0 && len(nums1) == 0 {
		return 0
	}
	if len(nums1) == 0 {
		return findMedian(nums2)
	}
	if len(nums2) == 0 {
		return findMedian(nums1)
	}
	a1 := nums1[0]
	a2 := nums1[len(nums1)-1]
	b1 := nums2[0]
	b2 := nums2[len(nums2)-1]
	// 没有交集
	if !(a2 > b1 && b2 > a1) {
		if a2 > b1 {
			return float64(b2+a1) / 2
		} else {
			return float64(a2+b1) / 2
		}
	}
	if a1 < b1 {
		return findMid(nums1, nums2, b1)
	}
	return findMid(nums2, nums1, a1)
}

func findMid(nums1 []int, nums2 []int, b1 int) float64 {
	// 便利数组num1
	i := 0
	j := len(nums1) - 1
	for i <= j {
		if nums1[i] < b1 {
			i++
		}
		if nums1[j] > b1 {
			j--
		}
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(nums1[i]+b1) / 2
	} else {
		return float64(b1)
	}
}
func findMedian(nums1 []int) float64 {
	if len(nums1)%2 == 0 {
		return float64(nums1[len(nums1)/2-1]+nums1[len(nums1)/2]) / 2
	}
	return float64(nums1[len(nums1)/2])
}
func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
}
