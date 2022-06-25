package magnet

import (
	"bytes"
	"io/fs"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	resources "github.com/raa0121/magnet/magnet/internal"
)

var (
	audioContext *audio.Context
	bgmPlayer *audio.Player
	jumpPlayer *audio.Player
	failPlayer *audio.Player
	slidePlayer *audio.Player
	damagePlayer *audio.Player
)

func init() {
	audioContext = audio.NewContext(44100)
	bgmPlayerInit()
	jumpPLayerInit()
	failPLayerInit()
	slidePLayerInit()
	damagePLayerInit()
}

func bgmPlayerInit() {
	b, err := fs.ReadFile(resources.Bgms, "Magnet.mp3")
	if err != nil {
		log.Fatal(err)
	}
	voice, err := mp3.Decode(audioContext, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	bgmPlayer, err = audio.NewPlayer(audioContext, voice)
	if err != nil {
		log.Fatal(err)
	}
}

func jumpPLayerInit() {
	b, err := fs.ReadFile(resources.Ses, "Jump.mp3")
	if err != nil {
		log.Fatal(err)
	}
	voice, err := mp3.Decode(audioContext, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	jumpPlayer, err = audio.NewPlayer(audioContext, voice)
	if err != nil {
		log.Fatal(err)
	}
}

func failPLayerInit() {
	b, err := fs.ReadFile(resources.Ses, "Fail.mp3")
	if err != nil {
		log.Fatal(err)
	}
	voice, err := mp3.Decode(audioContext, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	failPlayer, err = audio.NewPlayer(audioContext, voice)
	if err != nil {
		log.Fatal(err)
	}
}

func slidePLayerInit() {
	b, err := fs.ReadFile(resources.Ses, "Slide.mp3")
	if err != nil {
		log.Fatal(err)
	}
	voice, err := mp3.Decode(audioContext, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	slidePlayer, err = audio.NewPlayer(audioContext, voice)
	if err != nil {
		log.Fatal(err)
	}
}

func damagePLayerInit() {
	b, err := fs.ReadFile(resources.Ses, "Damage.mp3")
	if err != nil {
		log.Fatal(err)
	}
	voice, err := mp3.Decode(audioContext, bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	damagePlayer, err = audio.NewPlayer(audioContext, voice)
	if err != nil {
		log.Fatal(err)
	}
}
