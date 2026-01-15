package ui

import (
	"image/color"
	"strings"
	"vpdesktop/types"
	"vpdesktop/ui/dayview"
	"vpdesktop/ui/other"
	"vpdesktop/ui/start"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func DayViewWrapper(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	for _, k := range state.ClassesResponse.Klassen.Klassen {
		if strings.Contains(k.Kurz, state.SelectedClass) {
			state.DayViewState.Lessons = []types.LessonDisplayData{}
			for _, l := range k.Plan.Stunden {
				lesson := types.LessonDisplayData{
					Beginn: l.Beginn,
					Ende:   l.Ende,
					Fa:     types.ValueWithNote{Value: l.Fa.Value, Note: l.Fa.FaAe},
					Le:     types.ValueWithNote{Value: l.Le.Value, Note: l.Le.LeAe},
				}
				state.DayViewState.Lessons = append(state.DayViewState.Lessons, lesson)
			}
		}
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return dayview.Header(gtx, th, state, localizer)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return dayview.DrawDayViewUI(gtx, th, state)
		}),
	)
}

func StartWrapper(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	size := gtx.Constraints.Max

	// Paint background FIRST
	paint.LinearGradientOp{
		Stop1:  f32.Point{X: 0, Y: 0},
		Color1: color.NRGBA{R: 255, G: 0, B: 0, A: 255},
		Stop2:  f32.Point{X: float32(size.X), Y: float32(size.Y)},
		Color2: color.NRGBA{R: 0, G: 0, B: 255, A: 255},
	}.Add(gtx.Ops)

	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Placeholder for Start Header
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return start.DrawStartUI(gtx, *th, state, localizer)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// Placeholder for Start Footer
			return layout.Dimensions{Size: gtx.Constraints.Min}
		}),
	)
}

func SampleDataUIWrapper(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	return other.DrawsampleDataUI(gtx, *th, state, localizer)
}

func ClassSelectWrapper(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	return start.ClassSelectUI(gtx, th, state, localizer)
}
