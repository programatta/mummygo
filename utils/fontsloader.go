package utils

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//FontsLoader ...
type FontsLoader struct {
	fonts map[string]*truetype.Font
}

//NewFontsLoader es el constructor.
func NewFontsLoader() *FontsLoader {
	fl := &FontsLoader{}

	//Creamos el mapa de fuentes.
	fl.fonts = make(map[string]*truetype.Font)

	return fl
}

//Load carga la fuente indicada en filename
func (fl *FontsLoader) Load(filename string) error {
	f, err := os.Open(filename) //"assets/fonts/ka1.ttf")
	defer f.Close()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)

	const sizeInBytes = 10240 //10Kb

	fontData := make([]byte, 0)
	var data []byte
	data = make([]byte, sizeInBytes)
	for {
		n, err := reader.Read(data)
		if err != nil {
			return err
		}
		fontData = append(fontData, data...)
		if n < sizeInBytes {
			break
		}
	}
	ttfont, err := truetype.Parse(fontData)
	if err != nil {
		return err
	}

	fontName := filepath.Base(filename)

	//TODO: que se pueda cargar cuando lo pedimos.
	fl.fonts[fontName] = ttfont
	return nil
}

//GetFont devuelve la fuente cargada.
func (fl *FontsLoader) GetFont(fontname string, dpi float64, fontSize float64) font.Face {
	ttfont := fl.fonts[fontname]
	font := truetype.NewFace(ttfont, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	return font
}
