package mahjong

import (
	"fmt"
	"github.com/justcy/ygame/game/base"
	"testing"
)
func TestPoker(t *testing.T) {
	//wall := ShufflePoker(time.Now().UTC().UnixNano())

	wall := ShuffleMahjong(100,WithFixed("11,12,13,14"),WithDragon(false),WithFlower(true),WithWind(true),WithEightDragon(true),WithGhostDragon(true))
	mj :=Mahjong{}

	fmt.Printf("牌墙长度%d",len(wall.Tiles))
	fmt.Println(mj.toString(wall.Tiles))
}
func TestMahjong_toString(t *testing.T) {
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
			m := Mahjong{
				Card: tt.fields.Card,
			}
			if got := m.toString(tt.args.cards); got != tt.want {
				t.Errorf("toString() = %v, want %v", got, tt.want)
			}
		})
	}
}
