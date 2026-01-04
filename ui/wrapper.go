package ui

import (
	"vpdesktop/types"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

func Wrapper(gtx layout.Context, th *material.Theme, state *types.AppState) layout.Dimensions {

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return Header(gtx, th, state)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return DrawDayViewUI(gtx, th, state)
		}),
	)
}
