package poker

import (
	"github.com/justcy/ygame/game/base"
)

//桌子状态
const (
	PokerWallNIL        int = iota
	PokerWallAll            //全部扑克牌
)
var defaultPokerOptions = PokerOptions{PokerWallType: 1,IsJoker: false, IsGhost: false, Repeat: 1}

type PokerOptions struct {
	base.WallOptions
	PokerWallType int //1.全部牌 2.15张跑得快@todo
	IsJoker bool //是否有王牌
	IsGhost bool //是否有鬼牌
	Repeat  int  //几副牌
}
type PokerOption func(*PokerOptions)

func WithFixed(s string) PokerOption {
	return func(args *PokerOptions) {
		args.Fixed = s
	}
}
func WithJoker(j bool) PokerOption {
	return func(args *PokerOptions) {
		args.IsJoker = j
	}
}
func WithGhost(g bool) PokerOption {
	return func(args *PokerOptions) {
		args.IsGhost = g
	}
}
func WithRepeat(r int) PokerOption {
	return func(args *PokerOptions) {
		args.Repeat = r
	}
}
func WithPokerWallType(t int) PokerOption {
	return func(args *PokerOptions) {
		args.PokerWallType = t
	}
}

func ShufflePoker(seed int64,options ...PokerOption) *base.Wall {
	wall := &base.Wall{}
	wall.FullCard = make(base.Cards, 0)
	config := defaultPokerOptions
	for _, o := range options {
		o(&config)
	}

	for i := 0; i < config.Repeat; i++ {
		if config.PokerWallType == PokerWallAll {
			for x := PokerColorDiamond; x <= PokerColorSpade; x++ { //黑桃到方块
				for y := PokerPoint2; y <= PokerPointA; y++ { //2到A
					wall.Tiles = append(wall.Tiles, base.Card{Color: x, Number: y})
				}
			}
			//大小王
			if config.IsJoker {
				wall.Tiles = append(wall.Tiles, base.Card{Color: PokerColorJoker, Number: PokerPointX})
				wall.Tiles = append(wall.Tiles, base.Card{Color: PokerColorJoker, Number: PokerPointY})
			}
			if config.IsGhost{
				wall.Tiles = append(wall.Tiles, base.Card{Color: PokerColorJoker, Number: PokerPointZ})
			}
			wall.FullCard = append(wall.FullCard, wall.Tiles...)
		}
	}
	wall.Shuffle(seed)
	wall.Shuffle(seed)

	if config.Fixed != ""{
		wall.FixedCard(config.Fixed)
	}
	return wall
}
