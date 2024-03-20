package main

func TwoSum(nums []int, target int) []int {
	var arr_res []int
	m := make(map[int]int)
	for i, num := range nums {
		val := target - num
		if kes, ok := m[num]; ok {
			if num+val == target {
				arr_res = append(arr_res, kes)
				arr_res = append(arr_res, i)
			}
		} else {
			m[val] = i
		}

	}
	return arr_res
}
