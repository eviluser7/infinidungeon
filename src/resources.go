package main

import (
	"bytes"
	"image"
	"image/png"
	_ "image/png"
	"os"

	"github.com/eviluser7/infinidungeon/resources/fonts"
	"github.com/eviluser7/infinidungeon/resources/img"
	"github.com/eviluser7/infinidungeon/resources/img/stage1"
	"github.com/eviluser7/infinidungeon/resources/img/stage2"
	"github.com/eviluser7/infinidungeon/resources/img/stage3"
	"github.com/eviluser7/infinidungeon/resources/img/stage5"
	"github.com/eviluser7/infinidungeon/resources/img/stageEnd"
	"github.com/eviluser7/infinidungeon/resources/sfx"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"golang.org/x/image/font"
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

	// Level 3
	backgroundIII *ebiten.Image
	exitPreIII    *ebiten.Image
	exitPostIII   *ebiten.Image
	arrowUp       *ebiten.Image
	arrowRight    *ebiten.Image
	arrowDown     *ebiten.Image
	arrowLeft     *ebiten.Image
	arrowRD       *ebiten.Image
	arrowLR       *ebiten.Image
	circle        *ebiten.Image

	// Level 4
	stage4    *ebiten.Image
	greenBlur *ebiten.Image

	// Level 5
	backgroundV     *ebiten.Image
	backgroundVPre  *ebiten.Image
	backgroundVPost *ebiten.Image
	whiteShrine     *ebiten.Image

	// Level 6
	backgroundVI1 *ebiten.Image
	backgroundVI2 *ebiten.Image
	evilStanding1 *ebiten.Image
	evilStanding2 *ebiten.Image
	exclamation   *ebiten.Image

	// Character
	charIdle1 *ebiten.Image
	charIdle2 *ebiten.Image
	charIdle3 *ebiten.Image
	charWalk1 *ebiten.Image
	charWalk2 *ebiten.Image
	charWalk3 *ebiten.Image
	blur      *ebiten.Image
	moreBlur  *ebiten.Image

	// Env
	blueBlur     *ebiten.Image
	yellowBlur   *ebiten.Image
	blueShrine   *ebiten.Image
	yellowShrine *ebiten.Image
	greenShrine  *ebiten.Image
	shrineBlur1  *ebiten.Image
	shrineBlur2  *ebiten.Image
	shrineBlur3  *ebiten.Image
	purpleEffect *ebiten.Image
	blueEffect   *ebiten.Image
	yellowEffect *ebiten.Image

	// External
	wasd     *ebiten.Image
	spaceBar *ebiten.Image
	enterKey *ebiten.Image

	// Menu
	menu           *ebiten.Image
	backgroundMenu *ebiten.Image
	menuCredits    *ebiten.Image
	loadingScreen1 *ebiten.Image
	loadingScreen2 *ebiten.Image
	loadingScreen3 *ebiten.Image
	loadingScreen4 *ebiten.Image
	loadingScreen5 *ebiten.Image
	loadingScreen6 *ebiten.Image

	// Achievements
	achBar            *ebiten.Image
	achBar2           *ebiten.Image
	achBarLocked      *ebiten.Image
	ach1              *ebiten.Image
	ach2              *ebiten.Image
	ach3              *ebiten.Image
	ach4              *ebiten.Image
	ach5              *ebiten.Image
	ach6              *ebiten.Image
	ach7              *ebiten.Image
	ach8              *ebiten.Image
	ach9              *ebiten.Image
	leftArrow         *ebiten.Image
	rightArrow        *ebiten.Image
	achievementButton *ebiten.Image
	achievementText   string

	// Dialogue
	dialogueBox *ebiten.Image

	// Sounds
	audioContext = audio.NewContext(44100)
	enableShrine *audio.Player
	surprise     *audio.Player
	punches      *audio.Player
	dialogue     *audio.Player

	// Music
	music1 *audio.Player
	music2 *audio.Player
	music3 *audio.Player
	music5 *audio.Player

	// Fonts
	pixeledFont      font.Face
	pixeledFontSmall font.Face

	// Window
	icon16 image.Image
	icon32 image.Image
	icon48 image.Image
)

