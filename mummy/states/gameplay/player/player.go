package player

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/mummy/states/gameplay/interfaces"
	"github.com/programatta/mummygo/mummy/states/gameplay/object"
	"github.com/programatta/mummygo/utils"
)

//Player ...
type Player struct {
	spriteSheet      *utils.SpriteSheet
	animations       map[string]*utils.Animation
	currentDir       string
	posX             float64
	posY             float64
	dirY             int
	dirX             int
	toX              int
	toY              int
	currentFrameData utils.Rect
	stage            interfaces.IStageNotificable
	potions          int
	lives            int
	hasKey           bool
	hasPapyre        bool
	isLeavingLevel   bool
	state            tplayerState
	sc               float64
	alpha            float64
	gameplay         interfaces.IGamePlayNotificable
	blinkingTime     float64
	isBlinking       bool
}

//PlayerLeft ...
const PlayerLeft int = 0

//PlayerRight ...
const PlayerRight int = 1

//PlayerUp ...
const PlayerUp int = 2

//PlayerDown ...
const PlayerDown int = 3

//NewPlayer is a constructor.
func NewPlayer(spriteSheet *utils.SpriteSheet, stage interfaces.IStageNotificable) *Player {
	player := &Player{}

	player.spriteSheet = spriteSheet
	player.stage = stage

	//Animations.
	player.animations = make(map[string]*utils.Animation)

	//Down.
	animDown := utils.NewAnimation()
	animDown.AddFrame(spriteSheet.GetFrameByName("player-0.png"))
	animDown.AddFrame(spriteSheet.GetFrameByName("player-1.png"))
	animDown.AddFrame(spriteSheet.GetFrameByName("player-2.png"))
	animDown.AddFrame(spriteSheet.GetFrameByName("player-1.png"))
	player.animations["down"] = animDown

	//Left.
	animLeft := utils.NewAnimation()
	animLeft.AddFrame(spriteSheet.GetFrameByName("player-3.png"))
	animLeft.AddFrame(spriteSheet.GetFrameByName("player-4.png"))
	animLeft.AddFrame(spriteSheet.GetFrameByName("player-5.png"))
	animLeft.AddFrame(spriteSheet.GetFrameByName("player-4.png"))
	player.animations["left"] = animLeft

	//Right.
	animRight := utils.NewAnimation()
	animRight.AddFrame(spriteSheet.GetFrameByName("player-6.png"))
	animRight.AddFrame(spriteSheet.GetFrameByName("player-7.png"))
	animRight.AddFrame(spriteSheet.GetFrameByName("player-8.png"))
	animRight.AddFrame(spriteSheet.GetFrameByName("player-7.png"))
	player.animations["right"] = animRight

	//Up.
	animUp := utils.NewAnimation()
	animUp.AddFrame(spriteSheet.GetFrameByName("player-9.png"))
	animUp.AddFrame(spriteSheet.GetFrameByName("player-10.png"))
	animUp.AddFrame(spriteSheet.GetFrameByName("player-11.png"))
	animUp.AddFrame(spriteSheet.GetFrameByName("player-10.png"))
	player.animations["up"] = animUp

	player.currentDir = "down"

	return player
}

