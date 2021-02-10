package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/mummy"
)

func main() {

	ebiten.SetWindowSize(608, 512)
	ebiten.SetWindowTitle("MummyGo!")

	game := mummy.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
