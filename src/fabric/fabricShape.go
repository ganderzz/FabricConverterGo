package fabric

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

type fabricShape struct {
	ShapeType     string        `json:"type"`
	Left          float64       `json:"left"`
	Top           float64       `json:"top"`
	Width         float64       `json:"width"`
	Height        float64       `json:"height"`
	Fill          string        `json:"fill"`
	Stroke        string        `json:"stroke"`
	StrokeWidth   float64       `json:"strokeWidth"`
	Angle         float64       `json:"angle"`
	ScaleX        float64       `json:"scaleX"`
	ScaleY        float64       `json:"scaleY"`
	Radius        float64       `json:"radius"`
	Text          string        `json:"text"`
	FontSize      float64       `json:"fontSize"`
	LineHeight    float64       `json:"lineHeight"`
	X1            float64       `json:"x1"`
	X2            float64       `json:"x2"`
	Y1            float64       `json:"y1"`
	Y2            float64       `json:"y2"`
	StrokeLineCap string        `json:"strokeLineCap"`
	Objects       []fabricShape `json:"objects"`
}

const (
	circle    = "circle"
	rectangle = "rect"
	line      = "line"
	text      = "i-text"
	group     = "group"
)

func (s *fabricShape) drawShapeType(ctx *gg.Context) {
	switch s.ShapeType {
	case group:
		ctx.Push()
		ctx.Translate(s.Left, s.Top)
		ctx.RotateAbout(s.Angle, s.Left, s.Top)
		for i := len(s.Objects) - 1; i >= 0; i-- {
			s.Objects[i].Parse(ctx)
		}
		ctx.Pop()
		break

	case circle:
		ctx.DrawCircle(s.Left+s.Radius, s.Top+s.Radius, s.Radius)
		break

	case rectangle:
		ctx.DrawRectangle(s.Left, s.Top, s.Width, s.Height)
		break

	case line:
		ctx.Translate(s.Left, s.Top)
		setLineCap(ctx, *s)
		ctx.DrawLine(s.X1, s.Y1, s.X2, s.Y2)
		break

	case text:
		fnt, err := loadFont(&truetype.Options{Size: s.FontSize, DPI: 100})
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		ctx.SetFontFace(fnt)
		ctx.SetHexColor(s.Fill)
		ctx.DrawStringWrapped(s.Text, s.Left, s.Top, 0, 0, s.Width, s.LineHeight, gg.AlignLeft)
		break
	}
}

func setLineCap(ctx *gg.Context, shape fabricShape) {
	switch strings.ToLower(shape.StrokeLineCap) {
	case "butt":
		ctx.SetLineCapButt()
		break
	}
}

func loadFont(options *truetype.Options) (font.Face, error) {
	font, err := truetype.Parse(goregular.TTF)

	if err != nil {
		return nil, err
	}

	if options == nil {
		options = &truetype.Options{Size: 32}
	}

	return truetype.NewFace(font, options), nil
}

/**
 * setColor
 * Parses a string color (rgba or hex), and applies that color to the drawing context
 */
func setColor(ctx *gg.Context, color string) {
	rgbRegex, _ := regexp.Compile("rgb\\(\\s*(\\d{1,3})\\s*,\\s*(\\d{1,3})\\s*,\\s*(\\d{1,3})\\s*\\)")
	hexRegex, _ := regexp.Compile("^#")

	if hexRegex.MatchString(color) {
		ctx.SetHexColor(color)
	} else if rgbRegex.MatchString(color) {
		rgb := rgbRegex.FindStringSubmatch(color)
		r, _ := strconv.Atoi(rgb[1])
		g, _ := strconv.Atoi(rgb[2])
		b, _ := strconv.Atoi(rgb[3])

		ctx.SetRGB255(r, g, b)
	}
}

// Adds the current shape to the gg canvas
func (s *fabricShape) Parse(ctx *gg.Context) {
	ctx.Push()
	if s.ScaleX != 1 || s.ScaleY != 1 {
		ctx.ScaleAbout(s.ScaleX, s.ScaleY, s.Left, s.Top)
	}

	if s.Angle != 0 {
		ctx.RotateAbout(gg.Radians(s.Angle), s.Left, s.Top)
	}

	s.drawShapeType(ctx)

	if len(s.Stroke) > 0 && strings.ToLower(s.Stroke) != "transparent" {
		setColor(ctx, s.Stroke)
		ctx.SetLineWidth(s.StrokeWidth)

		ctx.StrokePreserve()
	}

	if len(s.Fill) > 0 && strings.ToLower(s.Fill) != "transparent" {
		setColor(ctx, s.Fill)
		ctx.Fill()
	}

	ctx.Pop()
}
