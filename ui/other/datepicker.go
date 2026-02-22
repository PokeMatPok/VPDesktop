package other

import (
	"image"
	"strconv"
	"time"
	"vpdesktop/types"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var datePickerCloseButton widget.Clickable
var datePickerOkButton widget.Clickable
var datePickerDayEditor widget.Editor
var datePickerMonthEditor widget.Editor
var datePickerYearEditor widget.Editor

func DatePickerOverlay(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	datePickerDayEditor.SingleLine = true
	datePickerDayEditor.MaxLen = 2
	datePickerDayEditor.Filter = "0123456789"
	datePickerMonthEditor.SingleLine = true
	datePickerMonthEditor.MaxLen = 2
	datePickerMonthEditor.Filter = "0123456789"
	datePickerYearEditor.SingleLine = true
	datePickerYearEditor.MaxLen = 4
	datePickerYearEditor.Filter = "0123456789"

	editorTheme := material.Theme{
		Shaper:   th.Shaper,
		TextSize: 25,
		Face:     font.Typeface("Times New Roman"),
	}

	editorTheme.Fg = th.Fg

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = gtx.Dp(300)
		gtx.Constraints.Min.Y = gtx.Dp(150)
		gtx.Constraints.Max.X = gtx.Dp(300)
		gtx.Constraints.Max.Y = gtx.Dp(150)

		rr := clip.RRect{
			Rect: image.Rectangle{Max: gtx.Constraints.Max},
			NE:   gtx.Dp(10),
			NW:   gtx.Dp(10),
			SE:   gtx.Dp(10),
			SW:   gtx.Dp(10),
		}
		paint.FillShape(gtx.Ops, th.Bg, rr.Op(gtx.Ops))

		return layout.UniformInset(4).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.H6(th, "Date Picker Placeholder").Layout(gtx)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					dateList := layout.List{Axis: layout.Horizontal}

					currentTime := time.Now()

					return dateList.Layout(gtx, 5, func(gtx layout.Context, index int) layout.Dimensions {
						var dim layout.Dimensions

						switch index {
						case 0:
							dim = material.Editor(&editorTheme, &datePickerDayEditor, strconv.Itoa(currentTime.Day())).Layout(gtx)
						case 1:
							dim = material.Label(th, 15, ":").Layout(gtx)
						case 2:
							dim = material.Editor(&editorTheme, &datePickerMonthEditor, strconv.Itoa(int(currentTime.Month()))).Layout(gtx)
						case 3:
							dim = material.Label(th, 15, ":").Layout(gtx)
						case 4:
							dim = material.Editor(&editorTheme, &datePickerYearEditor, strconv.Itoa(currentTime.Year())).Layout(gtx)
						}
						return dim
					})
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &datePickerCloseButton, "Cancel").Layout(gtx)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return material.Button(th, &datePickerOkButton, "OK").Layout(gtx)
						}),
					)
				}),
			)
		})
	})
}
