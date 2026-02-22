package ui

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"vpdesktop/types"
	"vpdesktop/ui/components"

	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
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
	schoolEditor            widget.Editor
	usernameEditor          widget.Editor
	passwordEditor          widget.Editor
	RememberLogin           widget.Bool
	statusText              widget.Label
	loginButton             widget.Clickable
	loginButtonBottomBorder components.ReactiveBottomBorder
	recentLoginButton       widget.Clickable
	sampleDataButton        widget.Clickable
	passBottomBorder        components.ReactiveBottomBorder
	userBottomBorder        components.ReactiveBottomBorder
	schoolBottomBorder      components.ReactiveBottomBorder
)

//go:embed assets/login_atmos.jpg
var backgroundBytes []byte
var backgroundImage *widget.Image

var recentLoginText richtext.InteractiveText

func init() {
	schoolEditor.SingleLine = true
	usernameEditor.SingleLine = true
	passwordEditor.SingleLine = true
	passwordEditor.Mask = '•'

	bottomBorderConfig := components.ReactiveBottomBorder{
		Height:     1,
		Color:      color.NRGBA{R: 180, G: 180, B: 180, A: 255},
		FocusColor: color.NRGBA{R: 90, G: 140, B: 255, A: 255},
		SpeedIn:    2,
		SpeedOut:   4,
	}

	passBottomBorder = bottomBorderConfig
	userBottomBorder = bottomBorderConfig
	schoolBottomBorder = bottomBorderConfig
	loginButtonBottomBorder = bottomBorderConfig

	image, err := jpeg.Decode(bytes.NewReader(backgroundBytes))
	if err != nil {
		panic("failed to decode embedded background image: " + err.Error())
	}

	backgroundImage = &widget.Image{
		Src: paint.NewImageOp(image),
		Fit: widget.Cover, // fills space, crops if needed
	}
}

func when(cond bool, child layout.FlexChild) layout.FlexChild {
	if cond {
		return child
	}
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{} // empty placeholder
	})
}

func vSpace(dp unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Spacer{Height: dp}.Layout(gtx)
	}
}

