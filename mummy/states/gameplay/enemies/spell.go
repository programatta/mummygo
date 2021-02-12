package enemies

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/mummy/states/gameplay/interfaces"
	"github.com/programatta/mummygo/utils"
	"github.com/programatta/mummygo/utils/pathfinding"
)

//Spell contiene la funcionalidad del hechizo.
type Spell struct {
	spriteSheet *utils.SpriteSheet
	posX, posY  float64
	gameplay    interfaces.IGamePlayNotificable
	state       tenemyState
	sc          float64
	rot         float64
	toX         float64
	toY         float64
	dirX        int
	dirY        int
	nodesPath   []*pathfinding.Node
	scaleLess   bool
}

//NewSpell es el constructor.
func NewSpell(spriteSheet *utils.SpriteSheet, x, y int, gameplay interfaces.IGamePlayNotificable) IEnemy {
	spell := &Spell{}

	spell.spriteSheet = spriteSheet
	spell.posX = float64(x)
	spell.posY = float64(y)
	spell.gameplay = gameplay

	spell.sc = 0
	spell.rot = 0
	spell.state = enemyShowing

	return spell
}

//Update actualiza la lógia del hechizo.
func (s *Spell) Update(dt float64) {
	if s.state == enemyShowing || s.state == enemyLeaving {
		s.doPresentation(dt)
	} else if s.state == enemyLookingfor || s.state == enemyNextStep {
		s.doIA()
	} else if s.state == enemyWalking {
		s.doMove(dt)
	}
}

//Draw dibuja el hechizo.
func (s *Spell) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if s.rot > 0 {
		op.GeoM.Translate(-16.0, -16.0)
		op.GeoM.Rotate(s.rot)
		op.GeoM.Translate(16.0, 16.0)
	}
	op.GeoM.Scale(s.sc, s.sc)
	op.GeoM.Translate(s.posX, s.posY)
	op.GeoM.Translate(16.0-(32*s.sc)*0.5, 16.0-(32*s.sc)*0.5)

	texture := s.spriteSheet.GetTexture()
	frameData := s.spriteSheet.GetFrameByName("object-58.png")

	screen.DrawImage(texture.SubImage(image.Rect(frameData.X, frameData.Y, frameData.X+frameData.W, frameData.Y+frameData.H)).(*ebiten.Image), op)
}

//Position devuelve las coordenadas X e Y del hechizo.
func (s *Spell) Position() (float64, float64) {
	return s.posX, s.posY
}

/*===========================================================================*/
/*                               Private Section                             */
/*===========================================================================*/

func (s *Spell) createMapData(px, py, sx, sy float64) *pathfinding.MapData {
	mapdata := pathfinding.MapData{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	//Player
	xPlayerLog := int(px+16) / 32
	yPlayerLog := int(py+16) / 32

	//Spell
	xSpellLog := int(sx+16) / 32
	ySpellLog := int(sy+16) / 32

	mapdata[yPlayerLog][xPlayerLog] = pathfinding.STOP
	mapdata[ySpellLog][xSpellLog] = pathfinding.START

	return &mapdata
}

func (s *Spell) doPresentation(dt float64) {
	if s.state == enemyShowing {
		if s.sc < 1 {
			s.sc += dt
		} else {
			s.sc = 1.0
			s.state = enemyLeaving
			s.toY = s.posY + 32
		}
	} else if s.state == enemyLeaving {
		if s.posY < s.toY {
			s.posY = s.posY + 32*dt
		} else if s.posY >= s.toY {
			s.posY = s.toY
			s.toY = 0
			s.state = enemyLookingfor
		}
	}
}

func (s *Spell) doIA() {
	if s.state == enemyLookingfor {
		px, py := s.gameplay.OnRequestPlayerPosition()
		mapData := s.createMapData(px, py, s.posX, s.posY)
		graph := pathfinding.NewGraph(mapData)
		s.nodesPath = pathfinding.Astar(graph)
		s.state = enemyNextStep
	} else if s.state == enemyNextStep {
		if len(s.nodesPath) == 1 {
			s.state = enemyLookingfor
			return
		}
		currentNode := s.nodesPath[0]
		toNode := s.nodesPath[1]

		//eliminamos el currentNode de la lista.
		s.nodesPath = s.nodesPath[1:]

		x0, y0 := currentNode.Y, currentNode.X
		x1, y1 := toNode.Y, toNode.X

		if y1-y0 == 0 {
			s.dirY = 0
			s.toY = 0
			if x1-x0 > 0 {
				s.dirX = 32
				s.toX = s.posX + float64(s.dirX)
			} else if x1-x0 < 0 {
				s.dirX = -32
				s.toX = s.posX + float64(s.dirX)
			}
		} else {
			s.dirX = 0
			s.toX = 0
			if y1-y0 < 0 {
				s.dirY = -32
				s.toY = s.posY + float64(s.dirY)
			} else if y1-y0 > 0 {
				s.dirY = 32
				s.toY = s.posY + float64(s.dirY)
			}
		}

		s.state = enemyWalking
	}
}

func (s *Spell) doMove(dt float64) {
	const vel float64 = 3
	s.rot += 15 * dt / 2

	if s.scaleLess {
		if s.sc > 0.4 {
			s.sc -= dt
		} else {
			s.sc = 0.4
			s.scaleLess = false
		}
	} else {
		if s.sc < 1 {
			s.sc += dt
		} else {
			s.sc = 1
			s.scaleLess = true
		}
	}

	if s.toY > 0 {
		if s.dirY > 0 {
			if s.posY < s.toY {
				s.posY = s.posY + float64(s.dirY)*dt*vel

				if s.posY >= s.toY {
					s.posY = s.toY
					s.dirY = 0
					s.toY = 0
					s.state = enemyNextStep
				}
			}
		} else if s.dirY < 0 {
			if s.posY > s.toY {
				s.posY = s.posY + float64(s.dirY)*dt*vel

				if s.posY <= s.toY {
					s.posY = s.toY
					s.dirY = 0
					s.toY = 0
					s.state = enemyNextStep
				}
			}
		}
	} else if s.toX > 0 {
		if s.dirX > 0 {
			if s.posX < s.toX {
				s.posX = s.posX + float64(s.dirX)*dt*vel

				if s.posX >= s.toX {
					s.posX = s.toX
					s.dirX = 0
					s.toX = 0
					s.state = enemyNextStep
				}
			}
		} else if s.dirX < 0 {
			if s.posX > s.toX {
				s.posX = s.posX + float64(s.dirX)*dt*vel

				if s.posX <= s.toX {
					s.posX = s.toX
					s.dirX = 0
					s.toX = 0
					s.state = enemyNextStep
				}
			}
		}
	}
}
