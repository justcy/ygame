package mahjong

import (
	"github.com/justcy/ygame/game/base"
	"strconv"
	"strings"
)

const (
	MjColorCharacters = 1 //万
	MjColorDots       = 2 //筒
	MjColorSticks     = 3 //条
	MjColorWind       = 4 //风
	MjColorDragon     = 5 //中、发、白
	MjColorFlower     = 6 //花牌 春夏秋冬 梅兰竹菊
)

// 扑克点数
const (
	MjPointNIL int8 = iota // 留空
	MjPoint1               //1 - 1
	MjPoint2               //2 - 1
	MjPoint3               //3 - 2
	MjPoint4
	MjPoint5
	MjPoint6
	MjPoint7
	MjPoint8
	MjPoint9
)

type Mahjong struct{ base.Card }

func (m Mahjong) toString(cards base.Cards) string {
	colors := make(map[int8]string,5)
	colors[MjColorCharacters] ="万"
	colors[MjColorDots] ="筒"
	colors[MjColorSticks] ="条"
	numbers := make(map[int8]int,9)
	numbers[MjPoint1] = 1
	numbers[MjPoint2] = 2
	numbers[MjPoint3] = 3
	numbers[MjPoint4] = 4
	numbers[MjPoint5] = 5
	numbers[MjPoint6] = 6
	numbers[MjPoint7] = 7
	numbers[MjPoint8] = 8
	numbers[MjPoint9] = 9
	result := []string{}
	for _, card := range cards {
		if card.Color  == MjColorDragon {
			str := ""
			if numbers[card.Number] == 1{
				str = "红中"
			}else if numbers[card.Number] == 2{
				str = "发财"
			}else if numbers[card.Number] == 3{
				str = "白板"
			}
			result = append(result, str)
		}else if card.Color  == MjColorWind {
			str := []string{"东","南","西","北"}
			result = append(result, str[numbers[card.Number]-1]+"风")
		}else if card.Color  == MjColorFlower {
			str := []string{"春","夏","秋","冬","梅","兰","竹","菊"}
			result = append(result, str[numbers[card.Number]-1])
		}else{
			result = append(result,  strconv.Itoa(numbers[card.Number]) + colors[card.Color])
		}
	}
	return strings.Join(result,",")
}
