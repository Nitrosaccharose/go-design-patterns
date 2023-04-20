package proxy

import "fmt"

type GameAble interface {
	makeMoney()
}

type Player struct {
	hp    int
	pp    int
	money int
	GameAble
}

func (p *Player) makeMoney() {
	fmt.Println("个人赚钱，赚了10个金币")
	p.money += 10
}

type BotProxy struct {
	player Player
	GameAble
}

func (bp *BotProxy) regenerateHealth() {
	fmt.Println("代理回血，回了22点血")
	bp.player.hp += 22
}
func (bp *BotProxy) regenerateMana() {
	fmt.Println("代理回蓝，回了33点蓝")
	bp.player.pp += 33
}

func (bp *BotProxy) playGame() {
	fmt.Println("代理开始游戏")
	bp.regenerateHealth()
	bp.player.makeMoney()
	bp.regenerateMana()
}