const (
	text1 = `Find the blue shrine`
	text2 = `Make your way to the pond of walls`
	text3 = `Find the yellow shrine`
	text4 = `Reach the valley of arrows`
	text5 = `Find the green shrine`
	text6 = `A rest is always good`
	text7 = `Reach the hole of confunsion`
	text8 = `Find the last shrine`
	text9 = `Finish the game.`
	intro = `massiveNerd presents...`
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

	// Level 3
	imgBgIII, _, err := image.Decode(bytes.NewReader(stage3.BgStage3_png))
	if err != nil {
		panic(err)
	}
	backgroundIII = ebiten.NewImageFromImage(imgBgIII)

	imgExitPreIII, _, err := image.Decode(bytes.NewReader(stage3.ExitStage3pre_png))
	if err != nil {
		panic(err)
	}
	exitPreIII = ebiten.NewImageFromImage(imgExitPreIII)

	imgExitPostIII, _, err := image.Decode(bytes.NewReader(stage3.ExitStage3post_png))
	if err != nil {
		panic(err)
	}
	exitPostIII = ebiten.NewImageFromImage(imgExitPostIII)

	imgArrowUp, _, err := image.Decode(bytes.NewReader(stage3.ArrowUp_png))
	if err != nil {
		panic(err)
	}
	arrowUp = ebiten.NewImageFromImage(imgArrowUp)

	imgArrowRight, _, err := image.Decode(bytes.NewReader(stage3.ArrowRight_png))
	if err != nil {
		panic(err)
	}
	arrowRight = ebiten.NewImageFromImage(imgArrowRight)

	imgArrowDown, _, err := image.Decode(bytes.NewReader(stage3.ArrowDown_png))
	if err != nil {
		panic(err)
	}
	arrowDown = ebiten.NewImageFromImage(imgArrowDown)

	imgArrowLeft, _, err := image.Decode(bytes.NewReader(stage3.ArrowLeft_png))
	if err != nil {
		panic(err)
	}
	arrowLeft = ebiten.NewImageFromImage(imgArrowLeft)

	imgArrowRD, _, err := image.Decode(bytes.NewReader(stage3.ArrowRD_png))
	if err != nil {
		panic(err)
	}
	arrowRD = ebiten.NewImageFromImage(imgArrowRD)

	imgArrowLR, _, err := image.Decode(bytes.NewReader(stage3.ArrowLR_png))
	if err != nil {
		panic(err)
	}
	arrowLR = ebiten.NewImageFromImage(imgArrowLR)

	imgCircle, _, err := image.Decode(bytes.NewReader(stage3.Circle_png))
	if err != nil {
		panic(err)
	}
	circle = ebiten.NewImageFromImage(imgCircle)

	// Level 4 is just one image, really
	imgStage4, _, err := image.Decode(bytes.NewReader(img.Stage4_png))
	if err != nil {
		panic(err)
	}
	stage4 = ebiten.NewImageFromImage(imgStage4)

	imgNicerBlur, _, err := image.Decode(bytes.NewReader(img.NicerBlur_png))
	if err != nil {
		panic(err)
	}
	greenBlur = ebiten.NewImageFromImage(imgNicerBlur)

	// Level 5
	imgStage5, _, err := image.Decode(bytes.NewReader(stage5.Stage5_png))
	if err != nil {
		panic(err)
	}
	backgroundV = ebiten.NewImageFromImage(imgStage5)

	imgStage5Pre, _, err := image.Decode(bytes.NewReader(stage5.Stage5Pre_png))
	if err != nil {
		panic(err)
	}
	backgroundVPre = ebiten.NewImageFromImage(imgStage5Pre)

	imgStage5Post, _, err := image.Decode(bytes.NewReader(stage5.Stage5Post_png))
	if err != nil {
		panic(err)
	}
	backgroundVPost = ebiten.NewImageFromImage(imgStage5Post)

	imgWhiteShrine, _, err := image.Decode(bytes.NewReader(stage5.WhiteShrine_png))
	if err != nil {
		panic(err)
	}
	whiteShrine = ebiten.NewImageFromImage(imgWhiteShrine)

	// Endgame
	imgEndBG1, _, err := image.Decode(bytes.NewReader(stageEnd.EndBG_png))
	if err != nil {
		panic(err)
	}
	backgroundVI1 = ebiten.NewImageFromImage(imgEndBG1)

	imgEndBG2, _, err := image.Decode(bytes.NewReader(stageEnd.EndBG2_png))
	if err != nil {
		panic(err)
	}
	backgroundVI2 = ebiten.NewImageFromImage(imgEndBG2)

	imgEvilStanding1, _, err := image.Decode(bytes.NewReader(stageEnd.MrEvilStanding_png))
	if err != nil {
		panic(err)
	}
	evilStanding1 = ebiten.NewImageFromImage(imgEvilStanding1)

	imgEvilStanding2, _, err := image.Decode(bytes.NewReader(stageEnd.MrEvilStanding2_png))
	if err != nil {
		panic(err)
	}
	evilStanding2 = ebiten.NewImageFromImage(imgEvilStanding2)

	imgExclamation, _, err := image.Decode(bytes.NewReader(stageEnd.Exclamation_png))
	if err != nil {
		panic(err)
	}
	exclamation = ebiten.NewImageFromImage(imgExclamation)
}

