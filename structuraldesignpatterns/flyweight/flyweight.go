package flyweight

// Player 玩家
type Player struct {
	name             int
	playerAppearance *PlayerAppearance
}

func (p *Player) getPlayerAppearance() *PlayerAppearance {
	return p.playerAppearance
}

func (p *Player) setPlayerAppearance(clothesColor string, pantsColor string, playerAppearanceFactory PlayerAppearanceFactory) {
	p.playerAppearance = playerAppearanceFactory.getPlayerAppearance(clothesColor, pantsColor)
}

// PlayerAppearance 玩家外观
type PlayerAppearance struct {
	clothesColor string
	pantsColor   string
}

type PlayerAppearanceFactory struct {
	playerAppearances map[string]*PlayerAppearance
}

func (paf *PlayerAppearanceFactory) getPlayerAppearance(clothesColor string, pantsColor string) *PlayerAppearance {
	clothesPantsColor := clothesColor + pantsColor
	if _, ok := paf.playerAppearances[clothesPantsColor]; ok {
		return paf.playerAppearances[clothesPantsColor]
	}
	playerAppearance := &PlayerAppearance{clothesColor: clothesColor, pantsColor: pantsColor}
	paf.playerAppearances[clothesPantsColor] = playerAppearance
	return playerAppearance
}
