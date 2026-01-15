package ui

import (
	"image"
	"image/color"
	"vpdesktop/types"

	"gioui.org/font/gofont"
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

var (
	schoolEditor      widget.Editor
	usernameEditor    widget.Editor
	passwordEditor    widget.Editor
	classEditor       widget.Editor
	RememberLogin     widget.Bool
	statusText        widget.Label
	loginButton       widget.Clickable
	recentLoginButton widget.Clickable
	sampleDataButton  widget.Clickable
)

var recentLoginText richtext.InteractiveText

func init() {
	schoolEditor.SingleLine = true
	usernameEditor.SingleLine = true
	passwordEditor.SingleLine = true
	classEditor.SingleLine = true
	passwordEditor.Mask = '•'
}

func when(cond bool, child layout.FlexChild) layout.FlexChild {
	if cond {
		return child
	}
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{} // empty placeholder
	})
}

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
				Rect: image.Rectangle{Max: gtx.Constraints.Min},
				NE:   gtx.Dp(radius),
				NW:   gtx.Dp(radius),
				SE:   gtx.Dp(radius),
				SW:   gtx.Dp(radius),
			}
			paint.FillShape(gtx.Ops, bg, rr.Op(gtx.Ops))
			return layout.Dimensions{Size: gtx.Constraints.Min}
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

func DrawLoginUI(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	if loginButton.Clicked(gtx) && !state.Login.LoginRequested && !state.Login.LoginInProgress && !state.Login.LoginSuccess && state.SelectedSchool != "" && state.SelectedUsername != "" && state.SelectedPassword != "" {

		state.SelectedSchool = schoolEditor.Text()
		state.SelectedUsername = usernameEditor.Text()
		state.SelectedPassword = passwordEditor.Text()
		state.SelectedClass = classEditor.Text()
		state.Login.RememberLogin = RememberLogin.Value

		state.Login.LoginRequested = true
	}

	if recentLoginButton.Clicked(gtx) {
		state.SelectedUsername = state.Login.RecentLogin.Username
		state.SelectedSchool = state.Login.RecentLogin.School
		state.SelectedPassword = state.Login.RecentLogin.Password

		usernameEditor.SetText(state.SelectedUsername)
		schoolEditor.SetText(state.SelectedSchool)
		passwordEditor.SetText(state.SelectedPassword)
	}

	if sampleDataButton.Clicked(gtx) {
		state.ActiveUI = "sample_data"
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

				title := material.H5(th, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_title"}))
				title.Color = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

				login_as_label := material.Label(th, unit.Sp(12), localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_as"}))
				login_as_label.Color = color.NRGBA{R: 0xfb, G: 0xfb, B: 0xfb, A: 0xff}
				login_as_label.Alignment = text.Middle

				return layout.Flex{
					Axis:    layout.Vertical,
					Spacing: layout.SpaceStart,
				}.Layout(gtx,
					layout.Rigid(title.Layout),

					layout.Rigid(vSpace(12)),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:    layout.Horizontal,
							Spacing: layout.SpaceStart,
						}.Layout(gtx,

							layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{Right: unit.Dp(8)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return applyBorder(gtx, border, func(gtx layout.Context) layout.Dimensions {
										return layout.UniformInset(unit.Dp(4)).Layout(
											gtx,
											material.Editor(th, &schoolEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "schoolnumber"})).Layout,
										)
									})
								})
							}),

							layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
								return applyBorder(gtx, border, func(gtx layout.Context) layout.Dimensions {
									return layout.UniformInset(unit.Dp(4)).Layout(
										gtx,
										material.Editor(th, &usernameEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "username"})).Layout,
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
									return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(4), Right: unit.Dp(8)}.Layout(gtx,
										material.Editor(th, &passwordEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password"})).Layout,
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
									return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(4), Right: unit.Dp(8)}.Layout(gtx,
										material.Editor(th, &classEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "class"})).Layout,
									)
								})
							}),
						)
					}),

					layout.Rigid(vSpace(10)),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.X = gtx.Dp(140)

							chckbx := material.CheckBox(th, &RememberLogin, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "remember_login"}))
							chckbx.IconColor = color.NRGBA{R: 90, G: 140, B: 255, A: 255}
							return chckbx.Layout(gtx)
						})
					}),

					layout.Rigid(vSpace(16)),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

							return loginButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return Card(
									gtx,
									color.NRGBA{R: 90, G: 140, B: 255, A: 255},
									8,
									12,
									func(gtx layout.Context) layout.Dimensions {

										return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,

											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												lbl := material.Body1(
													th,
													localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_btn"}),
												)
												lbl.Color = th.Bg // 👈 THIS is the key line
												return lbl.Layout(gtx)
											}),

											layout.Rigid(func(gtx layout.Context) layout.Dimensions {
												icon, err := widget.NewIcon(icons.HardwareKeyboardArrowRight)
												if err != nil {
													panic(err)
												}
												return icon.Layout(gtx, th.Bg) // match text
											}),
										)
									},
								)
							})
						})
					}),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.X = gtx.Dp(140)

							label := material.Label(th, unit.Sp(gtx.Dp(14)), state.Login.LoginNote)
							return label.Layout(gtx)
						})
					}),

					when(state.Login.RecentLogin.Username != "", layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

							gtx.Constraints.Min.X = min(gtx.Constraints.Max.X, gtx.Dp(220))

							return recentLoginButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

								// Background
								rr := clip.RRect{
									Rect: image.Rectangle{Max: gtx.Constraints.Max},
									NE:   gtx.Dp(10), NW: gtx.Dp(10),
									SE: gtx.Dp(10), SW: gtx.Dp(10),
								}
								paint.FillShape(
									gtx.Ops,
									color.NRGBA{R: 90, G: 90, B: 100, A: 255},
									rr.Op(gtx.Ops),
								)

								return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {

									spans := []richtext.SpanStyle{
										{
											Content: state.Login.RecentLogin.School + "\n",
											Color:   color.NRGBA{R: 160, G: 160, B: 160, A: 255},
											Size:    unit.Sp(13),
											Font:    gofont.Collection()[0].Font,
										},
										{
											Content: localizer.MustLocalize(&i18n.LocalizeConfig{
												MessageID: "login_as",
											}) + " " + state.Login.RecentLogin.Username,
											Color:       color.NRGBA{R: 255, G: 255, B: 255, A: 255},
											Size:        unit.Sp(15),
											Font:        gofont.Collection()[0].Font,
											Interactive: true,
										},
										{
											Content:     " (" + localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "delete_login"}) + ")",
											Color:       color.NRGBA{R: 255, G: 0, B: 0, A: 255},
											Size:        unit.Sp(12),
											Font:        gofont.Collection()[0].Font,
											Interactive: true,
										},
									}

									// Process richtext interaction events
									for {
										span, ev, ok := recentLoginText.Update(gtx)
										if !ok {
											break
										}
										if ev.Type == richtext.Click {
											if v, _ := span.Content(); v == " ("+localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "delete_login"})+")" {
												state.Login.RecentLoginDeletionRequested = true
											} else {
												recentLoginButton.Click() // forward to button logic
											}
										}
									}

									return richtext.Text(
										&recentLoginText,
										th.Shaper,
										spans...,
									).Layout(gtx)
								})
							})
						})
					})),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return sampleDataButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Body2(th,
									localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_button"}),
								)
								lbl.Color = color.NRGBA{R: 150, G: 150, B: 160, A: 255}
								return layout.UniformInset(unit.Dp(8)).Layout(gtx, lbl.Layout)
							})
						})
					}),
				)
			},
		)
	})
}
