package fabric

import (
	"math"
)

// FabricBaseObject Parent object of the canvas JSON object
type FabricBaseObject struct {
	Version string        `json:"version"`
	Objects []fabricShape `json:"objects"`
}

// GetBounds Gets the size of the fabric image
func (f *FabricBaseObject) GetBounds() (float64, float64) {
	minX := 0.0
	minY := 0.0
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)

	for _, o := range f.Objects {
		// Get the bounds of a shape--keeping in mind its rotation
		width := math.Sin(o.Angle)*o.Height + math.Cos(o.Angle)*o.Width + o.StrokeWidth
		height := math.Sin(o.Angle)*o.Width + math.Cos(o.Angle)*o.Height + o.StrokeWidth

		left := o.Left
		top := o.Top

		if o.ShapeType == line {
			left = o.X1
			top = o.Y1
		}

		minX = math.Min(left, minX)
		minY = math.Min(top, minY)
		maxX = math.Max(left+width, maxX)
		maxY = math.Max(top+height, maxY)
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
