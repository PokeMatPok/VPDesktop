package other

import (
	"vpdesktop/types"
	"vpdesktop/utils"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

var continueButton widget.Clickable

func DrawsampleDataUI(gtx layout.Context, th material.Theme, state *types.AppState, localizer *i18n.Localizer) layout.Dimensions {

	if continueButton.Clicked(gtx) {
		utils.LoadSampleData(state)
		state.ActiveUI = "class_select"
	}

	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Body1(&th, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_title"})).Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Body2(&th, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_body"})).Layout(gtx)
		}),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				return material.Button(&th, &continueButton, localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "sample_data_continue_button"})).Layout(gtx)
			},
		),
	)
}
