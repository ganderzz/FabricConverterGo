package main

import (
	"strings"

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

func (s *fabricShape) Parse(ctx *gg.Context) {
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

	if len(s.Fill) > 0 && strings.ToLower(s.Fill) != "transparent" {
		ctx.SetHexColor(s.Fill)
		ctx.Fill()
	}

	if len(s.Stroke) > 0 && strings.ToLower(s.Stroke) != "transparent" {
		ctx.SetHexColor(s.Stroke)
		ctx.SetLineWidth(s.StrokeWidth)

		ctx.Stroke()
	}
}
