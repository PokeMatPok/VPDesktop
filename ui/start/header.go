package start

import (
	"fmt"
	"image"
	"image/color"
	"vpdesktop/types"

	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/layout"
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

var logoText richtext.InteractiveText
var settingsButton widget.Clickable

func Header(gtx layout.Context, th material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {
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

	for settingsButton.Clicked(gtx) {
		fmt.Println("Settings button clicked")
	}

	if settingsButton.Hovered() {
		pointer.CursorPointer.Add(gtx.Ops)
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

			// Center title
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				//lbl := material.H6(th, state.ClassesResponse.Kopf.DatumPlan)

				spans := []richtext.SpanStyle{
					{
						Content: "welcome to\n",
						Color:   color.NRGBA{R: 140, G: 140, B: 140, A: 255},
						Size:    unit.Sp(11),
						Font:    gofont.Collection()[0].Font,
					},
					{
						Content: "VPDesktop",
						Color:   th.Fg,
						Size:    unit.Sp(20),
						Font:    gofont.Collection()[0].Font,
					},
				}

				datetext := richtext.Text(
					&logoText,
					th.Shaper,
					spans...,
				)
				datetext.Alignment = text.Middle

				// padding at the top and bottom
				return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return datetext.Layout(gtx)
				})
			}),

			// Settings button
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return settingsButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return settingsIcon.Layout(gtx, th.Fg)
					})
				})
			}),
		)
	})
}
