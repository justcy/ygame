package poker

import (
	"fmt"
	"github.com/justcy/ygame/game/base"
	"testing"
)

func TestPoker(t *testing.T) {
	//wall := ShufflePoker(time.Now().UTC().UnixNano())

	wall := ShufflePoker(100,WithRepeat(1), WithJoker(true), WithGhost(true),WithPokerWallType(1),WithFixed("11,12,13,14,116"))
	poker :=Poker{}

	fmt.Printf("牌墙长度%d",len(wall.Tiles))
	fmt.Println(poker.toString(wall.Tiles))
}
func TestPoker_toString(t *testing.T) {
	type fields struct {
		Card base.Card
	}
	type args struct {
		cards base.Cards
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Poker{
				Card: tt.fields.Card,
			}
			if got := p.toString(tt.args.cards); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
