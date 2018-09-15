package main

import (
	"image/color"
	"strconv"

	"github.com/fogleman/gg"
)

type fabricBaseObject struct {
	Version string        `json:"version"`
	Objects []fabricShape `json:"objects"`
}

type fabricShape struct {
	ShapeType   string  `json:"type"`
	Left        float64 `json:"left"`
	Top         float64 `json:"top"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	Fill        string  `json:"fill"`
	Stroke      string  `json:"stroke"`
	StrokeWidth float64 `json:"strokeWidth"`
	Radius      float64 `json:"radius"`
	X1          float64 `json:"x1"`
	X2          float64 `json:"x2"`
	Y1          float64 `json:"y1"`
	Y2          float64 `json:"y2"`
}

func HexToRgb(hex string) color.RGBA {
	s := hex
	if hex[0] == '#' {
		s = hex[0:]
	}

	r, _ := strconv.ParseUint(s[0:2], 16, 8)
	b, _ := strconv.ParseUint(s[2:4], 16, 8)
	g, _ := strconv.ParseUint(s[4:6], 16, 8)

	return color.RGBA{
		A: 1,
		R: uint8(r),
		B: uint8(b),
		G: uint8(g)}
}

func (s *fabricShape) Parse(ctx *gg.Context) {
	if len(s.Fill) > 0 {
		rgb := HexToRgb(s.Fill)
		ctx.SetRGB(float64(rgb.R), float64(rgb.B), float64(rgb.G))
		ctx.Fill()
	}

	if len(s.Stroke) > 0 {
		rgb := HexToRgb(s.Stroke)
		ctx.SetRGB(float64(rgb.R), float64(rgb.B), float64(rgb.G))
		ctx.SetLineWidth(s.StrokeWidth)

		ctx.Stroke()
	}

	switch s.ShapeType {
	case "circle":
		ctx.DrawCircle(s.Left, s.Top, s.Radius)
		break

	case "rect":
		ctx.DrawRectangle(s.Left, s.Top, s.Width, s.Height)
		break

	case "line":
		ctx.DrawLine(s.X1, s.Y1, s.X2, s.Y2)
		break
	}

}
