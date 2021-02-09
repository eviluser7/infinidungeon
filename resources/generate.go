//go:generate file2byteslice -package=img -input=./img/0.png -output=./img/bg0.go -var=Bg0_png
//go:generate file2byteslice -package=img -input=./img/1.png -output=./img/bg1.go -var=Bg1_png
//go:generate file2byteslice -package=img -input=./img/2.png -output=./img/bg2.go -var=Bg2_png
//go:generate file2byteslice -package=img -input=./img/3.png -output=./img/bg3.go -var=Bg3_png
//go:generate file2byteslice -package=img -input=./img/4.png -output=./img/bg4.go -var=Bg4_png
//go:generate file2byteslice -package=img -input=./img/5.png -output=./img/bg5.go -var=Bg5_png
//go:generate file2byteslice -package=img -input=./img/6.png -output=./img/bg6.go -var=Bg6_png
//go:generate file2byteslice -package=img -input=./img/7.png -output=./img/bg7.go -var=Bg7_png
//go:generate file2byteslice -package=img -input=./img/8.png -output=./img/bg8.go -var=Bg8_png
//go:generate file2byteslice -package=img -input=./img/9.png -output=./img/bg9.go -var=Bg9_png
//go:generate file2byteslice -package=img -input=./img/10.png -output=./img/bg10.go -var=Bg10_png
//go:generate file2byteslice -package=img -input=./img/11.png -output=./img/bg11.go -var=Bg11_png
//go:generate file2byteslice -package=img -input=./img/12.png -output=./img/bg12.go -var=Bg12_png
//go:generate file2byteslice -package=img -input=./img/13.png -output=./img/bg13.go -var=Bg13_png
//go:generate file2byteslice -package=img -input=./img/14.png -output=./img/bg14.go -var=Bg14_png
//go:generate file2byteslice -package=img -input=./img/15.png -output=./img/bg15.go -var=Bg15_png
//go:generate file2byteslice -package=img -input=./img/16.png -output=./img/bg16.go -var=Bg16_png
//go:generate file2byteslice -package=img -input=./img/17.png -output=./img/bg17.go -var=Bg17_png
//go:generate file2byteslice -package=img -input=./img/18pre.png -output=./img/bg18pre.go -var=Bg18pre_png
//go:generate file2byteslice -package=img -input=./img/18post.png -output=./img/bg18post.go -var=Bg18post_png
//go:generate file2byteslice -package=img -input=./img/19.png -output=./img/bg19.go -var=Bg19_png
//go:generate file2byteslice -package=img -input=./img/20.png -output=./img/bg20.go -var=Bg20_png
//go:generate file2byteslice -package=img -input=./img/21.png -output=./img/bg21.go -var=Bg21_png
//go:generate file2byteslice -package=img -input=./img/22.png -output=./img/bg22.go -var=Bg22_png
//go:generate file2byteslice -package=img -input=./img/char_idle_1.png -output=./img/charIdle1.go -var=CharIdle1_png
//go:generate file2byteslice -package=img -input=./img/char_idle_2.png -output=./img/charIdle2.go -var=CharIdle2_png
//go:generate file2byteslice -package=img -input=./img/char_idle_3.png -output=./img/charIdle3.go -var=CharIdle3_png
//go:generate file2byteslice -package=img -input=./img/char_walk_1.png -output=./img/charWalk1.go -var=CharWalk1_png
//go:generate file2byteslice -package=img -input=./img/char_walk_2.png -output=./img/charWalk2.go -var=CharWalk2_png
//go:generate file2byteslice -package=img -input=./img/char_walk_3.png -output=./img/charWalk3.go -var=CharWalk3_png
//go:generate file2byteslice -package=img -input=./img/blur.png -output=./img/blur.go -var=Blur_png
//go:generate file2byteslice -package=img -input=./img/blue_blur.png -output=./img/blueBlur.go -var=BlueBlur_png
//go:generate file2byteslice -package=img -input=./img/blueShrine.png -output=./img/blueShrine.go -var=BlueShrine_png
//go:generate file2byteslice -package=img -input=./img/shrineBlur1.png -output=./img/shrineBlur1.go -var=BlueShrineBlur_png
//go:generate file2byteslice -package=sfx -input=./sfx/enableShrine.wav -output=./sfx/enableShrine.go -var=EnableShrine_wav
//go:generate gofmt -s -w .

package resources

import (
	// Dummy imports for go.mod for some Go files with 'ignore' tags. For example, `go mod tidy` does not
	// recognize Go files with 'ignore' build tag.
	//
	// Note that this affects only importing this package, but not 'file2byteslice' commands in //go:generate.
	_ "github.com/hajimehoshi/file2byteslice"
)