func loadResources() {
	var err error

	// Icon
	f16, err := os.Open("../icon16.png")
	if err != nil {
		panic(err)
	}
	defer f16.Close()

	icon16, err = png.Decode(f16)
	if err != nil {
		panic(err)
	}

	f32, err := os.Open("../icon32.png")
	if err != nil {
		panic(err)
	}
	defer f32.Close()

	icon32, err = png.Decode(f32)
	if err != nil {
		panic(err)
	}

	f48, err := os.Open("../icon48.png")
	if err != nil {
		panic(err)
	}
	defer f48.Close()

	icon48, err = png.Decode(f48)
	if err != nil {
		panic(err)
	}

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

	imgMoreBlur, _, err := image.Decode(bytes.NewReader(img.MoreBlur_png))
	if err != nil {
		panic(err)
	}
	moreBlur = ebiten.NewImageFromImage(imgMoreBlur)

	// Env
	imgBlueBlur, _, err := image.Decode(bytes.NewReader(img.BlueBlur_png))
	if err != nil {
		panic(err)
	}
	blueBlur = ebiten.NewImageFromImage(imgBlueBlur)

	imgYellowBlur, _, err := image.Decode(bytes.NewReader(img.YellowBlur_png))
	if err != nil {
		panic(err)
	}
	yellowBlur = ebiten.NewImageFromImage(imgYellowBlur)

	imgBlueShrineBlur, _, err := image.Decode(bytes.NewReader(img.BlueShrineBlur_png))
	if err != nil {
		panic(err)
	}
	shrineBlur1 = ebiten.NewImageFromImage(imgBlueShrineBlur)

	imgYellowShrineBlur, _, err := image.Decode(bytes.NewReader(img.YellowShrineBlur_png))
	if err != nil {
		panic(err)
	}
	shrineBlur2 = ebiten.NewImageFromImage(imgYellowShrineBlur)

	imgBlueShrine, _, err := image.Decode(bytes.NewReader(img.BlueShrine_png))
	if err != nil {
		panic(err)
	}
	blueShrine = ebiten.NewImageFromImage(imgBlueShrine)

	imgYellowShrine, _, err := image.Decode(bytes.NewReader(img.YellowShrine_png))
	if err != nil {
		panic(err)
	}
	yellowShrine = ebiten.NewImageFromImage(imgYellowShrine)

	imgGreenShrine, _, err := image.Decode(bytes.NewReader(stage3.GreenShrine_png))
	if err != nil {
		panic(err)
	}
	greenShrine = ebiten.NewImageFromImage(imgGreenShrine)

	imgGreenBlur, _, err := image.Decode(bytes.NewReader(stage3.GreenShrineBlur_png))
	if err != nil {
		panic(err)
	}
	shrineBlur3 = ebiten.NewImageFromImage(imgGreenBlur)

	imgPurpleEffect, _, err := image.Decode(bytes.NewReader(img.PurpleEffect_png))
	if err != nil {
		panic(err)
	}
	purpleEffect = ebiten.NewImageFromImage(imgPurpleEffect)

	imgBlueEffect, _, err := image.Decode(bytes.NewReader(img.BlueEffect_png))
	if err != nil {
		panic(err)
	}
	blueEffect = ebiten.NewImageFromImage(imgBlueEffect)

	imgYellowEffect, _, err := image.Decode(bytes.NewReader(img.YellowEffect_png))
	if err != nil {
		panic(err)
	}
	yellowEffect = ebiten.NewImageFromImage(imgYellowEffect)

	// External
	imgWasd, _, err := image.Decode(bytes.NewReader(img.Wasd_png))
	if err != nil {
		panic(err)
	}
	wasd = ebiten.NewImageFromImage(imgWasd)

	imgSpacebar, _, err := image.Decode(bytes.NewReader(img.Spacebar_png))
	if err != nil {
		panic(err)
	}
	spaceBar = ebiten.NewImageFromImage(imgSpacebar)

	// Menu
	imgMenu, _, err := image.Decode(bytes.NewReader(img.Menu_png))
	if err != nil {
		panic(err)
	}
	menu = ebiten.NewImageFromImage(imgMenu)

	imgMenuBG, _, err := image.Decode(bytes.NewReader(img.MenuBG_png))
	if err != nil {
		panic(err)
	}
	backgroundMenu = ebiten.NewImageFromImage(imgMenuBG)

	imgMenuText, _, err := image.Decode(bytes.NewReader(img.MenuCredits_png))
	if err != nil {
		panic(err)
	}
	menuCredits = ebiten.NewImageFromImage(imgMenuText)

	imgLoading1, _, err := image.Decode(bytes.NewReader(img.LoadingScreen1_png))
	if err != nil {
		panic(err)
	}
	loadingScreen1 = ebiten.NewImageFromImage(imgLoading1)

	imgLoading2, _, err := image.Decode(bytes.NewReader(img.LoadingScreen2_png))
	if err != nil {
		panic(err)
	}
	loadingScreen2 = ebiten.NewImageFromImage(imgLoading2)

	imgLoading3, _, err := image.Decode(bytes.NewReader(img.LoadingScreen3_png))
	if err != nil {
		panic(err)
	}
	loadingScreen3 = ebiten.NewImageFromImage(imgLoading3)

	imgLoading4, _, err := image.Decode(bytes.NewReader(img.LoadingScreen4_png))
	if err != nil {
		panic(err)
	}
	loadingScreen4 = ebiten.NewImageFromImage(imgLoading4)

	imgLoading5, _, err := image.Decode(bytes.NewReader(img.LoadingScreen5_png))
	if err != nil {
		panic(err)
	}
	loadingScreen5 = ebiten.NewImageFromImage(imgLoading5)

	imgLoading6, _, err := image.Decode(bytes.NewReader(img.LoadingScreen6_png))
	if err != nil {
		panic(err)
	}
	loadingScreen6 = ebiten.NewImageFromImage(imgLoading6)

	// Achievement-related
	imgAchBar, _, err := image.Decode(bytes.NewReader(img.AchievementBar_png))
	if err != nil {
		panic(err)
	}
	achBar = ebiten.NewImageFromImage(imgAchBar)

	imgAchBar2, _, err := image.Decode(bytes.NewReader(img.AchievementBar2_png))
	if err != nil {
		panic(err)
	}
	achBar2 = ebiten.NewImageFromImage(imgAchBar2)

	imgAchBarLocked, _, err := image.Decode(bytes.NewReader(img.AchievementLocked_png))
	if err != nil {
		panic(err)
	}
	achBarLocked = ebiten.NewImageFromImage(imgAchBarLocked)

	imgAch1, _, err := image.Decode(bytes.NewReader(img.Achievement1_png))
	if err != nil {
		panic(err)
	}
	ach1 = ebiten.NewImageFromImage(imgAch1)

	imgAch2, _, err := image.Decode(bytes.NewReader(img.Achievement2_png))
	if err != nil {
		panic(err)
	}
	ach2 = ebiten.NewImageFromImage(imgAch2)

	imgAch3, _, err := image.Decode(bytes.NewReader(img.Achievement3_png))
	if err != nil {
		panic(err)
	}
	ach3 = ebiten.NewImageFromImage(imgAch3)

	imgAch4, _, err := image.Decode(bytes.NewReader(img.Achievement4_png))
	if err != nil {
		panic(err)
	}
	ach4 = ebiten.NewImageFromImage(imgAch4)

	imgAch5, _, err := image.Decode(bytes.NewReader(img.Achievement5_png))
	if err != nil {
		panic(err)
	}
	ach5 = ebiten.NewImageFromImage(imgAch5)

	imgAch6, _, err := image.Decode(bytes.NewReader(img.Achievement6_png))
	if err != nil {
		panic(err)
	}
	ach6 = ebiten.NewImageFromImage(imgAch6)

	imgAch7, _, err := image.Decode(bytes.NewReader(img.Achievement7_png))
	if err != nil {
		panic(err)
	}
	ach7 = ebiten.NewImageFromImage(imgAch7)

	imgAch8, _, err := image.Decode(bytes.NewReader(img.Achievement8_png))
	if err != nil {
		panic(err)
	}
	ach8 = ebiten.NewImageFromImage(imgAch8)

	imgAch9, _, err := image.Decode(bytes.NewReader(img.Achievement9_png))
	if err != nil {
		panic(err)
	}
	ach9 = ebiten.NewImageFromImage(imgAch9)

	imgLeftArrow, _, err := image.Decode(bytes.NewReader(img.LeftArrow_png))
	if err != nil {
		panic(err)
	}
	leftArrow = ebiten.NewImageFromImage(imgLeftArrow)

	imgRightArrow, _, err := image.Decode(bytes.NewReader(img.RightArrow_png))
	if err != nil {
		panic(err)
	}
	rightArrow = ebiten.NewImageFromImage(imgRightArrow)

	imgAchievements, _, err := image.Decode(bytes.NewReader(img.AchievementButton_png))
	if err != nil {
		panic(err)
	}
	achievementButton = ebiten.NewImageFromImage(imgAchievements)

	// Dialogue
	imgDialogueBox, _, err := image.Decode(bytes.NewReader(img.DialogueBox_png))
	if err != nil {
		panic(err)
	}
	dialogueBox = ebiten.NewImageFromImage(imgDialogueBox)

	imgEnterKey, _, err := image.Decode(bytes.NewReader(img.EnterKey_png))
	if err != nil {
		panic(err)
	}
	enterKey = ebiten.NewImageFromImage(imgEnterKey)
}

