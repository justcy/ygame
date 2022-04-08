package base

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Wall struct {
	Tiles    Cards // 所有牌
	FullCard Cards // 全牌花型一样一张
	Forward  int32  // 前游标
	Backward int32  // 后游标
}
//按照给定规则摆出牌型
func (wall *Wall) FixedCard(r string){
	if len(r) <= 0 {
		return
	}
	a := strings.Split(r, ",")
	if len(a) < 1 {
		return
	}
	head := make([]Card, 0)
	foot := wall.Tiles
	for _, v := range a {
		i, err := strconv.Atoi(v)
		if err == nil && i >= 11 {
			var x, y int8
			if i >= 11 && i <= 99 {
				x = int8(i / 10)
				y = int8(i % 10)
			} else if i >= 100 {
				x = int8(i / 100)
				y = int8(i % 100)
			}
			//log.Debugf("Tiles: %d C:%d N:%d", i, x, y)
			c := Card{Color: x, Number: y}
			if CardInSlice(c, foot) { //包含牌,才摆牌
				head = append(head, c)
				foot = sliceDel(foot, c)
			} else {
				//log.Debugf("bai not find %d:%d", c.Color, c.Number)
			}
		} else if err == nil && i < 0 {
			// lastFoot := len(foot)
			i = i * -1
			for j := 0; j < i; j++ {
				c := Card{Color: 0, Number: 0}
				head = append(head, c)
			}
			// log.Debugf("随机摆牌 %d 张 last:%d now:%d", i, len(foot))
		} else {
			fmt.Printf("bai not find I:%d e:%v", i, err)
		}
	}
	for i := range head { //将空白补上
		if head[i].Color == 0 && head[i].Number == 0 && len(foot) > 1 {
			head[i] = foot[0]
			foot = foot[1:]
		}
	}
	// log.Debugf("baipai: L:%d ", len(head))
	foot = ShuffleSliceCard(time.Now().UTC().UnixNano(),foot)
	wall.Tiles = append(head, foot...)
}
// Shuffle 洗牌
func (wall *Wall) Shuffle(seed int64) {
	wall.Tiles = ShuffleSliceCard(seed,wall.Tiles)
}
//全部牌
func (wall *Wall) GetAll() string {
	result := ""
	for _, v := range wall.Tiles {
		result = result + fmt.Sprintf("%d%d,", v.Color, v.Number)
	}
	if len(result) <= 1 {
		fmt.Println("GetAllPai Result Is Empty")
		return ""
	}
	li := len(result) - 1
	return result[0:li]
}
//Draw 发牌
func (wall *Wall) Draw(s int32) Cards {
	if len(wall.Tiles) >= int(s) {
		v := wall.Tiles[:s]
		wall.Tiles = wall.Tiles[s:]
		return v
	} else if len(wall.Tiles) > 0 {
		return wall.Tiles
	}else{
		return Cards{Card{Color: 0, Number: 0}}
	}
}
func (wall *Wall) DrawOne() Card {
	cards := wall.Draw(1)
	return cards[0]
}
func (wall *Wall) Length() int {
	return len(wall.Tiles)
}
