package mummy

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/mummy/states/credits"
	"github.com/programatta/mummygo/mummy/states/gameplay"
	"github.com/programatta/mummygo/mummy/states/menu"
	"github.com/programatta/mummygo/states"
	"github.com/programatta/mummygo/utils"
)

//Game contiene la funcionalidad de la carga de recursos y de la gestión de
//los estados de juego.
type Game struct {
	spriteSheet  *utils.SpriteSheet
	stateMgr     *states.StateMgr
	currentState states.IState
	nextState    states.IState
}

//NewGame es el constructor.
func NewGame() *Game {
	game := &Game{}
	game.init()
	return game
}

//Update actualiza la lógica del estado en curso (cumple la interface de ebiten).
func (g *Game) Update(screen *ebiten.Image) error {
	if g.nextState != nil {
		g.currentState = g.nextState
		g.nextState = nil
	}

	dt := 0.016
	if ebiten.CurrentFPS() > 0 {
		dt = 1 / ebiten.CurrentFPS()
	}

	g.currentState.ProcessEvents()
	g.currentState.Update(dt)
	g.nextState = g.stateMgr.ChangeState(g.currentState)
	return nil
}

//Draw dibuja el contenido del estado en curso (cumple la interface de ebiten).
func (g *Game) Draw(screen *ebiten.Image) {
	g.currentState.Draw(screen)
}

//Layout (cumple la interface de ebiten)
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 608, 512
}

//init se encarga de inicializar el objeto, cargando recursos y creando la ma-
//quina de estados del juego.
func (g *Game) init() {
	//Cargamos los recursos.
	g.spriteSheet = utils.NewSpriteSheet()
	if err := g.spriteSheet.Load("assets/images/sprites"); err != nil {
		log.Fatal(err)
	}

	//TODO: cargador de fuentes (igual que hemos hecho con el spritesheet)
	fontsloader := utils.NewFontsLoader()
	if err := fontsloader.Load("assets/fonts/ka1.ttf"); err != nil {
		log.Fatal(err)
	}

	if err := fontsloader.Load("assets/fonts/BarcadeBrawl.ttf"); err != nil {
		log.Fatal(err)
	}

	//Preparamos la maquina de estados de la aplicación.
	g.stateMgr = states.NewStateMgr()

	//Añadimos los estados de juego.
	g.stateMgr.AddState("menu", menu.NewMenu(fontsloader))
	g.stateMgr.AddState("credits", credits.NewCredits(fontsloader))
	g.stateMgr.AddState("gameplay", gameplay.NewGamePlay(g.spriteSheet, fontsloader))

	//Establecemos el estado actual.
	g.currentState = g.stateMgr.GetState("menu")
	g.nextState = nil
}

const layoutWidth int = 608
const layoutHeight int = 512
