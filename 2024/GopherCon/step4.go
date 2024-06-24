package main

import (
	"image"
	"image/color"
	"image/gif"

	"golang.org/x/image/draw"
)

func Resize4(src *gif.GIF, width, height int) *gif.GIF {
	srcWidth, srcHeight := src.Config.Width, src.Config.Height
	widthRatio := float64(width) / float64(srcWidth)
	heightRatio := float64(height) / float64(srcHeight)

	srcBounds := image.Rect(0, 0, srcWidth, srcHeight)
	dstBounds := image.Rect(0, 0, width, height)

	tempCanvas := image.NewNRGBA(srcBounds)
	bgColor := image.NewUniform(src.Config.ColorModel.(color.Palette)[src.BackgroundIndex])

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

	for i, srcFrame := range src.Image {
		srcFrameBounds := srcFrame.Bounds()
		dstFrameBounds := image.Rect(
			int(float64(srcFrameBounds.Min.X)*widthRatio),
			int(float64(srcFrameBounds.Min.Y)*heightRatio),
			int(float64(srcFrameBounds.Max.X)*widthRatio),
			int(float64(srcFrameBounds.Max.Y)*heightRatio),
		)

		stackedFrame := image.NewNRGBA(srcBounds)
		switch src.Disposal[i] {
		case gif.DisposalNone:
			draw.Draw(tempCanvas, srcFrameBounds, srcFrame, srcFrameBounds.Min, draw.Over)
			stackedFrame = deepCopyImage(tempCanvas)

		case gif.DisposalBackground:
			stackedFrame = deepCopyImage(tempCanvas)
			draw.Draw(stackedFrame, srcFrameBounds, srcFrame, srcFrameBounds.Min, draw.Over)

			// If the transparent color is in the frame palette, use it as the background color
			r, g, b, a := srcFrame.Palette[srcFrame.Palette.Index(color.Transparent)].RGBA()
			if r == 0 && g == 0 && b == 0 && a == 0 {
				draw.Draw(tempCanvas, srcBounds, image.Transparent, srcBounds.Min, draw.Src)
			} else {
				draw.Draw(tempCanvas, srcBounds, bgColor, srcBounds.Min, draw.Src)
			}

		case gif.DisposalPrevious:
			stackedFrame = deepCopyImage(tempCanvas)
			draw.Draw(stackedFrame, srcFrameBounds, srcFrame, srcFrameBounds.Min, draw.Over)
		}

		fittedFrame := image.NewPaletted(dstBounds, srcFrame.Palette)
		draw.CatmullRom.Scale(fittedFrame, dstBounds, stackedFrame, srcBounds, draw.Src, nil)

		dstFrame := image.NewPaletted(dstFrameBounds, srcFrame.Palette)
		draw.Draw(dstFrame, dstFrameBounds, fittedFrame, dstFrameBounds.Min, draw.Src)

		dst.Image = append(dst.Image, fittedFrame)
	}

	return dst
}

func deepCopyImage(src *image.NRGBA) *image.NRGBA {
	dst := &image.NRGBA{
		Pix:    make([]uint8, len(src.Pix)),
		Stride: src.Stride,
		Rect:   src.Rect,
	}
	copy(dst.Pix, src.Pix)

	return dst
}
