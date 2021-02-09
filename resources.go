package main

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/eviluser7/infinidungeon/resources/img"
	"github.com/eviluser7/infinidungeon/resources/img/stage1"
	"github.com/eviluser7/infinidungeon/resources/img/stage2"
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

	// Level 2
	background0II      *ebiten.Image
	background1II      *ebiten.Image
	background2II      *ebiten.Image
	background3II      *ebiten.Image
	background4II      *ebiten.Image
	background5II      *ebiten.Image
	background6II      *ebiten.Image
	background7II      *ebiten.Image
	background8II      *ebiten.Image
	background9II      *ebiten.Image
	background10II     *ebiten.Image
	background11II     *ebiten.Image
	background12II     *ebiten.Image
	background13II     *ebiten.Image
	background14II     *ebiten.Image
	background15II     *ebiten.Image
	background16II     *ebiten.Image
	background17II     *ebiten.Image
	background18preII  *ebiten.Image
	background18postII *ebiten.Image
	background19II     *ebiten.Image
	background20II     *ebiten.Image
	background21II     *ebiten.Image
	background22II     *ebiten.Image
	background23II     *ebiten.Image

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
	imgBg0, _, err := image.Decode(bytes.NewReader(stage1.Bg0_png))
	if err != nil {
		panic(err)
	}
	background0I = ebiten.NewImageFromImage(imgBg0)

	imgBg1, _, err := image.Decode(bytes.NewReader(stage1.Bg1_png))
	if err != nil {
		panic(err)
	}
	background1I = ebiten.NewImageFromImage(imgBg1)

	imgBg2, _, err := image.Decode(bytes.NewReader(stage1.Bg2_png))
	if err != nil {
		panic(err)
	}
	background2I = ebiten.NewImageFromImage(imgBg2)

	imgBg3, _, err := image.Decode(bytes.NewReader(stage1.Bg3_png))
	if err != nil {
		panic(err)
	}
	background3I = ebiten.NewImageFromImage(imgBg3)

	imgBg4, _, err := image.Decode(bytes.NewReader(stage1.Bg4_png))
	if err != nil {
		panic(err)
	}
	background4I = ebiten.NewImageFromImage(imgBg4)

	imgBg5, _, err := image.Decode(bytes.NewReader(stage1.Bg5_png))
	if err != nil {
		panic(err)
	}
	background5I = ebiten.NewImageFromImage(imgBg5)

	imgBg6, _, err := image.Decode(bytes.NewReader(stage1.Bg6_png))
	if err != nil {
		panic(err)
	}
	background6I = ebiten.NewImageFromImage(imgBg6)

	imgBg7, _, err := image.Decode(bytes.NewReader(stage1.Bg7_png))
	if err != nil {
		panic(err)
	}
	background7I = ebiten.NewImageFromImage(imgBg7)

	imgBg8, _, err := image.Decode(bytes.NewReader(stage1.Bg8_png))
	if err != nil {
		panic(err)
	}
	background8I = ebiten.NewImageFromImage(imgBg8)

	imgBg9, _, err := image.Decode(bytes.NewReader(stage1.Bg9_png))
	if err != nil {
		panic(err)
	}
	background9I = ebiten.NewImageFromImage(imgBg9)

	imgBg10, _, err := image.Decode(bytes.NewReader(stage1.Bg10_png))
	if err != nil {
		panic(err)
	}
	background10I = ebiten.NewImageFromImage(imgBg10)

	imgBg11, _, err := image.Decode(bytes.NewReader(stage1.Bg11_png))
	if err != nil {
		panic(err)
	}
	background11I = ebiten.NewImageFromImage(imgBg11)

	imgBg12, _, err := image.Decode(bytes.NewReader(stage1.Bg12_png))
	if err != nil {
		panic(err)
	}
	background12I = ebiten.NewImageFromImage(imgBg12)

	imgBg13, _, err := image.Decode(bytes.NewReader(stage1.Bg13_png))
	if err != nil {
		panic(err)
	}
	background13I = ebiten.NewImageFromImage(imgBg13)

	imgBg14, _, err := image.Decode(bytes.NewReader(stage1.Bg14_png))
	if err != nil {
		panic(err)
	}
	background14I = ebiten.NewImageFromImage(imgBg14)

	imgBg15, _, err := image.Decode(bytes.NewReader(stage1.Bg15_png))
	if err != nil {
		panic(err)
	}
	background15I = ebiten.NewImageFromImage(imgBg15)

	imgBg16, _, err := image.Decode(bytes.NewReader(stage1.Bg16_png))
	if err != nil {
		panic(err)
	}
	background16I = ebiten.NewImageFromImage(imgBg16)

	imgBg17, _, err := image.Decode(bytes.NewReader(stage1.Bg17_png))
	if err != nil {
		panic(err)
	}
	background17I = ebiten.NewImageFromImage(imgBg17)

	imgBg18p, _, err := image.Decode(bytes.NewReader(stage1.Bg18pre_png))
	if err != nil {
		panic(err)
	}
	background18preI = ebiten.NewImageFromImage(imgBg18p)

	imgBg18ps, _, err := image.Decode(bytes.NewReader(stage1.Bg18post_png))
	if err != nil {
		panic(err)
	}
	background18postI = ebiten.NewImageFromImage(imgBg18ps)

	imgBg19, _, err := image.Decode(bytes.NewReader(stage1.Bg19_png))
	if err != nil {
		panic(err)
	}
	background19I = ebiten.NewImageFromImage(imgBg19)

	imgBg20, _, err := image.Decode(bytes.NewReader(stage1.Bg20_png))
	if err != nil {
		panic(err)
	}
	background20I = ebiten.NewImageFromImage(imgBg20)

	imgBg21, _, err := image.Decode(bytes.NewReader(stage1.Bg21_png))
	if err != nil {
		panic(err)
	}
	background21I = ebiten.NewImageFromImage(imgBg21)

	imgBg22, _, err := image.Decode(bytes.NewReader(stage1.Bg22_png))
	if err != nil {
		panic(err)
	}
	background22I = ebiten.NewImageFromImage(imgBg22)

	imgBg23, _, err := image.Decode(bytes.NewReader(stage1.Bg23_png))
	if err != nil {
		panic(err)
	}
	background23I = ebiten.NewImageFromImage(imgBg23)

	// Level 2
	imgBg0II, _, err := image.Decode(bytes.NewReader(stage2.Bg0_png))
	if err != nil {
		panic(err)
	}
	background0II = ebiten.NewImageFromImage(imgBg0II)

	imgBg1II, _, err := image.Decode(bytes.NewReader(stage2.Bg1_png))
	if err != nil {
		panic(err)
	}
	background1II = ebiten.NewImageFromImage(imgBg1II)

	imgBg2II, _, err := image.Decode(bytes.NewReader(stage2.Bg2_png))
	if err != nil {
		panic(err)
	}
	background2II = ebiten.NewImageFromImage(imgBg2II)

	imgBg3II, _, err := image.Decode(bytes.NewReader(stage2.Bg3_png))
	if err != nil {
		panic(err)
	}
	background3II = ebiten.NewImageFromImage(imgBg3II)

	imgBg4II, _, err := image.Decode(bytes.NewReader(stage2.Bg4_png))
	if err != nil {
		panic(err)
	}
	background4II = ebiten.NewImageFromImage(imgBg4II)

	imgBg5II, _, err := image.Decode(bytes.NewReader(stage2.Bg5_png))
	if err != nil {
		panic(err)
	}
	background5II = ebiten.NewImageFromImage(imgBg5II)

	imgBg6II, _, err := image.Decode(bytes.NewReader(stage2.Bg6_png))
	if err != nil {
		panic(err)
	}
	background6II = ebiten.NewImageFromImage(imgBg6II)

	imgBg7II, _, err := image.Decode(bytes.NewReader(stage2.Bg7_png))
	if err != nil {
		panic(err)
	}
	background7II = ebiten.NewImageFromImage(imgBg7II)

	imgBg8II, _, err := image.Decode(bytes.NewReader(stage2.Bg8_png))
	if err != nil {
		panic(err)
	}
	background8II = ebiten.NewImageFromImage(imgBg8II)

	imgBg9II, _, err := image.Decode(bytes.NewReader(stage2.Bg9_png))
	if err != nil {
		panic(err)
	}
	background9II = ebiten.NewImageFromImage(imgBg9II)

	imgBg10II, _, err := image.Decode(bytes.NewReader(stage2.Bg10_png))
	if err != nil {
		panic(err)
	}
	background10II = ebiten.NewImageFromImage(imgBg10II)

	imgBg11II, _, err := image.Decode(bytes.NewReader(stage2.Bg11_png))
	if err != nil {
		panic(err)
	}
	background11II = ebiten.NewImageFromImage(imgBg11II)

	imgBg12II, _, err := image.Decode(bytes.NewReader(stage2.Bg12_png))
	if err != nil {
		panic(err)
	}
	background12II = ebiten.NewImageFromImage(imgBg12II)

	imgBg13II, _, err := image.Decode(bytes.NewReader(stage2.Bg13_png))
	if err != nil {
		panic(err)
	}
	background13II = ebiten.NewImageFromImage(imgBg13II)

	imgBg14II, _, err := image.Decode(bytes.NewReader(stage2.Bg14_png))
	if err != nil {
		panic(err)
	}
	background14II = ebiten.NewImageFromImage(imgBg14II)

	imgBg15II, _, err := image.Decode(bytes.NewReader(stage2.Bg15_png))
	if err != nil {
		panic(err)
	}
	background15II = ebiten.NewImageFromImage(imgBg15II)

	imgBg16II, _, err := image.Decode(bytes.NewReader(stage2.Bg16_png))
	if err != nil {
		panic(err)
	}
	background16II = ebiten.NewImageFromImage(imgBg16II)

	imgBg17II, _, err := image.Decode(bytes.NewReader(stage2.Bg17_png))
	if err != nil {
		panic(err)
	}
	background17II = ebiten.NewImageFromImage(imgBg17II)

	imgBg18pII, _, err := image.Decode(bytes.NewReader(stage2.Bg18pre_png))
	if err != nil {
		panic(err)
	}
	background18preII = ebiten.NewImageFromImage(imgBg18pII)

	imgBg18psII, _, err := image.Decode(bytes.NewReader(stage2.Bg18post_png))
	if err != nil {
		panic(err)
	}
	background18postII = ebiten.NewImageFromImage(imgBg18psII)

	imgBg19II, _, err := image.Decode(bytes.NewReader(stage2.Bg19_png))
	if err != nil {
		panic(err)
	}
	background19II = ebiten.NewImageFromImage(imgBg19II)

	imgBg20II, _, err := image.Decode(bytes.NewReader(stage2.Bg20_png))
	if err != nil {
		panic(err)
	}
	background20II = ebiten.NewImageFromImage(imgBg20II)

	imgBg21II, _, err := image.Decode(bytes.NewReader(stage2.Bg21_png))
	if err != nil {
		panic(err)
	}
	background21II = ebiten.NewImageFromImage(imgBg21II)

	imgBg22II, _, err := image.Decode(bytes.NewReader(stage2.Bg22_png))
	if err != nil {
		panic(err)
	}
	background22II = ebiten.NewImageFromImage(imgBg22II)

	imgBg23II, _, err := image.Decode(bytes.NewReader(stage2.Bg23_png))
	if err != nil {
		panic(err)
	}
	background23II = ebiten.NewImageFromImage(imgBg23II)
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
