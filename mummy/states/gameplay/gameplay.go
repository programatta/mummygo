package gameplay

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/programatta/mummygo/mummy/states/gameplay/enemies"
	"github.com/programatta/mummygo/mummy/states/gameplay/gamelevel"
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
	fontsloader      *utils.FontsLoader
	soundmgr         *utils.SoundMgr
	stage            *stage.Stage
	enemies          []enemies.IEnemy
	objects          []*object.CollectableObject
	player           *player.Player
	isNextLevel      bool //estado
	playerLeaveLevel bool //estado
	uigame           *UIGame
	currentLevel     int
	score            int
	currentState     tgamePlayState
	nextState        tgamePlayState
	chgameover       chan bool
	alpha            float64
	goimgblack       *ebiten.Image
	levels           []gamelevel.GameLevel
	oxygenConsumed   float64
}

//NewGamePlay es el constructor
func NewGamePlay(spriteSheet *utils.SpriteSheet, fontsloader *utils.FontsLoader, soundmgr *utils.SoundMgr) states.IState {
	g := &GamePlay{}

	g.nextStateID = "gameplay"
	g.spriteSheet = spriteSheet
	g.fontsloader = fontsloader
	g.soundmgr = soundmgr

	//Creamos el UI del juego (TODO: colocar iconos)
	g.uigame = NewUIGame(fontsloader, g.spriteSheet)

	//Creamos el escenario.
	g.stage = stage.NewStage(g.spriteSheet, g.soundmgr, g)

	//Creamos el array de objetos a recoger.
	g.objects = make([]*object.CollectableObject, 0)

	//Creamos el array de enemigos.
	g.enemies = make([]enemies.IEnemy, 0)

	//Creamos al jugador.
	g.player = player.NewPlayer(g.spriteSheet, g.soundmgr, g.stage)
	return g
}

/*===========================================================================*/
/*                               Interface IState                            */
/*===========================================================================*/

//Init ...
func (g *GamePlay) Init() {
	g.nextStateID = "gameplay"
	g.chgameover = nil

	g.prepareLevel(true)
}

