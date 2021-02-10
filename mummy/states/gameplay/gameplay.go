package gameplay

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/programatta/mummygo/mummy/states/gameplay/enemies"
	"github.com/programatta/mummygo/mummy/states/gameplay/object"
	"github.com/programatta/mummygo/mummy/states/gameplay/player"
	"github.com/programatta/mummygo/mummy/states/gameplay/stage"
	"github.com/programatta/mummygo/states"
	"github.com/programatta/mummygo/utils"
)

//GamePlay contiene la funcionalidad del juego.
type GamePlay struct {
	nextStateID      string
	spriteSheet      *utils.SpriteSheet
	stage            *stage.Stage
	mummies          []*enemies.Mummy
	objects          []*object.CollectableObject
	player           *player.Player
	isGameOver       bool //estado
	isNextLevel      bool //estado
	playerLeaveLevel bool //estado
	uigame           *UIGame
	level            int
}

//NewGamePlay es el constructor
func NewGamePlay(spriteSheet *utils.SpriteSheet) states.IState {
	g := &GamePlay{}

	g.nextStateID = "gameplay"
	g.spriteSheet = spriteSheet

	g.prepareLevel()

	//Create a UI.
	g.uigame = NewUIGame()

	//Create a stage.
	g.stage = stage.NewStage(g.spriteSheet, g)

	//Create a mummies.
	g.mummies = make([]*enemies.Mummy, 0)

	//Create a collectable objects
	g.objects = make([]*object.CollectableObject, 0)

	//Creates the player.
	w, h := ebiten.WindowSize()
	g.player = player.NewPlayer(g.spriteSheet, g.stage)
	g.player.SetPosition((w-64)/2+16, (h-32)/2-16)
	g.player.SetLives(3)

	return g
}

/*===========================================================================*/
/*                               Interface IState                            */
/*===========================================================================*/

//Init ...
func (g *GamePlay) Init() {
	g.nextStateID = "gameplay"
}

//ProcessEvents procesa los eventos del juego.
func (g *GamePlay) ProcessEvents() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		//os.Exit(0) //TODO: tenemos que ir a otro estado.
		g.nextStateID = "menu"
	}

	//Para debug.
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		g.stage.DebugOpenTombs()
		g.isNextLevel = true
	}

	//Movimiento del player.
	if !g.playerLeaveLevel {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			g.player.Move(player.PlayerLeft)
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			g.player.Move(player.PlayerRight)
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			g.player.Move(player.PlayerUp)
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			g.player.Move(player.PlayerDown)
		}
	}
}

//Update actualiza la lógica del juego.
func (g *GamePlay) Update(dt float64) {

	g.stage.Update(dt)
	g.player.Update(dt)
	if g.isNextLevel {
		x, y := g.player.Position()
		xLog := int(x+16) / 32
		yLog := int(y+16) / 32
		if xLog == 9 && yLog == 1 && g.player.CurrentDir() == "up" && !g.playerLeaveLevel {
			g.playerLeaveLevel = true
			g.player.LeaveLevel(g)
		}
	}

	//Mummies
	if len(g.mummies) > 0 {
		mummiesTmp := make([]*enemies.Mummy, len(g.mummies))
		copied := 0
		for _, mummy := range g.mummies {
			if g.checkPlayerIsAttackedByMummy(g.player, mummy) {
				mummy = nil
				//TODO: Verificar si tenemos pociones
				if g.player.Potions() > 0 {
					g.player.ConsumePotion()
				} else {
					g.player.LostLive()
					g.isGameOver = g.player.Lives() == 0
					if !g.isGameOver {
						w, h := ebiten.WindowSize()
						g.player.SetPosition((w-64)/2+16, (h-32)/2-16)
					}
				}
			} else {
				mummy.Update(dt)
				mummiesTmp[copied] = mummy
				copied++
			}
		}
		g.mummies = mummiesTmp[:copied]
	}

	//Collectable Objects.
	if len(g.objects) > 0 {
		objectsTmp := make([]*object.CollectableObject, len(g.objects))
		copied := 0
		for _, object := range g.objects {
			if g.checkCanPickUpObject(g.player, object) {
				g.player.AddObject(object)
				if g.player.HasKeyAndPapyre() {
					g.stage.OpenMainDoor()
					g.isNextLevel = true
				}
				object = nil
			} else {
				object.Update(dt)
				objectsTmp[copied] = object
				copied++
			}
		}
		g.objects = objectsTmp[:copied]
	}

	g.uigame.SetLives(g.player.Lives())
	g.uigame.SetPotions(g.player.Potions())
	g.uigame.SetLevel(g.level)
}

//Draw draws the game.
func (g *GamePlay) Draw(screen *ebiten.Image) {
	if g.isGameOver {
		return
	}

	g.stage.Draw(screen)

	for _, object := range g.objects {
		object.Draw(screen)
	}

	g.player.Draw(screen)

	for _, mummy := range g.mummies {
		mummy.Draw(screen)
	}

	g.uigame.Draw(screen)
}

//NextState ...
func (g *GamePlay) NextState() string {
	return g.nextStateID
}

/*===========================================================================*/
/*                      Interface IGamePlayNotificable                       */
/*===========================================================================*/

//OnCreateObject crea un objeto por tipo que contenia la tumba.
func (g *GamePlay) OnCreateObject(t, x, y int) {
	switch t {
	case 1: //Mummy
		mummy := enemies.NewMummy(g.spriteSheet, x, y)
		g.mummies = append(g.mummies, mummy)
		break
	case 2, 3, 4: //Potion, Key or Papyre
		object := object.NewCollectableObject(g.spriteSheet, t, x, y)
		g.objects = append(g.objects, object)
		break
	case 5: //Wizard
		break
	}
}

//OnPrepreNewLevel indica que el player ha abandonado el nivel por la puerta
//principal y procedemos a preparar otro nivel.
func (g *GamePlay) OnPrepreNewLevel() {
	//TODO: de momento para ver que funciona el ciclo.
	g.isGameOver = true

	//cargar un nuevo nivel.
	g.prepareLevel()
}

/*===========================================================================*/
/*                               Private Section                             */
/*===========================================================================*/

func (g *GamePlay) checkCanPickUpObject(player *player.Player, itemObject *object.CollectableObject) bool {

	iox, ioy := itemObject.Position()
	xObjLog := int(iox+16) / 32
	yObjLog := int(ioy+16) / 32

	xp, yp := player.Position()
	xLog := int(xp+16) / 32
	yLog := int(yp+16) / 32

	return (xObjLog == xLog) && (yObjLog == yLog)

}

func (g *GamePlay) checkPlayerIsAttackedByMummy(player *player.Player, mummy *enemies.Mummy) bool {
	if player.IsBlinking() {
		return false
	}

	mx, my := mummy.Position()
	xMummyLog := int(mx+16) / 32
	yMummyLog := int(my+16) / 32

	xp, yp := player.Position()
	xLog := int(xp+16) / 32
	yLog := int(yp+16) / 32

	return (xMummyLog == xLog) && (yMummyLog == yLog)
}

func (g *GamePlay) prepareLevel() {
	g.level++

	//TODO: cargar datos para el nivel.
}
