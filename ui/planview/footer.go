package planview

import (
	"fmt"
	"image"
	"vpdesktop/types"

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

var dateButton widget.Clickable
var viewButton widget.Clickable
var viewIcon *widget.Icon
var err error

func Footer(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {
	const footerHeight = 56

	// Background
	paint.FillShape(
		gtx.Ops,
		th.Bg,
		clip.RRect{
			Rect: image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Dp(footerHeight)),
			NE:   gtx.Dp(0),
			NW:   gtx.Dp(0),
			SE:   gtx.Dp(0),
			SW:   gtx.Dp(0),
		}.Op(gtx.Ops),
	)

	for dateButton.Clicked(gtx) {
		fmt.Println("Date button clicked")
	}

	for viewButton.Clicked(gtx) {
		fmt.Println("View button clicked")

		if state.ActiveUI == "dayview" {
			state.ActiveUI = "weekview"
		} else {
			state.ActiveUI = "dayview"
		}

		gtx.Execute(op.InvalidateCmd{})
	}

	if dateButton.Hovered() {
		pointer.CursorNotAllowed.Add(gtx.Ops)
	}

	if viewButton.Hovered() {
		pointer.CursorPointer.Add(gtx.Ops)
	}

	dateIcon, err := widget.NewIcon(icons.ActionDateRange)

	var viewIcon *widget.Icon
	switch state.ActiveUI {
	case "dayview":
		viewIcon, err = widget.NewIcon(icons.ActionViewDay)
		if err != nil {
			panic(err)
		}
	case "weekview":
		viewIcon, err = widget.NewIcon(icons.ActionViewWeek)
		if err != nil {
			panic(err)
		}

	default:
		viewIcon, err = widget.NewIcon(icons.ActionViewDay)
		if err != nil {
			panic(err)
		}
	}

	return layout.Inset{
		Left:  unit.Dp(16),
		Right: unit.Dp(16),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.Y = int(unit.Dp(footerHeight))

		return layout.Flex{
			Axis:      layout.Horizontal,
			Alignment: layout.Middle,
			Spacing:   layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return dateButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return dateIcon.Layout(gtx, th.ContrastFg)
					})
				})
			}),

			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return viewButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return viewIcon.Layout(gtx, th.ContrastFg)
					})
				})
			}),
		)
	})
}
