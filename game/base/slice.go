package base

import "math/rand"

func GetCardsColorsAndNumbers(hand []Card) (color, number []int8) {
	for _, card := range hand {
		color = append(color, card.Color)
		number = append(number, card.Number)
	}
	return color, number
}

//Unique int8 数组去重
func Unique(i []int8) []int8 {
	result := make([]int8, 0, len(i))
	m := make(map[int8]bool)
	for _, v := range i {
		if _, exists := m[v]; !exists {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

//CountCard  统计相同牌有多少张
func CountCard(hand []Card) map[int][]Card {
	r := map[int][]Card{}
	temp := UniqueCards(hand)
	for _, v := range temp {
		x := CountInSlice(v, hand)
		r[int(x)] = append(r[int(x)], v)
	}
	return r
}

func CountCardByNumber(hand []Card) map[int][]Card {
	r := map[int][]Card{}
	temp := UniqueCardsByType(hand)
	for _, v := range temp {
		finder := v.Number
		x := CountInSliceByNumber(finder, hand)
		r[int(x)] = append(r[int(x)], v)
	}
	return r
}

//查找相同的牌
func CountInSlice(finder Card, slice []Card) int32 {
	exists := int32(0)
	for _, v := range slice {
		if v == finder {
			exists++
		}
	}
	return exists
}

//CountInSliceByNumber 根据条件查找相同牌
func CountInSliceByNumber(finder int8, slice CardVec) int32 {
	exists := int32(0)
	for _, v := range slice {
		if v.Number == finder {
			exists++
		}
	}
	return exists
}

//UniqueCards  剔除重复的牌
func UniqueCards(s []Card) []Card {
	result := make([]Card, 0, len(s))
	m := make(map[Card]bool)
	for _, v := range s {
		if _, exists := m[v]; !exists {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

//UniqueCardsByType 根据点数剔除重复的牌
func UniqueCardsByType(s []Card) []Card {
	result := make([]Card, 0, len(s))
	m := make(map[int8]bool)
	for _, v := range s {
		if _, exists := m[v.Number]; !exists {
			m[v.Number] = true
			result = append(result, v)
		}
	}
	return result
}

//ShuffleSliceCard 根据随机数种子打乱牌
func ShuffleSliceCard(seed int64, src []Card) []Card {
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
