package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 5, 0, 0, 0}
	m := 3
	nums2 := []int{2, 4, 6}
	n := 3
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Start merging from the back
	i := m - 1     // Last index of the first m elements in nums1
	j := n - 1     // Last index of nums2
	k := m + n - 1 // Last index of nums1 (after merging)

	fmt.Println("Initial State:")
	fmt.Printf("nums1: %v\nnums2: %v\n", nums1, nums2)

	// Merge nums1 and nums2, starting from the end
	for i >= 0 && j >= 0 {
		fmt.Printf("\nComparing nums1[%d] = %d and nums2[%d] = %d\n", i, nums1[i], j, nums2[j])

		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			fmt.Printf("Placed nums1[%d] = %d at nums1[%d]\n", i, nums1[i], k)
			i--
		} else {
			nums1[k] = nums2[j]
			fmt.Printf("Placed nums2[%d] = %d at nums1[%d]\n", j, nums2[j], k)
			j--
		}
		k--
		fmt.Printf("nums1 after this step: %v\n", nums1)
	}

	// If there are remaining elements in nums2, copy them
	for j >= 0 {
		nums1[k] = nums2[j]
		fmt.Printf("\nCopying remaining nums2[%d] = %d to nums1[%d]\n", j, nums2[j], k)
		j--
		k--
		fmt.Printf("nums1 after copying: %v\n", nums1)
	}

	fmt.Println("\nFinal Merged nums1:", nums1)
}