func loadSounds() {
	var err error

	shrineSound, err := wav.Decode(audioContext, bytes.NewReader(sfx.EnableShrine_wav))
	if err != nil {
		panic(err)
	}
	enableShrine, err = audio.NewPlayer(audioContext, shrineSound)

	surpriseSound, err := wav.Decode(audioContext, bytes.NewReader(sfx.Surprise_wav))
	if err != nil {
		panic(err)
	}
	surprise, err = audio.NewPlayer(audioContext, surpriseSound)

	punchesSound, err := wav.Decode(audioContext, bytes.NewReader(sfx.Punch_wav))
	if err != nil {
		panic(err)
	}
	punches, err = audio.NewPlayer(audioContext, punchesSound)

	dialogueSound, err := wav.Decode(audioContext, bytes.NewReader(sfx.Dialogue_wav))
	if err != nil {
		panic(err)
	}
	dialogue, err = audio.NewPlayer(audioContext, dialogueSound)

	song1, err := wav.Decode(audioContext, bytes.NewReader(sfx.Song1_wav))
	if err != nil {
		panic(err)
	}
	music1, err = audio.NewPlayer(audioContext, song1)

	song2, err := wav.Decode(audioContext, bytes.NewReader(sfx.Song2_wav))
	if err != nil {
		panic(err)
	}
	music2, err = audio.NewPlayer(audioContext, song2)

	song3, err := wav.Decode(audioContext, bytes.NewReader(sfx.Song3_wav))
	if err != nil {
		panic(err)
	}
	music3, err = audio.NewPlayer(audioContext, song3)

	song5, err := wav.Decode(audioContext, bytes.NewReader(sfx.Song5_wav))
	if err != nil {
		panic(err)
	}
	music5, err = audio.NewPlayer(audioContext, song5)
}

func loadFonts() {
	var err error

	tt, err := truetype.Parse(fonts.Pixeled_ttf)

	const dpi = 72
	pixeledFont = truetype.NewFace(tt, &truetype.Options{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}

	tt2, err := truetype.Parse(fonts.Pixeled_ttf)

	pixeledFontSmall = truetype.NewFace(tt2, &truetype.Options{
		Size:    7,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
}
