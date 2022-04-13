package texas

import (
	"fmt"
	"github.com/justcy/ygame/game/base"
	"github.com/justcy/ygame/game/base/poker"
	"reflect"
	"sort"
)

type texas struct {
}

// 德州扑克牌型
const (
	TypeNIL           int8 = iota // 留空
	TypeHighCard                  // 高牌 1
	TypeOnePair                   // 一对 2
	TypeTwoPair                   // 两对 3
	TypeThree                     // 三条 4
	TypeStraight                  // 顺子 5
	TypeFlush                     // 同花 6
	TypeFullHouse                 // 葫芦 7
	TypeFour                      // 四条 8
	TypeStraightFlush             // 同花顺 9
	TypeRoyalFlush                // 皇家同花顺 10
)

var maxStraight = []int8{poker.PokerPointT, poker.PokerPointJ, poker.PokerPointQ, poker.PokerPointK, poker.PokerPointA}
var minStraight = []int8{poker.PokerPoint2, poker.PokerPoint3, poker.PokerPoint4, poker.PokerPoint5, poker.PokerPointA}

func (t *texas) MaxCards(cards base.Cards) []base.Card {
	return t.sortCards(cards)[0]
}

func (t *texas) sortCards(cards base.Cards) []base.Cards {
	var r []base.Cards
	if len(cards) <= 5 {
		return append(r, cards)
	}
	tempCards := t.combineCards(cards, 5)
	fmt.Println(tempCards)
	sort.Slice(tempCards, func(i, j int) bool {
		result, _ := t.CompareCards(tempCards[i], tempCards[j])
		fmt.Println(tempCards[i])
		fmt.Println(tempCards[j])
		fmt.Printf("result:%d \n", result)
		return result == -1
	})
	fmt.Println(tempCards)
	return tempCards
}

//CompareCards 比较两个牌型大小 cardsB > cardsA 1   cardsB < cardsA -1   cardsB = cardsA 0
func (t *texas) CompareCards(cardsA, cardsB []base.Card) (int, base.Cards) {
	if cardsA == nil && cardsB != nil {
		return 1, cardsB
	}
	if cardsA != nil && cardsB == nil {
		return -1, cardsA
	}
	if reflect.DeepEqual(cardsA, cardsB) {
		return 0, cardsA
	}
	cardTypeA := t.getCardType(cardsA)
	cardTypeB := t.getCardType(cardsB)
	if cardTypeB > cardTypeA {
		return 1, cardsB
	} else if cardTypeB < cardTypeA {
		return -1, cardsA
	}
	fmt.Printf("相同牌型比较 %d\n", cardTypeA)
	return t.equalCardTypeCompare(cardTypeA, cardsA, cardsB)
}

//combineCards 组合所有牌型
func (t *texas) combineCards(cards base.Cards, num int) []base.Cards {
	var r []base.Cards
	//if len(cards) <= 5 {
	//	return append(r, cards)
	//}
	n := len(cards)
	if n < num || num <= 0 {
		return r
	}
	for index, card := range cards {
		if num == 1 {
			temp := base.Cards{card}
			r = append(r, temp)
		} else {
			c := t.combineCards(cards[index+1:], num-1)
			for _, val := range c {
				temp := base.Cards{card}
				r = append(r, append(temp, val...))
			}
		}
	}
	return r
}

func (t *texas) getCardType(b []base.Card) int8 {
	countByNumber := base.CountCardByNumber(b)
	fmt.Println(countByNumber)
	if len(b) <= 2 {
		if len(countByNumber[2]) == 1 {
			return TypeOnePair
		}
		return TypeHighCard
	}
	isStraight := t.isStraight(countByNumber[1])
	isFlush := t.isFlush(countByNumber[1])
	if isStraight {
		isRoyalFlush := t.isRoyalFlush(countByNumber[1])
		if isRoyalFlush {
			return TypeRoyalFlush
		} else if isFlush {
			return TypeStraightFlush
		}
		return TypeStraight
	} else {
		if isFlush {
			return TypeFlush
		} else {
			if len(countByNumber[1]) == len(b) {
				return TypeHighCard
			} else if len(countByNumber[2]) == 1 && len(countByNumber[3]) == 0 {
				return TypeOnePair
			} else if len(countByNumber[2]) == 2 {
				return TypeTwoPair
			} else if len(countByNumber[3]) == 1 && len(countByNumber[1]) == 2 {
				return TypeThree
			} else if len(countByNumber[3]) == 1 && len(countByNumber[2]) == 1 {
				return TypeFullHouse
			} else if len(countByNumber[4]) == 1 {
				return TypeFour
			}
		}
	}
	return TypeHighCard
}

