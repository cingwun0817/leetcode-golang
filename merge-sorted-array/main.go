func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums1); j++ {
			if j > i {
				if nums1[i] > nums1[j] {
					// fmt.Println(nums1[i], nums1[j])

					num := nums1[i]
					nums1[i] = nums1[j]
					nums1[j] = num
				}
			}
		}
	}
}