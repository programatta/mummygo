package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/programatta/mummygo/mummy"
)

func main() {
	//Preparamos la pantalla del juego.
	ebiten.SetWindowSize(608, 512)
	ebiten.SetWindowTitle("MummyGo!")
	ebiten.SetCursorMode(ebiten.CursorModeHidden)

	//Cargamos el icono de la aplicación.
	pngfile := "assets/images/ic_launcher_round.png"
	_, imgicon, _ := ebitenutil.NewImageFromFile(pngfile, ebiten.FilterDefault)
	imgIcons := make([]image.Image, 0)
	imgIcons = append(imgIcons, imgicon)
	ebiten.SetWindowIcon(imgIcons)

	game := mummy.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
