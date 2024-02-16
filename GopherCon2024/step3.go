package main

import (
	"image"
	"image/gif"

	"golang.org/x/image/draw"
)

func Resize3(src *gif.GIF, width, height int) *gif.GIF {
	srcWidth, srcHeight := src.Config.Width, src.Config.Height
	widthRatio := float64(width) / float64(srcWidth)
	heightRatio := float64(height) / float64(srcHeight)

	srcBounds := image.Rect(0, 0, srcWidth, srcHeight)
	dstBounds := image.Rect(0, 0, width, height)

	tempCanvas := image.NewRGBA(srcBounds)

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

		draw.Draw(tempCanvas, srcFrameBounds, srcFrame, srcFrameBounds.Min, draw.Over)

		fittedFrame := image.NewPaletted(dstBounds, srcFrame.Palette)
		draw.CatmullRom.Scale(fittedFrame, dstBounds, tempCanvas, srcBounds, draw.Src, nil)

		dstFrame := image.NewPaletted(dstFrameBounds, srcFrame.Palette)
		draw.Draw(dstFrame, dstFrameBounds, fittedFrame, dstFrameBounds.Min, draw.Src)

		dst.Image = append(dst.Image, fittedFrame)
	}

	return dst
}
