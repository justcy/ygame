package base

func GetMapKeys(m interface{}) []int {
	temp := m.(map[int]interface{})
	keys := make([]int, len(temp))
	for k := range temp {
		keys = append(keys, k)
	}
	return keys
}