func (t *texas) isRoyalFlush(number []base.Card) bool {
	if len(number) < 5 {
		return false
	}
	_, temp := base.GetCardsColorsAndNumbers(base.SortCards(number))
	if reflect.DeepEqual(maxStraight, temp) {
		return true
	}
	return false
}
func (t *texas) isFlush(number []base.Card) bool {
	if len(number) < 5 {
		return false
	}
	color, _ := base.GetCardsColorsAndNumbers(number)
	return len(base.Unique(color)) == 1
}
func (t *texas) isStraight(number []base.Card) bool {
	if len(number) < 5 {
		return false
	}
	_, temp := base.GetCardsColorsAndNumbers(base.SortCards(number))
	if reflect.DeepEqual(minStraight, temp) {
		fmt.Println("最小顺子")
		return true
	}
	for i, key := range temp {
		next := i + 1
		if next < len(temp) && temp[next] != 0 && temp[next]-key != 1 {
			return false
		}
	}
	return true
}
func (t *texas) equalCardTypeCompare(cardType int8, a []base.Card, b []base.Card) (int, base.Cards) {
	countByNumberA := base.CountCardByNumber(a)
	countByNumberB := base.CountCardByNumber(b)
	switch cardType {
	case TypeRoyalFlush:
		return 0, a
	case TypeStraightFlush:
	case TypeStraight:
		_, numA := base.GetCardsColorsAndNumbers(base.SortCards(countByNumberA[1]))
		_, numB := base.GetCardsColorsAndNumbers(base.SortCards(countByNumberB[1]))

		if numA[0] == numB[0] {
			return 0, a
		} else {
			if reflect.DeepEqual(numB, minStraight) || numB[4] < numA[4] {
				return -1, a
			}
			if reflect.DeepEqual(numA, minStraight) || numB[4] > numA[4] {
				return 1, b
			}
		}
	case TypeThree:
		if countByNumberA[3][0].Number == countByNumberB[3][0].Number {
			//比较单牌大小
			oneCompare := t.compareOneCards(countByNumberA[1], countByNumberB[1])
			if -1 == oneCompare {
				return -1, a
			} else if 1 == oneCompare {
				return 1, b
			}
			return 0, a
		} else if countByNumberA[3][0].Number < countByNumberB[3][0].Number {
			return 1, b
		}
		return -1, a
	case TypeFullHouse:
		if countByNumberA[3][0].Number == countByNumberB[3][0].Number {
			if countByNumberA[2][0].Number == countByNumberB[2][0].Number {
				return 0, a
			} else if countByNumberA[2][0].Number < countByNumberB[2][0].Number {
				return 1, b
			}
			return -1, a
		} else if countByNumberA[3][0].Number < countByNumberB[3][0].Number {
			return 1, b
		}
		return -1, a
	case TypeFour:
		if countByNumberA[4][0].Number == countByNumberB[4][0].Number {
			//比较单牌大小
			oneCompare := t.compareOneCards(countByNumberA[1], countByNumberB[1])
			if -1 == oneCompare {
				return -1, a
			} else if 1 == oneCompare {
				return 1, b
			}
			return 0, a
		} else if countByNumberA[4][0].Number < countByNumberB[4][0].Number {
			return 1, b
		}
		return -1, a
	case TypeFlush:
	case TypeHighCard:
		oneCompare := t.compareOneCards(countByNumberA[1], countByNumberB[1])
		if -1 == oneCompare {
			return -1, a
		} else if 1 == oneCompare {
			return 1, b
		}
		return 0, a
	case TypeTwoPair:
		twoCompare := t.compareOneCards(countByNumberA[2], countByNumberB[2])
		if -1 == twoCompare {
			//比较单牌大小
			oneCompare := t.compareOneCards(countByNumberA[1], countByNumberB[1])
			if -1 == oneCompare {
				return -1, a
			} else if 1 == oneCompare {
				return 1, b
			}
			return 0, a
		} else if 1 == twoCompare {
			return 1, b
		}
		return 0, a
	case TypeOnePair:
		if countByNumberA[2][0].Number == countByNumberB[2][0].Number {
			//比较单牌大小
			oneCompare := t.compareOneCards(countByNumberA[1], countByNumberB[1])
			if -1 == oneCompare {
				return -1, a
			} else if 1 == oneCompare {
				return 1, b
			}
			return 0, a
		} else if countByNumberA[2][0].Number < countByNumberB[2][0].Number {
			return 1, b
		}
		return -1, a
	}
	return 0, a
}

func (t *texas) compareOneCards(a base.CardVec, b base.CardVec) int {
	fmt.Println("参数A")
	fmt.Println(a)
	fmt.Println("参数B")
	fmt.Println(b)
	if len(a) == 0 && len(b) == 0 {
		return 0
	}

	_, numA := base.GetCardsColorsAndNumbers(base.SortCards(a))
	_, numB := base.GetCardsColorsAndNumbers(base.SortCards(b))
	for i := len(numB) - 1; i > 0; i-- {
		if numB[i] > numA[i] {
			return 1
		} else if numB[i] < numA[i] {
			return -1
		}
	}
	//var numA, numB []int
	//for _, cardA := range a {
	//	numA = append(numA, int(cardA.Number))
	//}
	//for _, cardB := range b {
	//	numB = append(numB, int(cardB.Number))
	//}
	//sort.Sort(sort.Reverse(sort.IntSlice(numA)))
	//sort.Sort(sort.Reverse(sort.IntSlice(numB)))
	//for i, v := range numB {
	//	if v > numA[i] {
	//		return 1
	//	} else if v < numA[i] {
	//		return -1
	//	}
	//}
	return 0
}
