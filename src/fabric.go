package main

import (
	"math"
	"strings"

	"github.com/fogleman/gg"
)

type fabricBaseObject struct {
	Version string        `json:"version"`
	Objects []fabricShape `json:"objects"`
}

func (f *fabricBaseObject) GetBounds() (float64, float64) {
	minX := 0.0
	minY := 0.0
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)

	for _, o := range f.Objects {
		// Get the bounds of a shape--keeping in mind its rotation
		width := math.Sin(o.Angle)*o.Height + math.Cos(o.Angle)*o.Width + o.StrokeWidth
		height := math.Sin(o.Angle)*o.Width + math.Cos(o.Angle)*o.Height + o.StrokeWidth

		minX = math.Min(o.Left, minX)
		minY = math.Min(o.Top, minY)
		maxX = math.Max(o.Left+width, maxX)
		maxY = math.Max(o.Top+height, maxY)
	}

	// Doesn't make sense to make an image negative width or height
	if maxX < 0 {
		maxX = 0
	}

	if maxY < 0 {
		maxY = 0
	}

	return maxX - minX, maxY - minY
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
	Angle       float64 `json:"angle"`
	ScaleX      float64 `json:"scaleX"`
	ScaleY      float64 `json:"scaleY"`
	Radius      float64 `json:"radius"`
	Text        string  `json:"text"`
	X1          float64 `json:"x1"`
	X2          float64 `json:"x2"`
	Y1          float64 `json:"y1"`
	Y2          float64 `json:"y2"`
}

const (
	circle    = "circle"
	rectangle = "rect"
	line      = "line"
	text      = "i-text"
)

func (s *fabricShape) drawShapeType(ctx *gg.Context) {
	switch s.ShapeType {
	case circle:
		ctx.DrawCircle(s.Left, s.Top, s.Radius)
		break

	case rectangle:
		ctx.DrawRectangle(s.Left, s.Top, s.Width, s.Height)
		break

	case line:
		ctx.DrawLine(s.X1, s.Y1, s.X2, s.Y2)
		break

	case text:
		ctx.DrawString(s.Text, s.Left, s.Top)
		break
	}
}

func (s *fabricShape) Parse(ctx *gg.Context) {
	if s.ScaleX != 1 || s.ScaleY != 1 {
		ctx.ScaleAbout(s.ScaleX, s.ScaleY, s.Left, s.Top)
	}

	if s.Angle != 0 {
		ctx.RotateAbout(gg.Radians(s.Angle), s.Left, s.Top)
	}

	s.drawShapeType(ctx)

	if len(s.Stroke) > 0 && strings.ToLower(s.Stroke) != "transparent" {
		ctx.SetHexColor(s.Stroke)
		ctx.SetLineWidth(s.StrokeWidth)

		ctx.StrokePreserve()
	}

	if len(s.Fill) > 0 && strings.ToLower(s.Fill) != "transparent" {
		ctx.SetHexColor(s.Fill)
		ctx.Fill()
	}

	if s.Angle != 0 {
		ctx.RotateAbout(-gg.Radians(s.Angle), s.Left, s.Top)
	}

	if s.ScaleX != 1 || s.ScaleY != 1 {
		ctx.ScaleAbout(1, 1, s.Left, s.Top)
	}
}
