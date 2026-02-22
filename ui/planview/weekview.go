package planview

import (
	"fmt"
	"image"
	"image/color"
	"vpdesktop/types"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func TableCell(gtx layout.Context, th *material.Theme, isHeader bool, content []string) layout.FlexChild {
	headBorder := widget.Border{
		Color: color.NRGBA{R: 150, G: 150, B: 240, A: 255},
		Width: unit.Dp(2),
	}
	bodyBorder := widget.Border{
		Color: color.NRGBA{R: 150, G: 150, B: 150, A: 255},
		Width: unit.Dp(2),
	}

	border := bodyBorder
	if isHeader {
		border = headBorder
	}

	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min = image.Pt(75, 75)
		gtx.Constraints.Max = image.Pt(75, 75)
		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min = image.Pt(75, 75)
			gtx.Constraints.Max = image.Pt(75, 75)

			children := make([]layout.FlexChild, len(content))
			for i, line := range content {
				children[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, line).Layout(gtx)
				})
			}
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween, Alignment: layout.Middle}.Layout(gtx, children...)
			})
		})
	})
}

func DrawWeekViewUI(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		var list layout.List
		list.Axis = layout.Horizontal

		return list.Layout(gtx, len(state.WeekViewState.Days)+1, func(gtx layout.Context, dayIndex int) layout.Dimensions {
			gtx.Constraints.Max.X = min(gtx.Constraints.Max.X, gtx.Dp(600))

			var displayData []layout.FlexChild

			if dayIndex == 0 {

				day := state.WeekViewState.Days[0]
				for _, d := range state.WeekViewState.Days {
					if len(d.Lessons) > len(day.Lessons) {
						day = d
					}
				}

				displayData = append(displayData, TableCell(gtx, th, true, []string{
					"",
				}))

				for _, lesson := range day.Lessons {

					displayData = append(displayData, TableCell(gtx, th, true, []string{
						lesson.Beginn,
						"-",
						lesson.Ende,
					}))
				}
			} else {
				day := state.WeekViewState.Days[dayIndex-1]

				displayData = append(displayData, TableCell(gtx, th, true, []string{
					localizer.MustLocalize(&i18n.LocalizeConfig{
						MessageID: fmt.Sprintf("weekday_%d", dayIndex),
					}),
				}))

				for _, lesson := range day.Lessons {

					displayData = append(displayData,
						TableCell(gtx, th, false, []string{
							lesson.Fa.Value,
							lesson.Le.Value,
						}))
				}
			}

			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, displayData...)
		})
	})
}
