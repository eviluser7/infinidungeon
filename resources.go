package main

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/eviluser7/infinidungeon/resources/img"
	"github.com/eviluser7/infinidungeon/resources/sfx"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var (
	// Level 1
	background0I      *ebiten.Image
	background1I      *ebiten.Image
	background2I      *ebiten.Image
	background3I      *ebiten.Image
	background4I      *ebiten.Image
	background5I      *ebiten.Image
	background6I      *ebiten.Image
	background7I      *ebiten.Image
	background8I      *ebiten.Image
	background9I      *ebiten.Image
	background10I     *ebiten.Image
	background11I     *ebiten.Image
	background12I     *ebiten.Image
	background13I     *ebiten.Image
	background14I     *ebiten.Image
	background15I     *ebiten.Image
	background16I     *ebiten.Image
	background17I     *ebiten.Image
	background18preI  *ebiten.Image
	background18postI *ebiten.Image
	background19I     *ebiten.Image
	background20I     *ebiten.Image
	background21I     *ebiten.Image
	background22I     *ebiten.Image
	background23I     *ebiten.Image

	// Character
	charIdle1 *ebiten.Image
	charIdle2 *ebiten.Image
	charIdle3 *ebiten.Image
	charWalk1 *ebiten.Image
	charWalk2 *ebiten.Image
	charWalk3 *ebiten.Image
	blur      *ebiten.Image

	// Env
	blueBlur    *ebiten.Image
	blueShrine  *ebiten.Image
	shrineBlur1 *ebiten.Image

	// Sounds
	audioContext = audio.NewContext(44100)
	enableShrine *audio.Player
)

func loadMaps() {
	var err error

	// Level 1
	imgBg0, _, err := image.Decode(bytes.NewReader(img.Bg0_png))
	if err != nil {
		panic(err)
	}
	background0 = ebiten.NewImageFromImage(imgBg0)

	imgBg1, _, err := image.Decode(bytes.NewReader(img.Bg1_png))
	if err != nil {
		panic(err)
	}
	background1 = ebiten.NewImageFromImage(imgBg1)

	imgBg2, _, err := image.Decode(bytes.NewReader(img.Bg2_png))
	if err != nil {
		panic(err)
	}
	background2 = ebiten.NewImageFromImage(imgBg2)

	imgBg3, _, err := image.Decode(bytes.NewReader(img.Bg3_png))
	if err != nil {
		panic(err)
	}
	background3 = ebiten.NewImageFromImage(imgBg3)

	imgBg4, _, err := image.Decode(bytes.NewReader(img.Bg4_png))
	if err != nil {
		panic(err)
	}
	background4 = ebiten.NewImageFromImage(imgBg4)

	imgBg5, _, err := image.Decode(bytes.NewReader(img.Bg5_png))
	if err != nil {
		panic(err)
	}
	background5 = ebiten.NewImageFromImage(imgBg5)

	imgBg6, _, err := image.Decode(bytes.NewReader(img.Bg6_png))
	if err != nil {
		panic(err)
	}
	background6 = ebiten.NewImageFromImage(imgBg6)

	imgBg7, _, err := image.Decode(bytes.NewReader(img.Bg7_png))
	if err != nil {
		panic(err)
	}
	background7 = ebiten.NewImageFromImage(imgBg7)

	imgBg8, _, err := image.Decode(bytes.NewReader(img.Bg8_png))
	if err != nil {
		panic(err)
	}
	background8 = ebiten.NewImageFromImage(imgBg8)

	imgBg9, _, err := image.Decode(bytes.NewReader(img.Bg9_png))
	if err != nil {
		panic(err)
	}
	background9 = ebiten.NewImageFromImage(imgBg9)

	imgBg10, _, err := image.Decode(bytes.NewReader(img.Bg10_png))
	if err != nil {
		panic(err)
	}
	background10 = ebiten.NewImageFromImage(imgBg10)

	imgBg11, _, err := image.Decode(bytes.NewReader(img.Bg11_png))
	if err != nil {
		panic(err)
	}
	background11 = ebiten.NewImageFromImage(imgBg11)

	imgBg12, _, err := image.Decode(bytes.NewReader(img.Bg12_png))
	if err != nil {
		panic(err)
	}
	background12 = ebiten.NewImageFromImage(imgBg12)

	imgBg13, _, err := image.Decode(bytes.NewReader(img.Bg13_png))
	if err != nil {
		panic(err)
	}
	background13 = ebiten.NewImageFromImage(imgBg13)

	imgBg14, _, err := image.Decode(bytes.NewReader(img.Bg14_png))
	if err != nil {
		panic(err)
	}
	background14 = ebiten.NewImageFromImage(imgBg14)

	imgBg15, _, err := image.Decode(bytes.NewReader(img.Bg15_png))
	if err != nil {
		panic(err)
	}
	background15 = ebiten.NewImageFromImage(imgBg15)

	imgBg16, _, err := image.Decode(bytes.NewReader(img.Bg16_png))
	if err != nil {
		panic(err)
	}
	background16 = ebiten.NewImageFromImage(imgBg16)

	imgBg17, _, err := image.Decode(bytes.NewReader(img.Bg17_png))
	if err != nil {
		panic(err)
	}
	background17 = ebiten.NewImageFromImage(imgBg17)

	imgBg18p, _, err := image.Decode(bytes.NewReader(img.Bg18pre_png))
	if err != nil {
		panic(err)
	}
	background18pre = ebiten.NewImageFromImage(imgBg18p)

	imgBg18ps, _, err := image.Decode(bytes.NewReader(img.Bg18post_png))
	if err != nil {
		panic(err)
	}
	background18post = ebiten.NewImageFromImage(imgBg18ps)

	imgBg19, _, err := image.Decode(bytes.NewReader(img.Bg19_png))
	if err != nil {
		panic(err)
	}
	background19 = ebiten.NewImageFromImage(imgBg19)

	imgBg20, _, err := image.Decode(bytes.NewReader(img.Bg20_png))
	if err != nil {
		panic(err)
	}
	background20 = ebiten.NewImageFromImage(imgBg20)

	imgBg21, _, err := image.Decode(bytes.NewReader(img.Bg21_png))
	if err != nil {
		panic(err)
	}
	background21 = ebiten.NewImageFromImage(imgBg21)

	imgBg22, _, err := image.Decode(bytes.NewReader(img.Bg22_png))
	if err != nil {
		panic(err)
	}
	background22 = ebiten.NewImageFromImage(imgBg22)

	imgBg23, _, err := image.Decode(bytes.NewReader(img.Bg23_png))
	if err != nil {
		panic(err)
	}
	background23 = ebiten.NewImageFromImage(imgBg23)
}

