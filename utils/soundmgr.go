package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

//SoundMgr ...
type SoundMgr struct {
	audioContext *audio.Context
	audioPlayers map[string]*audio.Player
}

//NewSoundMgr es un constructor.
func NewSoundMgr(sampleRate int) *SoundMgr {
	soundmgr := &SoundMgr{}
	soundmgr.audioContext, _ = audio.NewContext(sampleRate)
	soundmgr.audioPlayers = make(map[string]*audio.Player)

	return soundmgr
}

//Load carga un sonido
func (sm *SoundMgr) Load(fullPathFileName string) error {
	sound, err := sm.loadAndPrepare(fullPathFileName)
	if err != nil {
		return err
	}

	soundName := filepath.Base(fullPathFileName)
	sm.audioPlayers[soundName] = sound
	return nil
}

//LoadSet carga un conjunto de sonidos desde un path
func (sm *SoundMgr) LoadSet(fullPathFileName string, soundNames []string) error {
	var err error = nil
	for _, soundname := range soundNames {
		fp := fmt.Sprintf("%s/%s", fullPathFileName, soundname)
		err = sm.Load(fp)
		if err != nil {
			break
		}
	}
	return err
}

//Sound devuelve un player a partir del nombre de la fuente de sonido.
func (sm *SoundMgr) Sound(soundName string) *audio.Player {
	return sm.audioPlayers[soundName]
}

func (sm *SoundMgr) loadAndPrepare(fullPathFileName string) (*audio.Player, error) {
	f, err := os.Open(fullPathFileName)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)
	const sizeInBytes = 10240 //10Kb

	soundData := make([]byte, 0)
	var data []byte
	data = make([]byte, sizeInBytes)
	for {
		n, err := reader.Read(data)
		if err != nil {
			return nil, err
		}
		soundData = append(soundData, data...)
		if n < sizeInBytes {
			break
		}
	}

	soundDec, err := wav.Decode(sm.audioContext, audio.BytesReadSeekCloser(soundData))
	if err != nil {
		return nil, err
	}
	soundPlayer, err := audio.NewPlayer(sm.audioContext, soundDec)
	return soundPlayer, err
}
