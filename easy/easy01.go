package easy

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for index, v := range nums {
		if _, ok := record[v]; ok {
			return []int{record[v], index}
		}
		record[target-v] = index
	}
	return nil
}
