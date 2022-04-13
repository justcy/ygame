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
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPointK},
			}},
			want: []base.Card{
				//{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint7},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
				{Color: poker.PokerColorSpade, Number: poker.PokerPointK},
			}},
		{name: "一对", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint7},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPointA},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint7},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPointA},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
		{name: "2对", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPointA},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint6},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint6},
				{Color: poker.PokerColorSpade, Number: poker.PokerPointA},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
		{name: "三条", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint7},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPointA},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPointA},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint7},
				{Color: poker.PokerColorClub, Number: poker.PokerPointA},
				{Color: poker.PokerColorSpade, Number: poker.PokerPointA},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
		{name: "顺子", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint5},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint6},
			}},
		{name: "同花", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint8},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint8},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
		{name: "葫芦", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				{Color: poker.PokerColorHeart, Number: poker.PokerPoint5},
			}},
		{name: "四条", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorClub, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint9},
				{Color: poker.PokerColorHeart, Number: poker.PokerPointA},
			}},
		{name: "同花顺", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
			}},
			{name: "皇家同花顺", args: args{
			cards: base.Cards{
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointT},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPoint8},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointJ},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointQ},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
				base.Card{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}},
			want: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointT},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointJ},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointQ},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
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
		{name: "比较牌型，len(a) != len(b)", args: args{a: base.CardVec{}, b: base.CardVec{
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
		}}, want: 0},
		{name: "比较牌型，两个都为空", args: args{a: base.CardVec{}, b: base.CardVec{}}, want: 0},
		{name: "比较牌型，A>B", args: args{a: base.CardVec{
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPointJ},
		}, b: base.CardVec{
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint8},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPointT},
		}},
			want: -1,
		},
		{name: "比较牌型，A<B", args: args{a: base.CardVec{
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPointJ},
		}, b: base.CardVec{
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint8},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPointQ},
		}},
			want: 1,
		},
		{name: "比较牌型，A=B", args: args{a: base.CardVec{
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPointJ},
		}, b: base.CardVec{
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint2},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
			base.Card{Color: poker.PokerColorHeart, Number: poker.PokerPointJ},
		}},
			want: 0,
		},
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
		{name: "高牌-2张", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
			}}, want: TypeHighCard},
		{name: "高牌", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint7},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
			}}, want: TypeHighCard},
		{name: "一对-2张", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint2},
			}}, want: TypeOnePair},
		{name: "一对", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint2},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
			}}, want: TypeOnePair},
		{name: "两对", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}}, want: TypeTwoPair},
		{name: "三条", args: args{
			b: []base.Card{
				{Color: poker.PokerColorClub, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint7},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
			}}, want: TypeThree},
		{name: "顺子", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint6},
			}}, want: TypeStraight},
		{name: "最小顺子", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorClub, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				{Color: poker.PokerColorSpade, Number: poker.PokerPointA},
			}}, want: TypeStraight},
		{name: "葫芦", args: args{
			b: []base.Card{
				{Color: poker.PokerColorClub, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPointK},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
			}}, want: TypeFullHouse},
		{name: "四条", args: args{
			b: []base.Card{
				{Color: poker.PokerColorClub, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint9},
				{Color: poker.PokerColorHeart, Number: poker.PokerPoint9},
				{Color: poker.PokerColorSpade, Number: poker.PokerPoint9},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
			}}, want: TypeFour},
		{name: "同花顺", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint2},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint3},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint4},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint5},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPoint6},
			}}, want: TypeStraightFlush},
		{name: "皇家同花顺", args: args{
			b: []base.Card{
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointT},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointJ},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointQ},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointK},
				{Color: poker.PokerColorDiamond, Number: poker.PokerPointA},
			}}, want: TypeRoyalFlush},
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
