package flyweight

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	playerAppearanceFactory := PlayerAppearanceFactory{playerAppearances: make(map[string]*PlayerAppearance)}
	player1 := Player{name: 1}
	player1.setPlayerAppearance("red", "blue", playerAppearanceFactory)
	player2 := Player{name: 2}
	player2.setPlayerAppearance("red", "blue", playerAppearanceFactory)
	player3 := Player{name: 3}
	player3.setPlayerAppearance("green", "blue", playerAppearanceFactory)
	player4 := Player{name: 4}
	player4.setPlayerAppearance("green", "blue", playerAppearanceFactory)
	fmt.Print("Name:", player1.name, " ", player1.getPlayerAppearance().clothesColor, " ", player1.getPlayerAppearance().pantsColor)
	fmt.Printf(" Appearance:%p\n", player1.getPlayerAppearance())
	fmt.Print("Name:", player2.name, " ", player2.getPlayerAppearance().clothesColor, " ", player2.getPlayerAppearance().pantsColor)
	fmt.Printf(" Appearance:%p\n", player2.getPlayerAppearance())
	fmt.Print("Name:", player3.name, " ", player3.getPlayerAppearance().clothesColor, " ", player3.getPlayerAppearance().pantsColor)
	fmt.Printf(" Appearance:%p\n", player3.getPlayerAppearance())
	fmt.Print("Name:", player4.name, " ", player4.getPlayerAppearance().clothesColor, " ", player4.getPlayerAppearance().pantsColor)
	fmt.Printf(" Appearance:%p\n", player4.getPlayerAppearance())

}