func loadResources() {
	var err error

	imgCharIdle1, _, err := image.Decode(bytes.NewReader(img.CharIdle1_png))
	if err != nil {
		panic(err)
	}
	charIdle1 = ebiten.NewImageFromImage(imgCharIdle1)

	imgCharIdle2, _, err := image.Decode(bytes.NewReader(img.CharIdle2_png))
	if err != nil {
		panic(err)
	}
	charIdle2 = ebiten.NewImageFromImage(imgCharIdle2)

	imgCharIdle3, _, err := image.Decode(bytes.NewReader(img.CharIdle3_png))
	if err != nil {
		panic(err)
	}
	charIdle3 = ebiten.NewImageFromImage(imgCharIdle3)

	imgCharWalk1, _, err := image.Decode(bytes.NewReader(img.CharWalk1_png))
	if err != nil {
		panic(err)
	}
	charWalk1 = ebiten.NewImageFromImage(imgCharWalk1)

	imgCharWalk2, _, err := image.Decode(bytes.NewReader(img.CharWalk2_png))
	if err != nil {
		panic(err)
	}
	charWalk2 = ebiten.NewImageFromImage(imgCharWalk2)

	imgCharWalk3, _, err := image.Decode(bytes.NewReader(img.CharWalk3_png))
	if err != nil {
		panic(err)
	}
	charWalk3 = ebiten.NewImageFromImage(imgCharWalk3)

	imgBlur, _, err := image.Decode(bytes.NewReader(img.Blur_png))
	if err != nil {
		panic(err)
	}
	blur = ebiten.NewImageFromImage(imgBlur)

	// Env
	imgBlueBlur, _, err := image.Decode(bytes.NewReader(img.BlueBlur_png))
	if err != nil {
		panic(err)
	}
	blueBlur = ebiten.NewImageFromImage(imgBlueBlur)

	imgBlueShrineBlur, _, err := image.Decode(bytes.NewReader(img.BlueShrineBlur_png))
	if err != nil {
		panic(err)
	}
	shrineBlur1 = ebiten.NewImageFromImage(imgBlueShrineBlur)

	imgBlueShrine, _, err := image.Decode(bytes.NewReader(img.BlueShrine_png))
	if err != nil {
		panic(err)
	}
	blueShrine = ebiten.NewImageFromImage(imgBlueShrine)
}

func loadSounds() {
	var err error

	shrineSound, err := wav.Decode(audioContext, bytes.NewReader(sfx.EnableShrine_wav))
	if err != nil {
		panic(err)
	}
	enableShrine, err = audio.NewPlayer(audioContext, shrineSound)
}
