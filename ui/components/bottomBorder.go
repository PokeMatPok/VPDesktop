package components

import (
	"image"
	"image/color"
	"math"
	"time"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type ReactiveBottomBorder struct {
	// configuration
	Height     int
	Color      color.NRGBA
	FocusColor color.NRGBA
	SpeedIn    float32
	SpeedOut   float32

	// internal
	progress float32
	lastTick time.Time
}

func (b *ReactiveBottomBorder) Layout(
	gtx layout.Context,
	focused bool,
	w layout.Widget,
) layout.Dimensions {
	if b.lastTick.IsZero() {
		b.lastTick = gtx.Now
	}

	dt := float32(gtx.Now.Sub(b.lastTick).Seconds())
	b.lastTick = gtx.Now

	if focused {
		b.progress = float32(
			math.Min(
				float64(b.progress+b.SpeedIn*dt),
				1.0,
			),
		)
	} else {
		b.progress = float32(
			math.Max(
				float64(b.progress-b.SpeedOut*dt),
				0.0,
			),
		)
	}

	// place redraw request if animation is not complete
	if (focused && b.progress < 1.0) || (!focused && b.progress > 0.0) {
		gtx.Execute(op.InvalidateCmd{})
	}

	var dims layout.Dimensions
	dims = layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			// layout child first
			dims = w(gtx)
			return dims
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			// draw bottom border using child dimensions
			r := image.Rect(
				0,
				dims.Size.Y-b.Height,
				dims.Size.X,
				dims.Size.Y,
			)

			paint.FillShape(gtx.Ops, b.Color, clip.Rect(r).Op())

			rr := image.Rect(
				0,
				dims.Size.Y-b.Height,
				max(0, int(float32(dims.Size.X)*float32(b.progress))),
				dims.Size.Y,
			)

			paint.FillShape(gtx.Ops, b.FocusColor, clip.Rect(rr).Op())

			return layout.Dimensions{Size: dims.Size}
		}),
	)
	return dims
}

func (b *ReactiveBottomBorder) ProgressValue() float32 {
	return b.progress
}
