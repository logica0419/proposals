package main

import (
	"image"
	"image/gif"

	"golang.org/x/image/draw"
)

func Resize2(src *gif.GIF, width, height int) *gif.GIF {
	srcWidth, srcHeight := src.Config.Width, src.Config.Height
	widthRatio := float64(width) / float64(srcWidth)
	heightRatio := float64(height) / float64(srcHeight)

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
		dstFrameBounds := image.Rect(
			int(float64(srcFrameBounds.Min.X)*widthRatio),
			int(float64(srcFrameBounds.Min.Y)*heightRatio),
			int(float64(srcFrameBounds.Max.X)*widthRatio),
			int(float64(srcFrameBounds.Max.Y)*heightRatio),
		)

		dstFrame := image.NewPaletted(dstFrameBounds, srcFrame.Palette)
		draw.CatmullRom.Scale(dstFrame, dstFrameBounds, srcFrame, srcFrameBounds, draw.Src, nil)

		dst.Image = append(dst.Image, dstFrame)
	}

	return dst
}
