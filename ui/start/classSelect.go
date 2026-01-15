package start

import (
	"strings"
	"vpdesktop/types"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

func applyBorder(gtx layout.Context, border widget.Border, w layout.Widget) layout.Dimensions {
	return border.Layout(gtx, w)
}

func vSpace(dp unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Height: dp}.Layout(gtx)
	}
}

var searchEditor widget.Editor
var searchButton widget.Clickable

func ClassSelectUI(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	searchEditor.SingleLine = true

	var classes []string
	for _, v := range state.ClassesResponse.Klassen.Klassen {
		if strings.Contains(v.Kurz, searchEditor.Text()) || searchEditor.Text() == "" {
			classes = append(classes,
				v.Kurz,
			)
		}
	}

	if state.ClassClickables == nil {
		state.ClassClickables = make(map[string]*widget.Clickable)
	}

	visible := make(map[string]struct{}, len(classes))
	for _, c := range classes {
		visible[c] = struct{}{}
	}

	for k := range state.ClassClickables {
		if _, ok := visible[k]; !ok {
			delete(state.ClassClickables, k)
		}
	}

	state.ClassList.Axis = layout.Vertical

	searchIcon, err := widget.NewIcon(icons.HardwareKeyboardArrowRight)
	if err != nil {
		panic(err)
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return applyBorder(gtx, widget.Border{
						Color:        th.Palette.Fg,
						CornerRadius: unit.Dp(8),
						Width:        unit.Dp(1),
					}, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Left:   unit.Dp(6),
							Top:    unit.Dp(10),
							Bottom: unit.Dp(10),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return material.Editor(th, &searchEditor, localizer.MustLocalize(&i18n.LocalizeConfig{
								MessageID: "search_class_hint",
							})).Layout(gtx)
						})
					})

				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return searchButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{
							Top:    unit.Dp(10),
							Bottom: unit.Dp(10),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return searchIcon.Layout(gtx, th.Fg)
						})

					})
				}),
			)
		}),
		layout.Flexed(float32(gtx.Constraints.Max.Y), func(gtx layout.Context) layout.Dimensions {
			return state.ClassList.Layout(gtx, len(classes), func(gtx layout.Context, i int) layout.Dimensions {

				if state.ClassClickables[classes[i]] == nil {
					state.ClassClickables[classes[i]] = new(widget.Clickable)
				}

				if state.ClassClickables[classes[i]].Clicked(gtx) {
					state.SelectedClass = classes[i]
					state.ActiveUI = "dayview"
				}

				return state.ClassClickables[classes[i]].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Body1(th, classes[i]).Layout(gtx)
				})
			})
		}),
	)
}
