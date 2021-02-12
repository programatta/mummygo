package enemies

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/mummy/states/gameplay/interfaces"
	"github.com/programatta/mummygo/utils"
	"github.com/programatta/mummygo/utils/pathfinding"
)

//Mummy ...
type Mummy struct {
	spriteSheet *utils.SpriteSheet
	posX, posY  float64
	gameplay    interfaces.IGamePlayNotificable
	animations  map[string]*utils.Animation
	currentDir  string
	state       tenemyState
	sc          float64
	toX         float64
	toY         float64
	dirX        int
	dirY        int
	nodesPath   []*pathfinding.Node
}

//NewMummy es el constructor.
func NewMummy(spriteSheet *utils.SpriteSheet, x, y int, gameplay interfaces.IGamePlayNotificable) IEnemy {
	mummy := &Mummy{}

	mummy.spriteSheet = spriteSheet
	mummy.posX = float64(x)
	mummy.posY = float64(y)
	mummy.gameplay = gameplay

	//Preparamos las animaciones.
	mummy.animations = make(map[string]*utils.Animation)

	//Down.
	animaDown := utils.NewAnimation()
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-0.png"))
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-1.png"))
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-2.png"))
	animaDown.AddFrame(spriteSheet.GetFrameByName("mummy-1.png"))
	mummy.animations["down"] = animaDown

	//Left.
	animaLeft := utils.NewAnimation()
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-3.png"))
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-4.png"))
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-5.png"))
	animaLeft.AddFrame(spriteSheet.GetFrameByName("mummy-4.png"))
	mummy.animations["left"] = animaLeft

	//Right.
	animaRight := utils.NewAnimation()
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-6.png"))
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-7.png"))
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-8.png"))
	animaRight.AddFrame(spriteSheet.GetFrameByName("mummy-7.png"))
	mummy.animations["right"] = animaRight

	//Up.
	animaUp := utils.NewAnimation()
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-9.png"))
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-10.png"))
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-11.png"))
	animaUp.AddFrame(spriteSheet.GetFrameByName("mummy-10.png"))
	mummy.animations["up"] = animaUp

	mummy.currentDir = "down"
	mummy.sc = 0
	mummy.state = enemyShowing

	return mummy
}

//Update actualiza la lógica de la momia.
func (m *Mummy) Update(dt float64) {
	m.animations[m.currentDir].UpdateFrame()

	if m.state == enemyShowing || m.state == enemyLeaving {
		m.doPresentation(dt)
	} else if m.state == enemyLookingfor || m.state == enemyNextStep {
		m.doIA()
	} else if m.state == enemyWalking {
		m.doMove(dt)
	}
}

//Draw dibuja el frame actual de la animación.
func (m *Mummy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(m.sc, m.sc)
	op.GeoM.Translate(m.posX, m.posY)
	op.GeoM.Translate(16.0-(32*m.sc)*0.5, 16.0-(32*m.sc)*0.5)

	texture := m.spriteSheet.GetTexture()

	frameData := m.animations[m.currentDir].GetFrame()
	screen.DrawImage(texture.SubImage(image.Rect(frameData.X, frameData.Y, frameData.X+frameData.W, frameData.Y+frameData.H)).(*ebiten.Image), op)
}

//Position devuelve las coordenadas X e Y de la momia.
func (m *Mummy) Position() (float64, float64) {
	return m.posX, m.posY
}

/*===========================================================================*/
/*                               Private Section                             */
/*===========================================================================*/

func (m *Mummy) createMapData(px, py, mx, my float64) *pathfinding.MapData {
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

	//Mummy
	xMummyLog := int(mx+16) / 32
	yMummyLog := int(my+16) / 32

	mapdata[yPlayerLog][xPlayerLog] = pathfinding.STOP
	mapdata[yMummyLog][xMummyLog] = pathfinding.START

	return &mapdata
}

func (m *Mummy) doPresentation(dt float64) {
	if m.state == enemyShowing {
		if m.sc < 1 {
			m.sc += dt
		} else {
			m.sc = 1.0
			m.state = enemyLeaving
			m.toY = m.posY + 32
		}
	} else if m.state == enemyLeaving {
		if m.posY < m.toY {
			m.posY = m.posY + 32*dt
		} else if m.posY >= m.toY {
			m.posY = m.toY
			m.toY = 0
			m.state = enemyLookingfor
		}
	}
}

func (m *Mummy) doIA() {
	if m.state == enemyLookingfor {
		px, py := m.gameplay.OnRequestPlayerPosition()
		mapData := m.createMapData(px, py, m.posX, m.posY)
		graph := pathfinding.NewGraph(mapData)
		m.nodesPath = pathfinding.Astar(graph)
		m.state = enemyNextStep
	} else if m.state == enemyNextStep {
		if len(m.nodesPath) == 0 {
			m.state = enemyLookingfor
		}
		currentNode := m.nodesPath[0]
		toNode := m.nodesPath[1]

		//eliminamos el currentNode de la lista.
		m.nodesPath = m.nodesPath[1:]

		x0, y0 := currentNode.Y, currentNode.X
		x1, y1 := toNode.Y, toNode.X

		if y1-y0 == 0 {
			m.dirY = 0
			m.toY = 0
			if x1-x0 > 0 {
				m.currentDir = "right"
				m.dirX = 32
				m.toX = m.posX + float64(m.dirX)
			} else if x1-x0 < 0 {
				m.currentDir = "left"
				m.dirX = -32
				m.toX = m.posX + float64(m.dirX)
			}
		} else {
			m.dirX = 0
			m.toX = 0
			if y1-y0 < 0 {
				m.currentDir = "up"
				m.dirY = -32
				m.toY = m.posY + float64(m.dirY)
			} else if y1-y0 > 0 {
				m.currentDir = "down"
				m.dirY = 32
				m.toY = m.posY + float64(m.dirY)
			}
		}

		if len(m.nodesPath) == 1 {
			m.state = enemyLookingfor
		} else {
			m.state = enemyWalking
		}
	}
}

func (m *Mummy) doMove(dt float64) {
	if m.toY > 0 {
		if m.dirY > 0 {
			if m.posY < m.toY {
				m.posY = m.posY + float64(m.dirY)*dt
			} else if m.posY >= m.toY {
				m.posY = m.toY
				m.dirY = 0
				m.toY = 0
				m.state = enemyNextStep
			}
		} else if m.dirY < 0 {
			if m.posY > m.toY {
				m.posY = m.posY + float64(m.dirY)*dt
			} else if m.posY <= m.toY {
				m.posY = m.toY
				m.dirY = 0
				m.toY = 0
				m.state = enemyNextStep
			}
		}
	} else if m.toX > 0 {
		if m.dirX > 0 {
			if m.posX < m.toX {
				m.posX = m.posX + float64(m.dirX)*dt
			} else if m.posX >= m.toX {
				m.posX = m.toX
				m.dirX = 0
				m.toX = 0
				m.state = enemyNextStep
			}
		} else if m.dirX < 0 {
			if m.posX > m.toX {
				m.posX = m.posX + float64(m.dirX)*dt
			} else if m.posX <= m.toX {
				m.posX = m.toX
				m.dirX = 0
				m.toX = 0
				m.state = enemyNextStep
			}
		}
	}
}
