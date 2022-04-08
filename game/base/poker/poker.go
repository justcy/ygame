package poker

import (
	"github.com/justcy/ygame/game/base"
	"strings"
)

// 扑克花色
const (
	PokerColorNIL     int8 = iota // 留空
	PokerColorDiamond             // 方块 1
	PokerColorClub                // 梅花 2
	PokerColorHeart               // 红心 3
	PokerColorSpade               // 黑桃 4
	PokerColorJoker               // 小王、大王、鬼牌 5
)

// 扑克点数
const (
	PokerPointNIL int8 = iota // 留空
	PokerPoint2               //2 - 1
	PokerPoint3               //3 - 2
	PokerPoint4
	PokerPoint5
	PokerPoint6
	PokerPoint7
	PokerPoint8
	PokerPoint9
	PokerPointT
	PokerPointJ
	PokerPointQ
	PokerPointK
	PokerPointA //A -13
	PokerPointX // 小王 -14
	PokerPointY // 大王 -15
	PokerPointZ // 鬼牌 -16
)

type Poker struct{ base.Card }

func (p Poker) toString(cards base.Cards) string {
	colors := make(map[int8]string,5)
	colors[PokerColorDiamond] ="方块"
	colors[PokerColorClub] ="梅花"
	colors[PokerColorHeart] ="红心"
	colors[PokerColorSpade] ="黑桃"
	colors[PokerColorJoker] ="鬼牌"
	numbers := make(map[int8]string,16)
	numbers[PokerPoint2] = "2"
	numbers[PokerPoint3] = "3"
	numbers[PokerPoint4] = "4"
	numbers[PokerPoint5] = "5"
	numbers[PokerPoint6] = "6"
	numbers[PokerPoint7] = "7"
	numbers[PokerPoint8] = "8"
	numbers[PokerPoint9] = "9"
	numbers[PokerPointT] = "10"
	numbers[PokerPointJ] = "J"
	numbers[PokerPointQ] = "Q"
	numbers[PokerPointK] = "K"
	numbers[PokerPointA] = "A"
	numbers[PokerPointX] = "小鬼"
	numbers[PokerPointY] = "大鬼"
	numbers[PokerPointZ] = "鬼牌"
	result := []string{}
	for _, card := range cards {
		if card.Color == PokerColorJoker {
			result = append(result, numbers[card.Number])
		}else{
			result = append(result, colors[card.Color] + numbers[card.Number])
		}
	}
	return strings.Join(result,",")
}

