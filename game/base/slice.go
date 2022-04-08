package base

import "math/rand"

func ShuffleSliceCard(seed int64,src []Card) []Card {
	dest := make([]Card, len(src))

	rand.Seed(seed)
	perm := rand.Perm(len(src))

	for i, v := range perm {
		dest[v] = src[i]
	}

	return dest
}

// IntInSlice 判断某个int值是否在切片中
func CardInSlice(finder Card, slice []Card) bool {
	exists := false
	for _, v := range slice {
		if v.Color == finder.Color && v.Number == finder.Number {
			exists = true
			break
		}
	}
	return exists
}

// SliceDel 删除slice中的某些元素
func SliceDel(slice []Card, values ...Card) []Card {
	if slice == nil || len(values) == 0 {
		return slice
	}
	for _, value := range values {
		slice = sliceDel(slice, value)
	}
	return slice
}
func sliceDel(slice []Card, value Card) []Card {
	if slice == nil {
		return slice
	}
	//temp := make(CardVec, 0, len(slice))
	for i, j := range slice {
		if j == value {
			return append(append([]Card{}, slice[:i]...), slice[i+1:]...)
		}
	}
	return slice
}