//Update updates a player logic
func (p *Player) Update(dt float64) {
	if p.toX > 0 || p.toY > 0 {
		p.animations[p.currentDir].UpdateFrame()
		p.currentFrameData = p.animations[p.currentDir].GetFrame()
	} else {
		p.currentFrameData = p.animations[p.currentDir].GetFrameIndex(1)
	}

	if p.isLeavingLevel {
		if p.state == playerLeaving {
			if p.posY > float64(p.toY) {
				p.posY = p.posY + float64(p.dirY)*dt //0.016
			} else if p.posY <= float64(p.toY) {
				p.posY = float64(p.toY)
				p.dirY = 0
				p.toY = 0
				p.state = playerHiding
			}
		} else if p.state == playerHiding {
			if p.sc > 0 {
				p.sc -= dt
				p.alpha -= dt
			} else {
				p.sc = 0
				p.alpha = 0
				p.state = playerLeftLevel
			}
		} else if p.state == playerLeftLevel {
			p.gameplay.OnPrepreNewLevel()
		}
	} else {
		//Modo blinking (al ser alcanzado por una momia)
		if p.blinkingTime > 0 {
			p.blinkingTime -= dt

			if p.state == playerBlinkLess {
				if p.alpha > 0.4 {
					p.alpha -= dt
				} else {
					p.alpha = 0.4
					p.state = playerBlinkMore
				}
			} else if p.state == playerBlinkMore {
				if p.alpha < 1 {
					p.alpha += dt
				} else {
					p.alpha = 1
					p.state = playerBlinkLess
				}
			}
		} else {
			p.alpha = 1
			p.isBlinking = false
			p.blinkingTime = 0
		}

		//TODO: Modo hechizado (al ser alcanzado por el hechizo).

		if p.toY > 0 {
			//Check collision.
			if p.checkCollision(int(p.posX), p.toY) {
				p.toY = 0
				p.dirY = 0
				return
			}

			if p.dirY > 0 {
				if p.posY < float64(p.toY) {
					p.posY = p.posY + float64(p.dirY)*dt //0.016
				} else if p.posY >= float64(p.toY) {
					p.posY = float64(p.toY)
					p.dirY = 0
					p.toY = 0
					p.updateMap(p.posX, p.posY, p.currentDir)
				}
			} else if p.dirY < 0 {
				if p.posY > float64(p.toY) {
					p.posY = p.posY + float64(p.dirY)*dt //0.016
				} else if p.posY <= float64(p.toY) {
					p.posY = float64(p.toY)
					p.dirY = 0
					p.toY = 0
					p.updateMap(p.posX, p.posY, p.currentDir)
				}
			}
		} else if p.toX > 0 {
			//Check collision.
			if p.checkCollision(p.toX, int(p.posY)) {
				p.toX = 0
				p.dirX = 0
				return
			}

			if p.dirX > 0 {
				if p.posX < float64(p.toX) {
					p.posX = p.posX + float64(p.dirX)*dt //0.016
				} else if p.posX >= float64(p.toX) {
					p.posX = float64(p.toX)
					p.dirX = 0
					p.toX = 0
					p.updateMap(p.posX, p.posY, p.currentDir)
				}
			} else if p.dirX < 0 {
				if p.posX > float64(p.toX) {
					p.posX = p.posX + float64(p.dirX)*dt //0.016
				} else if p.posX <= float64(p.toX) {
					p.posX = float64(p.toX)
					p.dirX = 0
					p.toX = 0
					p.updateMap(p.posX, p.posY, p.currentDir)
				}
			}
		}
	}
}

//Draw renders a current frame.
func (p *Player) Draw(screen *ebiten.Image) {
	//Nos ubicamos en el centro del sprite.
	op := &ebiten.DrawImageOptions{}
	if p.isLeavingLevel {
		op.GeoM.Scale(p.sc, p.sc)
	}
	op.GeoM.Translate(p.posX, p.posY)
	if p.isLeavingLevel {
		op.GeoM.Translate(16.0-(32*p.sc)*0.5, 16.0-(32*p.sc)*0.5)
		op.ColorM.Scale(1.0, 1.0, 1.0, p.alpha)
	}

	if p.isBlinking {
		op.ColorM.Scale(1.0, 1.0, 1.0, p.alpha)
	}

	texture := p.spriteSheet.GetTexture()

	screen.DrawImage(texture.SubImage(
		image.Rect(p.currentFrameData.X, p.currentFrameData.Y, p.currentFrameData.X+p.currentFrameData.W, p.currentFrameData.Y+p.currentFrameData.H)).(*ebiten.Image),
		op)
}

//Move preare directon for player
func (p *Player) Move(dir int) {
	switch dir {
	case PlayerUp:
		if p.toX == 0 && p.toY == 0 {
			p.dirY = -32
			p.toY = int(p.posY) + p.dirY
			p.currentDir = "up"
		}
		break
	case PlayerDown:
		if p.toX == 0 && p.toY == 0 {
			p.dirY = 32
			p.toY = int(p.posY) + p.dirY
			p.currentDir = "down"
		}
		break
	case PlayerLeft:
		if p.toX == 0 && p.toY == 0 {
			p.dirX = -32
			p.toX = int(p.posX) + p.dirX
			p.currentDir = "left"
		}
		break
	case PlayerRight:
		if p.toX == 0 && p.toY == 0 {
			p.dirX = 32
			p.toX = int(p.posX) + p.dirX
			p.currentDir = "right"
		}
		break
	}
}

