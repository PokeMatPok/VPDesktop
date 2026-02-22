package other

import (
	"image"
	"image/color"
	"vpdesktop/types"
	"vpdesktop/utils"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var continueButton widget.Clickable
var continueIcon *widget.Icon

func DrawsampleDataUI(gtx layout.Context, th material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	if continueButton.Clicked(gtx) {
		utils.LoadSampleData(state)
		state.ActiveUI = "class_select"
		gtx.Execute(op.InvalidateCmd{})
	}

	continueIcon, err := widget.NewIcon(icons.HardwareKeyboardArrowRight)
	if err != nil {
		panic(err)
	}

	size := gtx.Constraints.Max

	// Background gradient
	paint.LinearGradientOp{
		Stop1:  f32.Pt(0, 0),
		Color1: color.NRGBA{R: 38, G: 43, B: 51, A: 255},
		Stop2:  f32.Pt(float32(size.X), float32(size.Y)),
		Color2: color.NRGBA{R: 23, G: 26, B: 31, A: 255},
	}.Add(gtx.Ops)

	paint.PaintOp{}.Add(gtx.Ops)

	border := widget.Border{
		Color:        color.NRGBA{R: 255, G: 255, B: 255, A: 50},
		Width:        1,
		CornerRadius: unit.Dp(10),
	}

	buttonBorder := widget.Border{
		Color:        color.NRGBA{R: 255, G: 255, B: 255, A: 50},
		Width:        1,
		CornerRadius: unit.Dp(10),
	}

	return layout.Inset{
		Top:    unit.Dp(20),
		Bottom: unit.Dp(20),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

		return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

			return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

				gtx.Constraints.Min.X = 300
				gtx.Constraints.Max.X = 500

				paint.FillShape(
					gtx.Ops,
					color.NRGBA{R: 48, G: 54, B: 64, A: 255},
					clip.RRect{
						Rect: image.Rectangle{Max: gtx.Constraints.Max},
						NE:   gtx.Dp(10), NW: gtx.Dp(10),
						SE: gtx.Dp(10), SW: gtx.Dp(10),
					}.Op(gtx.Ops),
				)

				return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle, Spacing: layout.SpaceEvenly}.Layout(gtx,

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Body1(&th, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_title"})).Layout(gtx)
						})
					}),
					layout.Flexed(0.6, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Left:   unit.Dp(20),
							Right:  unit.Dp(20),
							Bottom: unit.Dp(10),
							Top:    unit.Dp(20),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Body2(&th, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_body"})).Layout(gtx)
						})
					}),
					layout.Rigid(
						func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{
								Top:    unit.Dp(20),
								Bottom: unit.Dp(20),
							}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return buttonBorder.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return continueButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										return layout.Inset{
											Left:   unit.Dp(40),
											Right:  unit.Dp(40),
											Top:    unit.Dp(10),
											Bottom: unit.Dp(10),
										}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											pointer.CursorPointer.Add(gtx.Ops)

											return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return continueIcon.Layout(gtx, color.NRGBA{R: 255, G: 255, B: 255, A: 120})
												}),
												layout.Rigid(func(gtx layout.Context) layout.Dimensions {
													return material.Body1(&th, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_continue_button"})).Layout(gtx)
												}),
											)
										})
									})
								})
							})
						},
					),
				)
			})
		})
	})
}
