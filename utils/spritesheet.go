package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//SpriteSheet structure
type SpriteSheet struct {
	texture *ebiten.Image
	frames  map[string]Rect
}

//NewSpriteSheet is a constructor.
func NewSpriteSheet() *SpriteSheet {
	return &SpriteSheet{}
}

//Load function loads texture and process json file.
func (s *SpriteSheet) Load(filename string) error {
	var err error

	//Load png file.
	pngfile := fmt.Sprintf("%s.png", filename)
	s.texture, _, err = ebitenutil.NewImageFromFile(pngfile, ebiten.FilterDefault)
	if err != nil {
		return err
	}

	//Load json file.
	jsonfile := fmt.Sprintf("%s.json", filename)
	data, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		return err
	}

	type sprjson struct {
		Frames []tframe `json:"frames"`
		Meta   tmeta    `json:"meta"`
	}

	var dat sprjson
	if err := json.Unmarshal(data, &dat); err != nil {
		return err
	}
	s.frames = make(map[string]Rect)
	for _, frame := range dat.Frames {
		s.frames[frame.Filename] = frame.Frame
	}

	return nil
}

//GetTexture returns a spritesheet texture.
func (s *SpriteSheet) GetTexture() *ebiten.Image {
	return s.texture
}

//GetFrameByName returns position in texture
func (s *SpriteSheet) GetFrameByName(name string) Rect {
	return s.frames[name]
}

type tsize struct {
	W int `json:"w"`
	H int `json:"h"`
}
type Rect struct {
	tsize
	X int `json:"x"`
	Y int `json:"y"`
}

type tframe struct {
	Filename         string `json:"filename"`         //	"filename": "desert-3.png",
	Frame            Rect   `json:"frame"`            //"frame": {"x":0,"y":0,"w":32,"h":32},
	Rotated          bool   `json:"rotated"`          //"rotated": false,
	Trimmed          bool   `json:"trimmed"`          //"trimmed": false,
	SpriteSourceSize Rect   `json:"spriteSourceSize"` // "spriteSourceSize": {"x":0,"y":0,"w":32,"h":32},
	SourceSize       tsize  `json:"sourceSize"`       //"sourceSize": {"w":32,"h":32}
}

type tmeta struct {
	App         string `json:"app"`         //"app": "https://www.codeandweb.com/texturepacker",
	Version     string `json:"version"`     //"version": "1.0",
	Image       string `json:"image"`       //"image": "sprites.png",
	Format      string `json:"format"`      //"format": "RGBA8888",
	Size        tsize  `json:"size"`        //"size": {"w":512,"h":96},
	Scale       string `json:"scale"`       //"scale": "1",
	SmartUpdate string `json:"smartupdate"` //"smartupdate": "$TexturePacker:SmartUpdate:3f43aedae3b01ae0320879f7dcc8327c:be04319743ecef53acabcef518ac51ea:1eabdf11f75e3a4fe3147baf7b5be24b$"
}
