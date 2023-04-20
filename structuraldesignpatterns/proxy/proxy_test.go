package proxy

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	player := &Player{hp: 50, pp: 50, money: 10}
	botProxy := &BotProxy{player: *player}

	fmt.Println(player)
	botProxy.playGame()
	fmt.Println(botProxy.player)
}
