package ui

import (
	"image"
	"vpdesktop/types"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

func Header(gtx layout.Context, th *material.Theme, state *types.AppState) layout.Dimensions {
	const headerHeight = 56

	// Background
	paint.FillShape(
		gtx.Ops,
		th.Bg,
		clip.RRect{
			Rect: image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Dp(headerHeight)),
			NE:   gtx.Dp(0),
			NW:   gtx.Dp(0),
			SE:   gtx.Dp(0),
			SW:   gtx.Dp(0),
		}.Op(gtx.Ops),
	)

	homeIcon, err := widget.NewIcon(icons.ActionHome)
	if err != nil {
		panic(err)
	}

	logOutIcon, err := widget.NewIcon(icons.ActionSettings)
	if err != nil {
		panic(err)
	}

	return layout.Inset{
		Left:  unit.Dp(16),
		Right: unit.Dp(16),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.Y = int(unit.Dp(headerHeight))

		return layout.Flex{
			Axis:      layout.Horizontal,
			Alignment: layout.Middle,
		}.Layout(gtx,

			// Left icon
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return homeIcon.Layout(gtx, th.Fg)
				})
			}),

			// Center title
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				lbl := material.H6(th, state.ClassesResponse.Kopf.DatumPlan)
				lbl.Alignment = text.Middle
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return lbl.Layout(gtx)
				})
			}),

			// Logout button
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return logOutIcon.Layout(gtx, th.Fg)
				})
			}),
		)
	})
}
