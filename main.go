package main

import (
	"embed"
	"image/color"
	"log"
	"os"
	"vpdesktop/api"
	"vpdesktop/cache"
	"vpdesktop/types"
	"vpdesktop/ui"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/cloudfoundry/jibber_jabber"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tomlv2 "github.com/pelletier/go-toml/v2"
	"github.com/zalando/go-keyring"
	"golang.org/x/text/language"
)

//go:embed locales/*.toml
var LocaleFS embed.FS
var password string
var localizer *i18n.Localizer
var bundle *i18n.Bundle
var AppState types.AppState

func init() {
	var err error

	AppState = types.AppState{
		ActiveUI: "login",
		Login: types.LoginState{
			LoginPhase:      "school_entry",
			LoginRequested:  false,
			LoginInProgress: false,
			LoginSuccess:    false,
			LoginNote:       "",
		},
		AnimationStates: make(map[string]*types.AnimationState),

		ViewMode: "day",
	}

	if cache.HasCacheFile("recent_logins") {
		credentials, err := cache.ReadJSONCacheFile[types.LoginCredentials]("recent_logins")

		if err == nil {
			AppState.Login.RecentLogin = credentials

			pass, err := keyring.Get("vpdesktop", credentials.School+credentials.Username)
			if err == nil {
				AppState.Login.RecentLogin.Password = pass
				AppState.SelectedSchool = credentials.School
				AppState.SelectedUsername = credentials.Username
				AppState.SelectedPassword = password
			}
		}
	}

	locale, err := jibber_jabber.DetectLanguage()
	if err != nil || locale != "de" && locale != "en" {
		locale = "en"
	}
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", tomlv2.Unmarshal)
	bundle.LoadMessageFileFS(LocaleFS, "locales/"+locale+".toml")

	localizer = i18n.NewLocalizer(bundle, locale)

	cache.EnsureCacheDirExists()

	go func() {
		window := new(app.Window)

		window.Option(app.Title("VPDesktop"))
		window.Option(app.Size(800, 600))

		err := run(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	var ops op.Ops
	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:

			switch AppState.ActiveUI {
			case "login":

				if AppState.Login.RecentLoginDeletionRequested == true {
					keyring.Delete("vpdesktop", AppState.Login.RecentLogin.School+AppState.Login.RecentLogin.Username)
					cache.DeleteCacheFile("recent_logins")

					AppState.Login.RecentLogin = types.LoginCredentials{}
					AppState.Login.RecentLoginDeletionRequested = false

				}

				if AppState.Login.LoginRequested && !AppState.Login.LoginInProgress {
					AppState.Login.LoginRequested = false
					AppState.Login.LoginInProgress = true
					AppState.Login.LoginNote = localizer.MustLocalize(&i18n.LocalizeConfig{
						MessageID: "fetch_plans_progress",
					})

					go func(school, user, pass string) {
						res, err := fetchTimetable(school, user, pass)

						weekDate, err := api.FetchWeeklyClasses(AppState.SelectedSchool, AppState.SelectedUsername, AppState.SelectedPassword, int(AppState.ClassesResponse.Kopf.TageProWoche))

						if err != nil {
							log.Println("Error fetching weekly classes:", err)
						} else {
							AppState.WeekClassesResponse = weekDate
						}

						AppState.ClassesResponse = res

						AppState.Login.LoginInProgress = false
						if err != nil {
							AppState.Login.LoginNote = localizer.MustLocalize(&i18n.LocalizeConfig{
								MessageID: "fetch_plans_error_reason_any",
							})
						} else {
							AppState.Login.LoginNote = ""
						}

						AppState.Login.LoginSuccess = err == nil
						if AppState.Login.LoginSuccess {
							AppState.Login.LoginNote = localizer.MustLocalize(&i18n.LocalizeConfig{
								MessageID: "fetch_plans_success",
							})

							cache.WriteJSONCacheFile("recent_logins",
								types.LoginCredentials{
									School:   school,
									Username: user,
								},
							)

							keyring.Set("vpdesktop", school+user, pass)

							AppState.ActiveUI = "class_select"
						}

						AppState.Login.LoginInProgress = false
						AppState.Login.LoginRequested = false

						AppState.SelectedDate = res.Kopf.DatumPlan
					}(
						AppState.SelectedSchool,
						AppState.SelectedUsername,
						AppState.SelectedPassword,
					)
				}

				// theme
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Fg = color.NRGBA{R: 161, G: 161, B: 161, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.DrawLoginUI(gtx, theme, &AppState, localizer)
				e.Frame(gtx.Ops)

			case "dayview":
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Fg = color.NRGBA{R: 50, G: 50, B: 50, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.DayViewWrapper(gtx, theme, &AppState, localizer)

				if AppState.ShowDatePicker {
					ui.DatePickerOverlayWrapper(gtx, theme, &AppState, localizer)
				}

				e.Frame(gtx.Ops)

			case "weekview":
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Fg = color.NRGBA{R: 50, G: 50, B: 50, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.WeekViewWrapper(gtx, theme, &AppState, localizer)

				if AppState.ShowDatePicker {
					ui.DatePickerOverlayWrapper(gtx, theme, &AppState, localizer)
				}

				e.Frame(gtx.Ops)

			case "start":
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Fg = color.NRGBA{R: 161, G: 161, B: 161, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.StartWrapper(gtx, theme, &AppState, localizer)

				e.Frame(gtx.Ops)

			case "class_select":
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Fg = color.NRGBA{R: 161, G: 161, B: 161, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.ClassSelectWrapper(gtx, theme, &AppState, localizer)

				e.Frame(gtx.Ops)

			case "sample_data":
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Fg = color.NRGBA{R: 161, G: 161, B: 161, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.SampleDataUIWrapper(gtx, theme, &AppState, localizer)

				e.Frame(gtx.Ops)
			}
		}
	}
}

func getLocalizedString(messageID string) string {
	localizedString := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID: "weekkday_1"})
	return localizedString
}
func getLocalizedStrings(messageIDs []string) []string {
	localizedStrings := make([]string, len(messageIDs))
	for i, messageID := range messageIDs {
		localizedStrings[i] = localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: messageID})
	}

	return localizedStrings
}

func fetchTimetable(school, username, password string) (types.ClassesResponse, error) {
	url := api.ComposeURL("stundenplan24.de", types.Classes, school, nil, nil)
	response, err := api.VPMobilClassesRequest(url, username, password)

	if err != nil {
		return types.ClassesResponse{}, err
	}

	return response, nil

}

func main() {}
