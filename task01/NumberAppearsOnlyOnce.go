func singleNumber(nums []int) int {
	temp := make(map[int]int)
	for _, v := range nums {
		_, exists := temp[v]
		if exists {
			temp[v] += 1
		} else {
			temp[v] = 1
		}
	}

	for k, v := range temp {
		if v == 1 {
			return k
		}
	}
	return -1
}