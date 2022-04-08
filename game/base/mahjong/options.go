package mahjong

import "github.com/justcy/ygame/game/base"

var defaultPokerOptions = MahjongOptions{OnlyCharacters: false,HasCharacters: true, HasWind: false, HasDragon: false, HasFlower:false, GhostDragon: false, GhostWhite: false, EightDragon: false}
type MahjongOptions struct {
	base.WallOptions
	OnlyCharacters bool //是否只有万
	HasCharacters bool //是否有万
	HasWind       bool //有风
	HasDragon     bool //有字
	HasFlower     bool //有花
	GhostDragon   bool //红中为鬼牌
	GhostWhite    bool //白板为鬼牌
	EightDragon   bool //八个红中
}
type MahjongOption func(*MahjongOptions)

func WithFixed(s string) MahjongOption {
	return func(args *MahjongOptions) {
		args.Fixed = s
	}
}
func WithOnlyCharacters(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.OnlyCharacters = b
	}
}
func WithCharacters(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.HasCharacters = b
	}
}
func WithWind(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.HasWind = b
	}
}
func WithDragon(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.HasDragon = b
	}
}
func WithFlower(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.HasFlower = b
	}
}
func WithGhostDragon(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.GhostDragon = b
	}
}
func WithGhostWhite(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.GhostWhite = b
	}
}
func WithEightDragon(b bool) MahjongOption {
	return func(args *MahjongOptions) {
		args.EightDragon = b
	}
}

func ShuffleMahjong(seed int64,options ...MahjongOption) *base.Wall  {
	wall := &base.Wall{}
	config := defaultPokerOptions
	for _, o := range options {
		o(&config)
	}

	colorStart := int8(1)
	if config.HasCharacters == false { //没有万
		colorStart = 2
	}
	colorEnd := int8(3)
	if config.OnlyCharacters {
		colorStart = 1
		colorEnd = 1
	}
	for ; colorStart <= colorEnd; colorStart++ { //万 筒 条
		for y := int8(1); y <= 9; y++ {
			wall.FullCard = append(wall.FullCard, base.Card{Color: colorStart, Number: y})
		}
	}
	if config.HasWind { //有风 东 南 西 北
		for y := int8(1); y <= 4; y++ {
			wall.FullCard = append(wall.FullCard, base.Card{Color: MjColorWind, Number: y})
		}
	}
	if config.HasDragon { //中 发 白
		for y := int8(1); y <= 3; y++ {
			wall.FullCard = append(wall.FullCard, base.Card{Color: MjColorDragon, Number: y})
		}
	}else{
		if config.GhostDragon { //红中为鬼牌
			wall.FullCard = append(wall.FullCard, base.Card{Color: MjColorDragon, Number: 1})

			if config.EightDragon { //八个红中
				wall.FullCard = append(wall.FullCard, base.Card{Color: MjColorDragon, Number: 1})
			}
		} else if config.GhostWhite { //白板为鬼牌
			wall.FullCard = append(wall.FullCard, base.Card{Color: MjColorDragon, Number: 3})
		}
	}
	for z := 0; z < 4; z++ { //每样4张
		for i := range wall.FullCard {
			wall.Tiles = append(wall.Tiles, wall.FullCard[i])
		}
	}
	if config.HasFlower { //花牌
		for z := int8(1); z <= 8; z++ {
			wall.FullCard = append(wall.FullCard, base.Card{Color: MjColorFlower, Number: z})
			wall.Tiles = append(wall.Tiles, base.Card{Color: MjColorFlower, Number: z})
		}
	}

	wall.Shuffle(seed)
	wall.Shuffle(seed)

	if config.Fixed != ""{
		wall.FixedCard(config.Fixed)
	}
	return wall
}