func DrawLoginUI(gtx layout.Context, th *material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	if loginButton.Clicked(gtx) && !state.Login.LoginRequested && !state.Login.LoginInProgress {

		switch state.Login.LoginPhase {
		case "school_entry":
			if schoolEditor.Text() != "" {
				state.SelectedSchool = schoolEditor.Text()
				state.Login.LoginPhase = "user_selection"
				gtx.Execute(op.InvalidateCmd{})
			}
		case "user_selection":
			if usernameEditor.Text() != "" {
				state.SelectedUsername = usernameEditor.Text()
				state.Login.LoginPhase = "password_entry"
				gtx.Execute(op.InvalidateCmd{})
			}
		case "password_entry":
			if passwordEditor.Text() != "" {
				state.SelectedPassword = passwordEditor.Text()
				state.Login.LoginRequested = true
				gtx.Execute(op.InvalidateCmd{})
			}
		}
	}

	if recentLoginButton.Clicked(gtx) {
		state.SelectedUsername = state.Login.RecentLogin.Username
		state.SelectedSchool = state.Login.RecentLogin.School
		state.SelectedPassword = state.Login.RecentLogin.Password

		usernameEditor.SetText(state.SelectedUsername)
		schoolEditor.SetText(state.SelectedSchool)
		passwordEditor.SetText(state.SelectedPassword)

		state.Login.LoginRequested = true
		gtx.Execute(op.InvalidateCmd{})
	}

	if sampleDataButton.Clicked(gtx) {
		state.ActiveUI = "sample_data"
	}

	titleTheme := material.Theme{
		Shaper:   th.Shaper,
		TextSize: 28,
		Face:     "Times New Roman",
	}
	titleTheme.Fg = th.Fg

	title := material.H5(&titleTheme, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_title"}))
	title.Color = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}

	login_as_label := material.Label(th, unit.Sp(12), localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_as"}))
	login_as_label.Color = color.NRGBA{R: 0xfb, G: 0xfb, B: 0xfb, A: 0xff}
	login_as_label.Alignment = text.Middle

	var UIweight float32 = 1.0
	if gtx.Constraints.Max.X > gtx.Sp(600) {
		UIweight = 0.55
	}

	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
		WeightSum: 1.0,
	}.Layout(gtx,

		layout.Flexed(UIweight, func(gtx layout.Context) layout.Dimensions {

			size := gtx.Constraints.Max

			//gradient 🤔
			paint.LinearGradientOp{
				Stop1:  f32.Pt(0, 0),
				Color1: color.NRGBA{R: 38, G: 43, B: 51, A: 255},
				Stop2:  f32.Pt(float32(size.X), float32(size.Y)),
				Color2: color.NRGBA{R: 23, G: 26, B: 31, A: 255},
			}.Add(gtx.Ops)

			paint.PaintOp{}.Add(gtx.Ops)

			return layout.Flex{
				Axis:      layout.Vertical,
				Alignment: layout.Middle,
				Spacing:   layout.SpaceBetween,
				WeightSum: 1,
			}.Layout(gtx,

				layout.Flexed(0.8, func(gtx layout.Context) layout.Dimensions {

					border := widget.Border{
						Color:        color.NRGBA{R: 255, G: 255, B: 255, A: 50},
						Width:        1,
						CornerRadius: unit.Dp(10),
					}

					return layout.Inset{Top: unit.Dp(30), Bottom: unit.Dp(10), Left: unit.Dp(20), Right: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

							paint.FillShape(
								gtx.Ops,
								color.NRGBA{R: 48, G: 54, B: 64, A: 255},
								clip.RRect{
									Rect: image.Rectangle{Max: gtx.Constraints.Max},
									NE:   gtx.Dp(10), NW: gtx.Dp(10),
									SE: gtx.Dp(10), SW: gtx.Dp(10),
								}.Op(gtx.Ops),
							)

							return layout.Inset{Top: unit.Dp(20), Bottom: unit.Dp(20), Left: unit.Dp(20), Right: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{
									Axis:      layout.Vertical,
									Spacing:   layout.SpaceBetween,
									Alignment: layout.End,
								}.Layout(gtx,

									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Inset{Top: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return title.Layout(gtx)
											})
										})
									}),

									layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
										return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											return layout.Inset{Right: unit.Dp(8)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

												switch state.Login.LoginPhase {
												case "school_entry":

													return schoolBottomBorder.Layout(gtx, gtx.Source.Focused(&schoolEditor), func(gtx layout.Context) layout.Dimensions {

														gtx.Constraints.Min.X = min(gtx.Constraints.Max.X-gtx.Dp(20), gtx.Dp(300))

														return layout.UniformInset(unit.Dp(4)).Layout(
															gtx,
															material.Editor(th, &schoolEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "schoolnumber"})).Layout,
														)
													})

												case "user_selection":
													return userBottomBorder.Layout(gtx, gtx.Source.Focused(&usernameEditor), func(gtx layout.Context) layout.Dimensions {

														gtx.Constraints.Min.X = min(gtx.Constraints.Max.X-gtx.Dp(20), gtx.Dp(300))

														return layout.UniformInset(unit.Dp(4)).Layout(
															gtx,
															material.Editor(th, &usernameEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "username"})).Layout,
														)
													})
												case "password_entry":
													return layout.Flex{
														Axis:    layout.Vertical,
														Spacing: layout.SpaceEnd,
													}.Layout(gtx,
														layout.Rigid(func(gtx layout.Context) layout.Dimensions {
															return passBottomBorder.Layout(gtx, gtx.Source.Focused(&passwordEditor), func(gtx layout.Context) layout.Dimensions {

																gtx.Constraints.Min.X = min(gtx.Constraints.Max.X-gtx.Dp(20), gtx.Dp(300))

																return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(4), Right: unit.Dp(8)}.Layout(gtx,
																	material.Editor(th, &passwordEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password"})).Layout,
																)
															})
														}),

														layout.Rigid(func(gtx layout.Context) layout.Dimensions {
															return layout.Inset{Top: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
																return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
																	gtx.Constraints.Min.X = gtx.Dp(140)

																	chckbx := material.CheckBox(th, &RememberLogin, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "remember_login"}))
																	chckbx.IconColor = color.NRGBA{R: 90, G: 140, B: 255, A: 255}
																	return chckbx.Layout(gtx)
																})
															})
														}),
													)

												}

												return layout.Dimensions{Size: gtx.Constraints.Min} // placeholder for other phases
											})
										})
									}),
									/*layout.Rigid(func(gtx layout.Context) layout.Dimensions {

										return layout.Flex{
											Axis:    layout.Horizontal,
											Spacing: layout.SpaceStart,
										}.Layout(gtx,



											layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
												return userBottomBorder.Layout(gtx, gtx.Source.Focused(&usernameEditor), func(gtx layout.Context) layout.Dimensions {
													return layout.UniformInset(unit.Dp(4)).Layout(
														gtx,
														material.Editor(th, &usernameEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "username"})).Layout,
													)
												})
											}),
										)
									}),

									layout.Rigid(func(gtx layout.Context) layout.Dimensions {

										passInAnimState := state.AnimationStates["login_password"]

										if passInAnimState == nil {
											state.AnimationStates["login_password"] = &types.AnimationState{
												Progress: 0,
											}
											passInAnimState = state.AnimationStates["login_password"]
										}

										return layout.Flex{
											Axis:    layout.Horizontal,
											Spacing: layout.SpaceBetween,
										}.Layout(gtx,

											layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
												return passBottomBorder.Layout(gtx, gtx.Source.Focused(&passwordEditor), func(gtx layout.Context) layout.Dimensions {
													// It is recommended to add a small Inset so the text doesn't touch the border
													return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4), Left: unit.Dp(4), Right: unit.Dp(8)}.Layout(gtx,
														material.Editor(th, &passwordEditor, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password"})).Layout,
													)
												})
											}),
										)
									}),*/

									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Flex{
											Axis:    layout.Horizontal,
											Spacing: layout.SpaceStart, // pushes button to the right
										}.Layout(gtx,

											layout.Rigid(func(gtx layout.Context) layout.Dimensions {

												return layout.Inset{Top: unit.Dp(10), Bottom: unit.Dp(10), Right: unit.Dp(20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

													border.Color = color.NRGBA{R: 90, G: 140, B: 255, A: 255}

													return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

														return layout.UniformInset(10).Layout(gtx, func(gtx layout.Context) layout.Dimensions {

															return loginButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

																return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,

																	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
																		lbl := material.Body1(
																			th,
																			localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_btn"}),
																		)
																		lbl.Color = th.Fg
																		return lbl.Layout(gtx)
																	}),

																	layout.Rigid(func(gtx layout.Context) layout.Dimensions {
																		icon, err := widget.NewIcon(icons.HardwareKeyboardArrowRight)
																		if err != nil {
																			panic(err)
																		}
																		return icon.Layout(gtx, th.Fg) // match text
																	}),
																)
															})
														})
													})
												})
											}))
									}),

									when(state.Login.RecentLogin.Username != "", layout.Flexed(0.2, func(gtx layout.Context) layout.Dimensions {
										return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											return material.Body1(th, localizer.MustLocalize(&i18n.LocalizeConfig{
												MessageID: "login_as",
											})).Layout(gtx)
										})

									})),

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
															Content:     state.Login.RecentLogin.Username,
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
								)
							})
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

				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Bottom: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return sampleDataButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								lbl := material.Body2(th,
									localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_button"}),
								)
								lbl.Color = color.NRGBA{R: 150, G: 150, B: 160, A: 255}
								return layout.UniformInset(unit.Dp(8)).Layout(gtx, lbl.Layout)
							})
						})
					})
				}),
			)
		},
		),
		when(UIweight != 1.0, layout.Flexed(0.45, func(gtx layout.Context) layout.Dimensions {
			size := gtx.Constraints.Max // fill available Flexed space
			paint.FillShape(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 0, A: 255}, clip.Rect{Max: size}.Op())
			return backgroundImage.Layout(gtx)
		})),
	)
}
