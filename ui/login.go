package ui

import (
	"image"
	"image/color"
	"vpmobil_app/types"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var (
	schoolEditor   widget.Editor
	usernameEditor widget.Editor
	passwordEditor widget.Editor
	classEditor    widget.Editor
	statusText     widget.Label
	loginButton    widget.Clickable
)

func Card(
	gtx layout.Context,
	bg color.NRGBA,
	radius unit.Dp,
	inset unit.Dp,
	w layout.Widget,
) layout.Dimensions {

	return layout.Stack{}.Layout(gtx,
		// Background
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rr := clip.RRect{
				Rect: image.Rectangle{Max: gtx.Constraints.Max.Sub(image.Point{X: 20, Y: 20})},
				NE:   gtx.Dp(radius),
				NW:   gtx.Dp(radius),
				SE:   gtx.Dp(radius),
				SW:   gtx.Dp(radius),
			}
			paint.FillShape(gtx.Ops, bg, rr.Op(gtx.Ops))
			return layout.Dimensions{Size: gtx.Constraints.Max}
		}),

		// Content
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(inset).Layout(gtx, w)
		}),
	)
}

func applyBorder(gtx layout.Context, border widget.Border, w layout.Widget) layout.Dimensions {
	return border.Layout(gtx, w)
}

func inputBox(
	gtx layout.Context,
	th *material.Theme,
	ed *widget.Editor,
	hint string,
	border widget.Border,
) layout.Dimensions {

	bg := color.NRGBA{R: 40, G: 40, B: 50, A: 255}

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rr := clip.RRect{
				Rect: image.Rectangle{Max: gtx.Constraints.Max},
				NE:   gtx.Dp(8), NW: gtx.Dp(8),
				SE: gtx.Dp(8), SW: gtx.Dp(8),
			}
			paint.FillShape(gtx.Ops, bg, rr.Op(gtx.Ops))
			return layout.Dimensions{Size: gtx.Constraints.Max}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return applyBorder(gtx, border, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(8)).Layout(
					gtx,
					material.Editor(th, ed, hint).Layout,
				)
			})
		}),
	)
}

func vSpace(dp unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Height: dp}.Layout(gtx)
	}
}

func DrawLoginUI(gtx layout.Context, th *material.Theme, state *types.AppState, locale map[string]string) layout.Dimensions {

	if loginButton.Clicked(gtx) && !state.LoginRequested && !state.LoginInProgress && !state.LoginSuccess && state.SelectedSchool != "" && state.SelectedUsername != "" && state.SelectedPassword != "" {
		state.LoginRequested = true
	}

	return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

		gtx.Constraints.Max.X = gtx.Dp(360)

		return Card(
			gtx,
			th.Bg, // background color
			16,    // corner radius
			20,    // padding
			func(gtx layout.Context) layout.Dimensions {

				border := widget.Border{
					Color:        color.NRGBA{R: 180, G: 180, B: 180, A: 255},
					Width:        unit.Dp(1),
					CornerRadius: unit.Dp(8),
				}

				schoolEditor.SingleLine = true
				usernameEditor.SingleLine = true
				passwordEditor.SingleLine = true
				classEditor.SingleLine = true
				passwordEditor.Mask = '•'

				state.SelectedSchool = schoolEditor.Text()
				state.SelectedUsername = usernameEditor.Text()
				state.SelectedPassword = passwordEditor.Text()
				state.SelectedClass = classEditor.Text()

				title := material.H5(th, locale["title"])
				title.Color = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

				return layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceEvenly,
				}.Layout(gtx,
					layout.Rigid(title.Layout),

					layout.Rigid(vSpace(12)),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,

							layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
								return applyBorder(gtx, border, func(gtx layout.Context) layout.Dimensions {
									return layout.UniformInset(unit.Dp(4)).Layout(
										gtx,
										material.Editor(th, &schoolEditor, locale["schoolnumber"]).Layout,
									)
								})
							}),

							layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
								return applyBorder(gtx, border, func(gtx layout.Context) layout.Dimensions {
									return layout.UniformInset(unit.Dp(4)).Layout(
										gtx,
										material.Editor(th, &usernameEditor, locale["username"]).Layout,
									)
								})
							}),
						)
					}),

					layout.Rigid(vSpace(10)),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,

							layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
								return applyBorder(gtx, border, func(gtx layout.Context) layout.Dimensions {
									// It is recommended to add a small Inset so the text doesn't touch the border
									return layout.UniformInset(unit.Dp(4)).Layout(gtx,
										material.Editor(th, &passwordEditor, locale["password"]).Layout,
									)
								})
							}),
						)
					}),

					layout.Rigid(vSpace(12)),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceBetween,
						}.Layout(gtx,

							layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
								return applyBorder(gtx, border, func(gtx layout.Context) layout.Dimensions {
									// It is recommended to add a small Inset so the text doesn't touch the border
									return layout.UniformInset(unit.Dp(4)).Layout(gtx,
										material.Editor(th, &classEditor, locale["class"]).Layout,
									)
								})
							}),
						)
					}),

					layout.Rigid(vSpace(16)),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.X = gtx.Dp(140)

							btn := material.Button(th, &loginButton, locale["login_btn"])
							btn.Background = color.NRGBA{R: 90, G: 140, B: 255, A: 255}
							return btn.Layout(gtx)
						})
					}),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.X = gtx.Dp(140)

							label := material.Label(th, unit.Sp(gtx.Dp(14)), state.LoginNote)
							return label.Layout(gtx)
						})
					}),
				)
			},
		)
	})
}
