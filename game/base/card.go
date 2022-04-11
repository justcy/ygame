package base

import "sort"

const (
	CardEumNIL    int8 = iota // 留空
	CardEumNumber             // 点数 1
	CardEumColor              // 颜色 2
)
type Card struct {
	Color  int8 //花色
	Number int8 //数字
}
type Icard interface {
	//Shuffle(seed int64,) Wall
	toString(cards Cards) string
}
type Cards []Card
type CardVec []Card

func (ms CardVec) Len() int {
	return len(ms)
}
func (ms CardVec) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
func (ms CardVec) Less(i, j int) bool {
	if ms[i].Color != ms[j].Color {
		return ms[i].Color < ms[j].Color
	} else {
		return ms[i].Number < ms[j].Number
	}
}
func (ms CardVec) Sort() {
	sort.Sort(ms)
}
