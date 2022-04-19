package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		j, exist := m[target-num]

		if exist {
			return []int{i, j}
		}

		m[num] = i
	}

	return []int{-1, -1}
}

// func twoSum(nums []int, target int) []int {
// 	var result []int

// 	for index1, value1 := range nums {
// 		for index2, value2 := range nums {
// 			if index1 != index2 && value1+value2 == target && len(result) == 0 {

// 				result = append(result, index1)
// 				result = append(result, index2)

// 				break
// 			}
// 		}
// 	}

// 	return result
// }

func main() {
	nums := []int{3, 2, 4}
	target := 6

	// twoSum(nums, target)
	result := twoSum(nums, target)
	fmt.Println(result)
}
