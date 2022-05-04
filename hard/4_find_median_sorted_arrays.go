/**
 * @Author raven
 * @Description
 * @Date 2022/4/5
 **/
package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	i := 0
	j := 0
	mergeArr := make([]int, 0, n1+n2)
	for {
		if i >= n1 && j >= n2 {
			break
		}
		if i < n1 && j < n2 {
			if nums1[i] <= nums2[j] {
				mergeArr = append(mergeArr, nums1[i])
				i++
			} else {
				mergeArr = append(mergeArr, nums2[j])
				j++
			}
		} else if i < n1 && j >= n2 {
			mergeArr = append(mergeArr, nums1[i])
			i++
		} else if j < n2 && i >= n1 {
			mergeArr = append(mergeArr, nums2[j])
			j++
		}
	}
	if (n1+n2)%2 == 0 {
		mid := (n1 + n2) / 2
		return float64(mergeArr[mid-1]+mergeArr[mid]) / 2
	} else {
		return float64(mergeArr[(n1+n2)/2])
	}
}
