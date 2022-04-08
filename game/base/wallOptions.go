package base

type WallOptions struct {
	Fixed string //摆牌
}
type IWallOptions interface {
	ShuffleCard(options ...IWallOptions) *Wall
}
type WOption func(options *WallOptions)



