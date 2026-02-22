package start

import (
	"image/color"
	"vpdesktop/types"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

var classSearchButton widget.Clickable
var searchIcon *widget.Icon

func Row(gtx layout.Context, th *material.Theme, state *types.AppState, text string, wrapper *widget.Clickable, col color.NRGBA) layout.FlexChild {
	border := widget.Border{
		Color:        col,
		Width:        1,
		CornerRadius: 10,
	}

	searchIcon, err := widget.NewIcon(icons.ActionSearch)
	if err != nil {
		panic(err)
	}

	return layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Max.Y -= gtx.Dp(70)
		return wrapper.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			pointer.CursorPointer.Add(gtx.Ops)

			return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(10).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											return layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return searchIcon.Layout(gtx, color.NRGBA{R: 255, G: 255, B: 255, A: 120})
											})
										})
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											return layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return material.Body1(th, text).Layout(gtx)
											})
										})
									}),
								)
							})
						})
					})
				}),
			)
		})
	})
}

func DrawStartUI(gtx layout.Context, th material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	if classSearchButton.Clicked(gtx) {
		state.ActiveUI = "class_select"
		gtx.Execute(op.InvalidateCmd{})
	}

	size := gtx.Constraints.Max

	// gradient 😕
	paint.LinearGradientOp{
		Stop1:  f32.Pt(0, 0),
		Color1: color.NRGBA{R: 38, G: 43, B: 51, A: 255},
		Stop2:  f32.Pt(float32(size.X), float32(size.Y)),
		Color2: color.NRGBA{R: 23, G: 26, B: 31, A: 255},
	}.Add(gtx.Ops)

	paint.PaintOp{}.Add(gtx.Ops)

	children := []layout.FlexChild{}

	children = append(children, Row(gtx, &th, state, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "search_prompt"}), &classSearchButton, color.NRGBA{R: 255, G: 255, B: 255, A: 50}))

	return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceBetween}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return Header(gtx, th, state, localizer)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical, Spacing: layout.SpaceAround}.Layout(gtx, children...)
			})
		}),
	)
}
