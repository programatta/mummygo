package object

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/programatta/mummygo/utils"
)

//CollectableObject ...
type CollectableObject struct {
	spriteSheet *utils.SpriteSheet
	soundmgr    *utils.SoundMgr
	typeObject  int
	posX, posY  float64
	sc          float64
	alpha       float64
	state       tobjectState //0 = escalando, 1= moviendo, 2 = final
	toY         float64
}

//NewCollectableObject is a constructor
func NewCollectableObject(spriteSheet *utils.SpriteSheet, soundmgr *utils.SoundMgr, t, x, y int) *CollectableObject {
	co := &CollectableObject{}
	co.spriteSheet = spriteSheet
	co.soundmgr = soundmgr
	co.typeObject = t
	co.posX = float64(x)
	co.posY = float64(y)
	co.sc = 0
	co.alpha = 1
	co.state = objShowing
	co.toY = 0
	return co
}

//Update ...
func (co *CollectableObject) Update(dt float64) {
	//Estados
	//ESTADO1 - escala 0.0x -> 1
	//ESTADO2 - pos_a -> pos_b (cae)
	if co.state == objShowing {
		if co.sc < 1 {
			co.sc += dt
		} else {
			co.sc = 1.0
			co.state = objLeaving
			co.toY = co.posY + 32
		}
	} else if co.state == objLeaving {
		if co.posY < co.toY {
			co.posY = co.posY + 32*dt //0.016
		} else if co.posY >= co.toY {
			co.posY = co.toY
			co.toY = 0
			co.state = objBlinkLess

			if co.typeObject == 2 {
				//Pocion
				potionshow := co.soundmgr.Sound("potionleave.wav")
				if !potionshow.IsPlaying() {
					potionshow.Rewind()
					potionshow.Play()
				}
			} else {
				//Llave y papiro.
				itemshow := co.soundmgr.Sound("itemleave.wav")
				if !itemshow.IsPlaying() {
					itemshow.Rewind()
					itemshow.Play()
				}
			}
		}
	} else if co.state == objBlinkLess {
		if co.alpha > 0.4 {
			co.alpha -= dt / 2
		} else {
			co.alpha = 0.4
			co.state = objBlinkMore
		}
	} else if co.state == objBlinkMore {
		if co.alpha < 1 {
			co.alpha += dt / 2
		} else {
			co.alpha = 1
			co.state = objBlinkLess
		}
	}
}

//Draw ...
func (co *CollectableObject) Draw(screen *ebiten.Image) {
	itemName := ""
	switch co.typeObject {
	case 2:
		itemName = "object-0.png"
		break
	case 3:
		itemName = "object-3.png"
		break
	case 4:
		itemName = "object-14.png"
		break
	}

	frameData := co.spriteSheet.GetFrameByName(itemName)
	texture := co.spriteSheet.GetTexture()

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(co.sc, co.sc)
	//	op.GeoM.Translate(co.posX+16.0-(32*co.sc)*0.5, co.posY+16.0-(32*co.sc)*0.5)
	op.GeoM.Translate(co.posX, co.posY)
	op.GeoM.Translate(16.0-(32*co.sc)*0.5, 16.0-(32*co.sc)*0.5)
	//op.Filter = ebiten.FilterLinear | ebiten.FilterDefault
	op.ColorM.Scale(1.0, 1.0, 1.0, co.alpha)
	screen.DrawImage(texture.SubImage(
		image.Rect(frameData.X, frameData.Y, frameData.X+frameData.W, frameData.Y+frameData.H)).(*ebiten.Image),
		op)
}

//TypeObject devuelve el tipo de objeto.
func (co *CollectableObject) TypeObject() int {
	return co.typeObject
}

//Position devuelve las coordenadas X e Y del objeto.
func (co *CollectableObject) Position() (float64, float64) {
	return co.posX, co.posY
}

//PickedUp reproduce un sonido cuando el gameplay le informa de que ha sido cogido.
func (co *CollectableObject) PickedUp() {
	if co.typeObject == 2 {
		//Pocion.
		potionPlayer := co.soundmgr.Sound("potiondrink.wav")
		potionPlayer.Rewind()
		potionPlayer.Play()
	} else {
		//Llave y papiro.
		itemPlayer := co.soundmgr.Sound("pickupitem.wav")
		itemPlayer.Rewind()
		itemPlayer.Play()
	}
}

type tobjectState int

const (
	objShowing   tobjectState = tobjectState(0)
	objLeaving   tobjectState = tobjectState(1)
	objBlinkLess tobjectState = tobjectState(2)
	objBlinkMore tobjectState = tobjectState(3)
)
