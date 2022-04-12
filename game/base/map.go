package base

func GetMapKeys(m map[int]CardVec) []int {
	keys := make([]int, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
