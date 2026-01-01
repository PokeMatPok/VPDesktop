package ui

import (
	"strings"
	"vpmobil_app/types"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

func DrawDayViewUI(gtx layout.Context, th *material.Theme, state *types.AppState) layout.Dimensions {

	label := material.H5(th, state.SelectedDate+" - "+state.SelectedClass)

	var match int
	for i, k := range state.ClassesResponse.Klassen.Klassen {
		if strings.Contains(k.Kurz, state.SelectedClass) {
			match = i
			break
		}
	}

	var DisplayValues []string
	for _, m := range state.ClassesResponse.Klassen.Klassen[match].Plan.Stunden {
		DisplayValues = append(DisplayValues, m.Beginn+" - "+m.Ende+" "+m.Fa.Value+" mit "+m.Le.Value)
	}

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{
			Axis: layout.Vertical,
		}.Layout(gtx,
			// Top label
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return label.Layout(gtx)
			}),

			// List of DisplayValues
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx, func() []layout.FlexChild {
					children := make([]layout.FlexChild, 0, len(DisplayValues))
					for _, v := range DisplayValues {
						text := v // avoid loop variable capture
						children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.Body1(th, text).Layout(gtx)
						}))
					}
					return children
				}()...)
			}),
		)
	})

}