//AddObject ...
func (p *Player) AddObject(object *object.CollectableObject) {
	if object.TypeObject() == 3 {
		p.hasKey = true
	} else if object.TypeObject() == 4 {
		p.hasPapyre = true
	} else {
		p.potions++
	}
}

//LeaveLevel ...
func (p *Player) LeaveLevel(gameplay interfaces.IGamePlayNotificable) {
	p.gameplay = gameplay
	p.isLeavingLevel = true
	p.state = playerLeaving
	p.sc = 1
	p.alpha = 1
	p.dirY = -32
	p.toY = 7 //Forzamos a esta posición que da un efecto muy bueno.
	p.currentDir = "up"
}

//HasKeyAndPapyre devuelve true si el player ha cogido la llave y el papiro.
func (p *Player) HasKeyAndPapyre() bool {
	return p.hasKey && p.hasPapyre
}

//Position devuelve los valores de X e Y del jugador.
func (p *Player) Position() (float64, float64) {
	return p.posX, p.posY
}

//SetPosition establece la posición del jugador en pantalla.
func (p *Player) SetPosition(x, y int) {
	p.posX = float64(x)
	p.posY = float64(y)
}

//Lives devuelve el número de vidas del jugador.
func (p *Player) Lives() int {
	return p.lives
}

//SetLives establece las vidas del jugador.
func (p *Player) SetLives(lives int) {
	p.lives = lives
}

//LostLive decrementa una vida cuando colisionamos con una momia.
func (p *Player) LostLive() {
	p.lives--

	//Reseteamos las direcciones y los pontos a donde ir.
	p.currentDir = "down"
	p.dirX = 0
	p.dirY = 0
	p.toX = 0
	p.toY = 0

	//blinking
	p.isBlinking = true
	p.state = playerBlinkLess
	p.blinkingTime = 5 //segundos.
}

//Potions devuelve el ńumero de pociones que tiene el jugador.
func (p *Player) Potions() int {
	return p.potions
}

//ConsumePotion decrementa el número de pociones del jugador.
//Se llama cuando colisionamos con una momia.
func (p *Player) ConsumePotion() {
	p.potions--
}

//CurrentDir devuelve la cadena de dirección en la que se desplaza el jugador.
func (p *Player) CurrentDir() string {
	return p.currentDir
}

//IsBlinking devuelve true si el player está saliendo de una muerte anterior.
//En este estado no afectan colisiones con las momias.
func (p *Player) IsBlinking() bool {
	return p.isBlinking
}

func (p *Player) checkCollision(x, y int) bool {

	xlog := x / 32
	ylog := y / 32

	hasCol := false
	switch p.stage.GetTypeAt(xlog, ylog) {
	case 1: //Wall
		fallthrough
	case 2: //Tomb Door
		fallthrough
	case 3: //Tomb Door Open
		fallthrough
	case 4: //Main Door
		hasCol = true
		break
	}

	return hasCol
}

func (p *Player) updateMap(x, y float64, dir string) {
	xlog := int(x) / 32
	ylog := int(y) / 32

	if p.stage.GetTypeAt(xlog, ylog) == 0 {
		switch dir {
		case "up":
			p.stage.SetTypeAt(xlog, ylog, 6)
			break
		case "down":
			p.stage.SetTypeAt(xlog, ylog, 7)
			break
		case "left":
			p.stage.SetTypeAt(xlog, ylog, 8)
			break
		case "right":
			p.stage.SetTypeAt(xlog, ylog, 9)
			break
		}
	}
}

type tplayerState int

const (
	playerHiding    tplayerState = tplayerState(0)
	playerLeaving   tplayerState = tplayerState(1)
	playerLeftLevel tplayerState = tplayerState(2)
	playerBlinkLess tplayerState = tplayerState(3)
	playerBlinkMore tplayerState = tplayerState(4)
)
