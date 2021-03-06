package main

import (
	"image/color"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"os"
)

// 调料板
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // 第一个颜色的位置
	blackIndex = 1 // 第二个颜色的位置
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles = 5 // 半径
		res = 0.001 // 粒度
		size = 100
		nFrames = 64 // 帧数
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles * 2 * math.Pi; t += res {
			x := math.Sin(t);
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}