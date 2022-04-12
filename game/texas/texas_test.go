package texas

import (
	"github.com/justcy/ygame/game/base"
	"github.com/justcy/ygame/game/base/poker"
	"reflect"
	"testing"
)

func Test_texas_CompareCards(t1 *testing.T) {
	type args struct {
		cardsA []base.Card
		cardsB []base.Card
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 base.Cards
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			got, got1 := t.CompareCards(tt.args.cardsA, tt.args.cardsB)
			if got != tt.want {
				t1.Errorf("CompareCards() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("CompareCards() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_texas_MaxCards(t1 *testing.T) {
	type args struct {
		cards base.Cards
	}
	tests := []struct {
		name string
		args args
		want []base.Card
	}{
		// TODO: Add test cases.
		{name: "高牌", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint7},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint7},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
		{name: "顺子", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint7},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
			}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			if got := t.MaxCards(tt.args.cards); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("MaxCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_texas_combineCards(t1 *testing.T) {
	type args struct {
		cards base.Cards
		num   int
	}
	tests := []struct {
		name string
		args args
		want []base.Cards
	}{
		// TODO: Add test cases.
		{name: "两张牌", args: args{cards: base.Cards{base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2}, base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2}}, num: 5}, want: []base.Cards{base.Cards{base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2}, base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2}}}},
		{name: "5张牌", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint2},
			}, num: 5},
			want: []base.Cards{
				base.Cards{
					base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
					base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
					base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint4},
					base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint5},
					base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint2},
				},
			}},
		{name: "6张牌", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint7},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
			}, num: 5},
			want: []base.Cards{
				base.Cards{
					base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
					base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
					base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
					base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
					base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint7},
				},
			}},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			if got := t.combineCards(tt.args.cards, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("combineCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_texas_compareOneCards(t1 *testing.T) {
	type args struct {
		a base.CardVec
		b base.CardVec
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			if got := t.compareOneCards(tt.args.a, tt.args.b); got != tt.want {
				t1.Errorf("compareOneCards() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_texas_equalCardTypeCompare(t1 *testing.T) {
	type args struct {
		cardType int8
		a        []base.Card
		b        []base.Card
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 base.Cards
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			got, got1 := t.equalCardTypeCompare(tt.args.cardType, tt.args.a, tt.args.b)
			if got != tt.want {
				t1.Errorf("equalCardTypeCompare() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t1.Errorf("equalCardTypeCompare() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_texas_getCardType(t1 *testing.T) {
	type args struct {
		b []base.Card
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			if got := t.getCardType(tt.args.b); got != tt.want {
				t1.Errorf("getCardType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_texas_isRoyalFlush(t1 *testing.T) {
	type args struct {
		number map[int]base.CardVec
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			if got := t.isRoyalFlush(tt.args.number[1]); got != tt.want {
				t1.Errorf("isRoyalFlush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_texas_isStraight(t1 *testing.T) {
	type args struct {
		number map[int]base.CardVec
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			if got := t.isStraight(tt.args.number[1]); got != tt.want {
				t1.Errorf("isStraight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_texas_sortCards(t1 *testing.T) {
	type args struct {
		cards base.Cards
	}
	tests := []struct {
		name string
		args args
		want []base.Cards
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &texas{}
			if got := t.sortCards(tt.args.cards); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("sortCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
