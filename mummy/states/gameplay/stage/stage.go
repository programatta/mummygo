package stage

import (
	"image"
	"math/rand"
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/mummy/states/gameplay/gamelevel"
	gp "github.com/programatta/mummygo/mummy/states/gameplay/interfaces"
	"github.com/programatta/mummygo/utils"
)

//Stage ...
type Stage struct {
	spriteSheet *utils.SpriteSheet
	soundmgr    *utils.SoundMgr
	gameplay    gp.IGamePlayNotificable
	logicMap    [][]int
	tombs       []*Tomb
}

//NewStage is a constructor.
func NewStage(spriteSheet *utils.SpriteSheet, soundmgr *utils.SoundMgr, gameplay gp.IGamePlayNotificable) *Stage {
	stage := &Stage{}

	stage.spriteSheet = spriteSheet
	stage.soundmgr = soundmgr
	stage.gameplay = gameplay

	return stage
}

//Update ...
func (s *Stage) Update(dt float64) {
	if len(s.tombs) > 0 {
		tombsTmp := make([]*Tomb, len(s.tombs))
		copied := 0

		for _, tomb := range s.tombs {
			tomb.Update(s.logicMap)
			if tomb.canOpen {
				s.createObjectType(tomb.contentType, tomb.x+1, tomb.y+1)
				s.logicMap[tomb.y+1][tomb.x+1] = 3 //Tomb open
				tomb = nil
				openDoorPlayer := s.soundmgr.Sound("opendoor.wav")
				openDoorPlayer.Rewind()
				openDoorPlayer.Play()
			} else {
				tombsTmp[copied] = tomb
				copied++
			}
		}
		s.tombs = tombsTmp[:copied]
	}
}

//Draw ...
func (s *Stage) Draw(screen *ebiten.Image) {
	texture := s.spriteSheet.GetTexture()
	for y, row := range s.logicMap {
		for x, value := range row {
			var pos utils.Rect
			switch value {
			case 0: //Ground
				pos = s.spriteSheet.GetFrameByName("desert-5.png")
				break
			case 1: //Wall
				pos = s.spriteSheet.GetFrameByName("tiles-36.png")
				break
			case 2: //Tomb Door
				pos = s.spriteSheet.GetFrameByName("tiles-57.png")
				break
			case 3: //Tomb Door Open
				pos = s.spriteSheet.GetFrameByName("tiles-58.png")
				break
			case 4: //Main Door
				pos = s.spriteSheet.GetFrameByName("tiles-54.png")
				break
			case 5: //Main Door Open
				pos = s.spriteSheet.GetFrameByName("tiles-55.png")
				break
			case 6: //Steps (Up)
				pos = s.spriteSheet.GetFrameByName("desert-33.png")
				break
			case 7: //Steps (Down)
				pos = s.spriteSheet.GetFrameByName("desert-52.png")
				break
			case 8: //Steps (Left)
				pos = s.spriteSheet.GetFrameByName("desert-51.png")
				break
			case 9: //Steps (Right)
				pos = s.spriteSheet.GetFrameByName("desert-32.png")
				break
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x)*32, float64(y)*32)

			screen.DrawImage(texture.SubImage(
				image.Rect(pos.X, pos.Y, pos.X+pos.W, pos.Y+pos.H)).(*ebiten.Image),
				op)
		}
	}
}

//PrepareStageForLevel ...
func (s *Stage) PrepareStageForLevel(level gamelevel.GameLevel) {

	contentType := make(map[int]int)
	tombIds := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	s.logicMap = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1, 2, 1, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	s.tombs = make([]*Tomb, 0)

	seed := rand.NewSource(time.Now().Local().Unix())
	rnd := rand.New(seed)

	//Posicionamos llave y papiro.
	keyPos := rnd.Intn(8)        //[0,8) => 0,1,2,3,4,5,6,7
	papyrePos := rnd.Intn(8) + 8 //[8,16)=>8,9,10,11,12,13,14,15

	//Borramos de abajo hacia arriba para no descolocar posiciones.
	//.: primero el papiro.
	s.removePosition(&tombIds, papyrePos)
	contentType[papyrePos] = 4

	//.: luego la llave.
	s.removePosition(&tombIds, keyPos)
	contentType[keyPos] = 3

	//Resto de elementos de nivel.
	for level.Mummies > 0 {
		pos := rnd.Intn(len(tombIds))
		valpos := tombIds[pos]
		s.removePosition(&tombIds, pos)
		contentType[valpos] = 1
		level.Mummies--
	}

	for level.Potins > 0 {
		pos := rnd.Intn(len(tombIds))
		valpos := tombIds[pos]
		s.removePosition(&tombIds, pos)
		contentType[valpos] = 2
		level.Potins--
	}

	for level.Spells > 0 {
		pos := rnd.Intn(len(tombIds))
		valpos := tombIds[pos]
		s.removePosition(&tombIds, pos)
		contentType[valpos] = 5
		level.Spells--
	}

	for _, valpos := range tombIds {
		contentType[valpos] = 0
	}

	//Ordenamos las posiciones.
	keys := make([]int, 0, len(contentType))
	for k := range contentType {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i, k := range keys {
		//Esquina superior izquierda de cada tumba.
		x1 := 2 + 4*(i/4)
		y1 := 2 + 3*(i%4)

		tomb := NewTomb(x1, y1, contentType[k])
		s.tombs = append(s.tombs, tomb)
	}
}

//OpenMainDoor ...
func (s *Stage) OpenMainDoor() {
	if s.logicMap[0][9] != 5 {
		s.logicMap[0][9] = 5

		mainOpenDoorPlayer := s.soundmgr.Sound("mainopendoor.wav")
		mainOpenDoorPlayer.Rewind()
		mainOpenDoorPlayer.Play()
	}
}

//GetTypeAt ...
func (s *Stage) GetTypeAt(x, y int) int {
	t := s.logicMap[y][x]
	return t //s.logicMap[x][y]
}

//SetTypeAt ...
func (s *Stage) SetTypeAt(x, y, t int) {
	if s.logicMap[y][x] == 0 {
		s.logicMap[y][x] = t
		s.gameplay.OnScoreByStep()
	}
}

//DebugOpenTombs ...
func (s *Stage) DebugOpenTombs() {
	for _, row := range s.logicMap {
		for i := range row {
			if row[i] == 0 {
				row[i] = 6 //Marcados como pisados.
			}
		}
	}
	s.logicMap[0][9] = 5 //Puerta principal abierta.
}

func (s *Stage) createObjectType(t, x, y int) {
	xfis := x * 32
	yfis := y * 32

	s.gameplay.OnCreateObject(t, xfis, yfis)
}

func (s *Stage) removePosition(array *[]int, pos int) {
	copy((*array)[pos:], (*array)[pos+1:])
	(*array)[len((*array))-1] = -1
	*array = ((*array)[:len((*array))-1])
}
