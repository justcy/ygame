package base

import "sort"

type Card struct {
	Color  int8 //花色
	Number int8 //数字
}
type Icard interface {
	//Shuffle(seed int64,) Wall
	toString(cards Cards) string
}
type Cards []Card
type cardVec []Card

func (ms cardVec) Len() int {
	return len(ms)
}
func (ms cardVec) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
func (ms cardVec) Less(i, j int) bool {
	if ms[i].Color != ms[j].Color {
		return ms[i].Color < ms[j].Color
	} else {
		return ms[i].Number < ms[j].Number
	}
}
func (ms cardVec) Sort() {
	sort.Sort(ms)
}
