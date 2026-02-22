package planview

import (
	"fmt"
	"image"
	"image/color"
	"vpdesktop/types"

	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/richtext"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var dateText richtext.InteractiveText
var homeButton widget.Clickable
var settingsButton widget.Clickable

func Header(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {
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

	for homeButton.Clicked(gtx) {
		state.ActiveUI = "start"
		gtx.Execute(op.InvalidateCmd{})
	}

	for settingsButton.Clicked(gtx) {
		fmt.Println("Settings button clicked")
	}

	if homeButton.Hovered() {
		pointer.CursorPointer.Add(gtx.Ops)
	}

	if settingsButton.Hovered() {
		pointer.CursorNotAllowed.Add(gtx.Ops)
	}

	homeIcon, err := widget.NewIcon(icons.ActionHome)
	if err != nil {
		panic(err)
	}

	settingsIcon, err := widget.NewIcon(icons.ActionSettings)
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
				return homeButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return homeIcon.Layout(gtx, th.ContrastFg)
					})
				})
			}),

			// Center title
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				//lbl := material.H6(th, state.ClassesResponse.Kopf.DatumPlan)

				spans := []richtext.SpanStyle{
					{
						Content: localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "header_next_day"}),
						Color:   color.NRGBA{R: 140, G: 140, B: 140, A: 255},
						Size:    unit.Sp(11),
						Font:    gofont.Collection()[0].Font,
					},
					{
						Content: "\n" + state.NextDayText,
						Color:   th.ContrastFg,
						Size:    unit.Sp(16),
						Font:    gofont.Collection()[0].Font,
					},
				}

				datetext := richtext.Text(
					&dateText,
					th.Shaper,

					spans...,
				)
				datetext.Alignment = text.Middle

				// padding at the top and bottom
				return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return datetext.Layout(gtx)
				})
			}),

			// Logout button
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return settingsButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return settingsIcon.Layout(gtx, th.ContrastFg)
					})
				})
			}),
		)
	})
}
