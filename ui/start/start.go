package start

import (
	"image"
	"image/color"
	"vpdesktop/types"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func DrawStartUI(gtx layout.Context, th material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {
	size := gtx.Constraints.Max

	// Background gradient
	paint.LinearGradientOp{
		Stop1:  f32.Pt(0, 0),
		Color1: color.NRGBA{R: 255, G: 0, B: 0, A: 255},
		Stop2:  f32.Pt(float32(size.X), float32(size.Y)),
		Color2: color.NRGBA{R: 0, G: 0, B: 255, A: 255},
	}.Add(gtx.Ops)

	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return Header(gtx, th, state, localizer)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {

			gtx.Constraints.Max.Y -= gtx.Dp(70)
			return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(10).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						// Draw a colored rectangle as the background of the inset
						cs := gtx.Constraints
						rect := image.Rectangle{
							Min: image.Point{X: 0, Y: 0},
							Max: image.Point{X: cs.Max.X, Y: cs.Max.Y},
						}
						paint.FillShape(gtx.Ops, color.NRGBA{R: 240, G: 240, B: 240, A: 255}, clip.Rect(rect).Op())
						return layout.Dimensions{Size: cs.Max}
					})
				}),
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(10).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						// Draw a colored rectangle as the background of the inset
						cs := gtx.Constraints
						rect := image.Rectangle{
							Min: image.Point{X: 0, Y: 0},
							Max: image.Point{X: cs.Max.X, Y: cs.Max.Y},
						}
						paint.FillShape(gtx.Ops, color.NRGBA{R: 240, G: 240, B: 240, A: 255}, clip.Rect(rect).Op())
						return layout.Dimensions{Size: cs.Max}
					})
				}),
			)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {

			gtx.Constraints.Max.Y -= gtx.Dp(70)
			return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(10).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						// Draw a colored rectangle as the background of the inset
						cs := gtx.Constraints
						rect := image.Rectangle{
							Min: image.Point{X: 0, Y: 0},
							Max: image.Point{X: cs.Max.X, Y: cs.Max.Y},
						}
						paint.FillShape(gtx.Ops, color.NRGBA{R: 240, G: 240, B: 240, A: 255}, clip.Rect(rect).Op())
						return layout.Dimensions{Size: cs.Max}
					})
				}),
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(10).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						// Draw a colored rectangle as the background of the inset
						cs := gtx.Constraints
						rect := image.Rectangle{
							Min: image.Point{X: 0, Y: 0},
							Max: image.Point{X: cs.Max.X, Y: cs.Max.Y},
						}
						paint.FillShape(gtx.Ops, color.NRGBA{R: 240, G: 240, B: 240, A: 255}, clip.Rect(rect).Op())
						return layout.Dimensions{Size: cs.Max}
					})
				}),
			)
		}),
	)
}
