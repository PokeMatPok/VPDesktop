package planview

import (
	"image/color"
	"vpdesktop/types"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func DrawDayViewUI(gtx layout.Context, th *material.Theme, state *types.AppState) layout.Dimensions {

	border := widget.Border{
		Color: color.NRGBA{R: 210, G: 210, B: 210, A: 255},
		Width: unit.Dp(2),
	}

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		var list layout.List // vertical scroll
		list.Axis = layout.Vertical

		return list.Layout(gtx, len(state.DayViewState.Lessons), func(gtx layout.Context, i int) layout.Dimensions {

			gtx.Constraints.Max.X = min(gtx.Constraints.Max.X, gtx.Dp(600))

			lesson := state.DayViewState.Lessons[i]

			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					// Time column (fixed width)
					return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Right: unit.Dp(8)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return material.Body1(th, lesson.Beginn+" - "+lesson.Ende).Layout(gtx)
							})
						})
					})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					// Subject column (flexible)
					return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Body1(th, lesson.Fa.Value).Layout(gtx)
						})
					})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					// Teacher column (flexible)
					return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Body1(th, lesson.Le.Value).Layout(gtx)
						})
					})
				}),
			)
		})

	})

}