//ProcessEvents procesa los eventos del juego.
func (g *GamePlay) ProcessEvents() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.nextStateID = "menu"
		g.nextState = exit
	}

	//Para debug.
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		g.stage.DebugOpenTombs()
		g.isNextLevel = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.player.LostLive()
		g.nextState = gameover
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
	ambiencePlayer := g.soundmgr.Sound("game.wav")
	if ambiencePlayer.IsPlaying() {
		ambiencePlayer.Play()
	} else {
		ambiencePlayer.Rewind()
		ambiencePlayer.Play()
	}

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
	if len(g.enemies) > 0 {
		enemiesTmp := make([]enemies.IEnemy, len(g.enemies))
		copied := 0
		for _, enemy := range g.enemies {
			if g.checkPlayerIsAttackedByEnemy(g.player, enemy) {
				if g.player.Potions() > 0 {
					g.player.ConsumePotion()
					enemy.Death()
					//Matamos una momia 125 puntos.
					g.score += 125
				} else {
					switch enemy.(type) {
					case *enemies.Mummy:
						g.player.LostLive()
						isGameOver := g.player.Lives() == 0
						if !isGameOver {
							w, h := ebiten.WindowSize()
							g.player.SetPosition((w-64)/2+16, (h-32)/2-16)
							g.oxygenConsumed = 0
						} else {
							g.nextState = gameover
						}
						break
					case *enemies.Spell:
						g.player.Bewitched()
						break
					}
				}
				enemy = nil
			} else {
				enemy.Update(dt)
				enemiesTmp[copied] = enemy
				copied++
			}
		}
		g.enemies = enemiesTmp[:copied]
	}

	//Collectable Objects.
	if len(g.objects) > 0 {
		objectsTmp := make([]*object.CollectableObject, len(g.objects))
		copied := 0
		for _, object := range g.objects {
			if g.checkCanPickUpObject(g.player, object) {
				//Por cada objeto cogido damos puntos.
				if object.TypeObject() == 3 || object.TypeObject() == 4 {
					//Por la llave y el papiro 100 puntos.
					g.score += 100
				} else {
					//las pociones 35
					g.score += 35
				}
				object.PickedUp()

				g.player.AddObject(object)
				if g.player.HasKey() && g.player.HasPapyre() {
					g.stage.OpenMainDoor()
					g.isNextLevel = true
					//Por completar el nivel incrementamos 500 puntos.
					g.score += 500
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

	//Consumo de oxigeno.
	warnlevel := 0
	g.oxygenConsumed += dt
	if g.oxygenConsumed >= 100 {
		g.player.LostLive()
		isGameOver := g.player.Lives() == 0
		if !isGameOver {
			w, h := ebiten.WindowSize()
			g.player.SetPosition((w-64)/2+16, (h-32)/2-16)
			g.oxygenConsumed = 0
		} else {
			g.nextState = gameover
			g.oxygenConsumed = 100
		}
	} else {
		if g.oxygenConsumed > 50 {
			warnlevel = 1
		}
		if g.oxygenConsumed > 75 {
			warnlevel = 2
		}
	}

	g.uigame.SetLives(g.player.Lives())
	g.uigame.SetPotions(g.player.Potions())
	g.uigame.SetLevel(g.currentLevel)
	g.uigame.SetScore(g.score)
	g.uigame.SetKey(g.player.HasKey())
	g.uigame.SetPapyre(g.player.HasPapyre())
	g.uigame.SetPercentOxigen(g.oxygenConsumed, warnlevel)

	if g.currentState == gameover {
		if g.alpha < 1 {
			g.alpha += dt
		} else {
			g.alpha = 1
			gos := g.soundmgr.Sound("gameover.wav")
			if !gos.IsPlaying() && g.chgameover == nil {
				g.chgameover = make(chan bool)
				go func(ch chan bool) {
					gos.Rewind()
					gos.Play()
					for gos.IsPlaying() {
						ch <- true
					}
					ch <- false
					close(ch)
				}(g.chgameover)
			} else {
				isPlaying, _ := <-g.chgameover
				if !isPlaying {
					g.nextState = exit
					g.nextStateID = "menu"
				}
			}
		}
	}
	if g.currentState == nextlevel {
		if g.alpha < 1 {
			g.alpha += dt
		} else {
			g.alpha = 1
			//cargar un nuevo nivel.
			ch := make(chan bool)
			go func(ch chan bool) {
				levelup := g.soundmgr.Sound("levelup.wav")
				levelup.Rewind()
				levelup.Play()
				ch <- true
			}(ch)

			<-ch
			close(ch)
			g.prepareLevel(false)
		}
	}
	if g.currentState == victorylevel {
		if g.alpha < 1 {
			g.alpha += dt
		} else {
			g.alpha = 1
			wgs := g.soundmgr.Sound("wingame.wav")
			if !wgs.IsPlaying() && g.chgameover == nil {

				//Eliminamos todos los enemigos.
				g.enemies = make([]enemies.IEnemy, 0)

				g.chgameover = make(chan bool)
				go func(ch chan bool) {
					wgs.Rewind()
					wgs.Play()
					for wgs.IsPlaying() {
						ch <- true
					}
					ch <- false
					close(ch)
				}(g.chgameover)
			} else {
				isPlaying, _ := <-g.chgameover
				if !isPlaying {
					g.nextState = exit
					g.nextStateID = "menu"
				}
			}
		}
	}
	if g.currentState == exit {
		if g.alpha < 1 {
			g.alpha += dt
		} else {
			g.alpha = 1
			g.nextState = end
		}
	}
}

//Draw dibuja los elementos del juego.
func (g *GamePlay) Draw(screen *ebiten.Image) {
	g.stage.Draw(screen)
	switch g.currentState {
	case playing:
		for _, object := range g.objects {
			object.Draw(screen)
		}

		g.player.Draw(screen)

		for _, enemy := range g.enemies {
			enemy.Draw(screen)
		}

		g.uigame.Draw(screen)
		break

	case gameover:
		g.fadeToBlack(screen)
		if g.alpha == 1 {
			uistring := fmt.Sprint("GAME OVER")
			fontSize := 18
			screenWidth, screenHeight := screen.Size()
			x := (screenWidth - len(uistring)*fontSize) / 2
			y := (screenHeight - fontSize) / 2
			font := g.fontsloader.GetFont("BarcadeBrawl.ttf", 72, 18)
			text.Draw(screen, uistring, font, x, y, color.White)
		}
		break

	case nextlevel:
		fallthrough
	case exit:
		fallthrough
	case end:
		g.fadeToBlack(screen)
		break

	case victorylevel:
		g.fadeToBlack(screen)
		if g.alpha == 1 {
			uistring := fmt.Sprint("YOU WIN")
			uiscore := fmt.Sprintf("YOUR SCORE:%d", g.score)
			fontSize := 18
			screenWidth, screenHeight := screen.Size()

			//texto: you win!
			x := (screenWidth - len(uistring)*fontSize) / 2
			y := (screenHeight - fontSize) / 2
			font := g.fontsloader.GetFont("BarcadeBrawl.ttf", 72, 18)
			text.Draw(screen, uistring, font, x, y, color.White)

			//texto: your score xxxx.
			x2 := (screenWidth - len(uiscore)*fontSize) / 2
			y2 := (screenHeight + fontSize*2) / 2
			text.Draw(screen, uiscore, font, x2, y2, color.White)
		}

		break
	}
}

//NextState ...
func (g *GamePlay) NextState() string {
	if g.currentState != g.nextState {
		g.currentState = g.nextState
	}

	if g.currentState == end {
		return g.nextStateID
	}
	return "gameplay"
}

//End ...
func (g *GamePlay) End() {
	//Paramos la musica ambiente.
	ambiencePlayer := g.soundmgr.Sound("game.wav")
	ambiencePlayer.Pause()
	ambiencePlayer.Rewind()

	//Eliminamos los enemigos para parar los sonidos.
	g.enemies = make([]enemies.IEnemy, 0)
}

/*===========================================================================*/
/*                      Interface IGamePlayNotificable                       */
/*===========================================================================*/

//OnCreateObject crea un objeto por tipo que contenia la tumba.
func (g *GamePlay) OnCreateObject(t, x, y int) {
	switch t {
	case 1: //Mummy
		mummy := enemies.NewMummy(g.spriteSheet, g.soundmgr, x, y, g)
		g.enemies = append(g.enemies, mummy)
		break
	case 2, 3, 4: //Potion, Key or Papyre
		object := object.NewCollectableObject(g.spriteSheet, g.soundmgr, t, x, y)
		g.objects = append(g.objects, object)
		break
	case 5: //Wizard
		spell := enemies.NewSpell(g.spriteSheet, g.soundmgr, x, y, g)
		g.enemies = append(g.enemies, spell)
		break
	}

	//Por cada tumba abierta incrementamos 20 puntos.
	g.score += 20
}

//OnPrepreNewLevel indica que el player ha abandonado el nivel por la puerta
//principal y procedemos a preparar otro nivel o finalizar el juego ya que ha
//alcanzado el máximo nivel.
func (g *GamePlay) OnPrepreNewLevel() {
	if g.currentLevel == len(g.levels) {
		g.nextState = victorylevel
	} else {
		g.nextState = nextlevel
	}
}

//OnRequestPlayerPosition devuelve la posición del player solicitado por un
//enemigo.
func (g *GamePlay) OnRequestPlayerPosition() (float64, float64) {
	return g.player.Position()
}

//OnScoreByStep establece puntos por cada paso dado en celdas vacias.
func (g *GamePlay) OnScoreByStep() {
	g.score += 5
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

func (g *GamePlay) checkPlayerIsAttackedByEnemy(player *player.Player, enemy enemies.IEnemy) bool {
	if player.IsUnattackable() {
		return false
	}

	ex, ey := enemy.Position()
	xEnemyLog := int(ex+16) / 32
	yEnemyLog := int(ey+16) / 32

	xp, yp := player.Position()
	xLog := int(xp+16) / 32
	yLog := int(yp+16) / 32

	return (xEnemyLog == xLog) && (yEnemyLog == yLog)
}

func (g *GamePlay) prepareLevel(isNew bool) {
	if isNew {
		g.currentLevel = 0
		g.score = 0
		g.player.SetLives(3)
	}
	g.currentLevel++

	//cargar datos para el nivel.
	if g.levels == nil {
		g.loadLevels("assets/data/levels")
	}

	//Posicionamos al jugador.
	w, h := ebiten.WindowSize()
	g.player.Reset()
	g.player.SetPosition((w-64)/2+16, (h-32)/2-16)

	//Limpiamos objetos de juego.
	if len(g.objects) > 0 {
		g.objects = make([]*object.CollectableObject, 0)
	}
	if len(g.enemies) > 0 {
		g.enemies = make([]enemies.IEnemy, 0)
	}

	g.stage.PrepareStageForLevel(g.levels[g.currentLevel-1])

	//Estado inicial.
	g.currentState = playing
	g.nextState = g.currentState
	g.alpha = 0.0
	g.goimgblack = nil
	g.isNextLevel = false
	g.playerLeaveLevel = false
	g.oxygenConsumed = 0
}

func (g *GamePlay) loadLevels(filename string) {
	jsonfile := fmt.Sprintf("%s.json", filename)
	data, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		log.Fatal(err)
	}

	var dat levelsjson
	if err := json.Unmarshal(data, &dat); err != nil {
		log.Fatal(err)
	}

	g.levels = make([]gamelevel.GameLevel, 0)
	for _, level := range dat.Levels {
		g.levels = append(g.levels, level)
	}
}

func (g *GamePlay) fadeToBlack(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	sx, sy := ebiten.WindowSize()
	if g.goimgblack == nil {
		g.goimgblack, _ = ebiten.NewImage(1, 1, ebiten.FilterDefault)
		g.goimgblack.Fill(color.Black)
	}
	op.GeoM.Scale(float64(sx), float64(sy))
	op.GeoM.Translate(0, 0)

	//if g.alpha != 1 {
	op.ColorM.Scale(1.0, 1.0, 1.0, g.alpha)
	screen.DrawImage(g.goimgblack, op)
}

type tgamePlayState int

const (
	playing      tgamePlayState = tgamePlayState(0)
	gameover     tgamePlayState = tgamePlayState(1)
	nextlevel    tgamePlayState = tgamePlayState(2)
	victorylevel tgamePlayState = tgamePlayState(3)
	exit         tgamePlayState = tgamePlayState(4)
	end          tgamePlayState = tgamePlayState(5)
)

type levelsjson struct {
	Levels []gamelevel.GameLevel `json:"levels"`
}
