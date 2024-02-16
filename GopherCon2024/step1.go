package main

import (
	"image"
	"image/gif"

	"golang.org/x/image/draw"
)

func Resize1(src *gif.GIF, width, height int) *gif.GIF {
	dst := &gif.GIF{
		Image:     make([]*image.Paletted, 0, len(src.Image)),
		Delay:     src.Delay,
		LoopCount: src.LoopCount,
		Disposal:  src.Disposal,
		Config: image.Config{
			ColorModel: src.Config.ColorModel,
			Width:      width,
			Height:     height,
		},
		BackgroundIndex: src.BackgroundIndex,
	}

	for _, srcFrame := range src.Image {
		srcFrameBounds := srcFrame.Bounds()
		dstFrameBounds := image.Rect(0, 0, width, height)

		dstFrame := image.NewPaletted(dstFrameBounds, srcFrame.Palette)
		draw.CatmullRom.Scale(dstFrame, dstFrameBounds, srcFrame, srcFrameBounds, draw.Src, nil)

		dst.Image = append(dst.Image, dstFrame)
	}

	return dst
}
