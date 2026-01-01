package main

import (
	"embed"
	"image/color"
	"log"
	"os"
	"vpmobil_app/api"
	"vpmobil_app/types"
	"vpmobil_app/ui"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget/material"
	"github.com/cloudfoundry/jibber_jabber"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	tomlv2 "github.com/pelletier/go-toml/v2"
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

		LoginRequested:  false,
		LoginInProgress: false,
		LoginSuccess:    false,
		LoginNote:       "",
	}

	locale, err := jibber_jabber.DetectLanguage()
	if err != nil || locale != "de" && locale != "en" {
		locale = "en"
	}
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", tomlv2.Unmarshal)
	bundle.LoadMessageFileFS(LocaleFS, "locales/"+locale+".toml")

	localizer = i18n.NewLocalizer(bundle, locale)

	// still in development
	/*cache.EnsureCacheDirExists()
	cache.EnsureSchoolCacheDir("example")
	cache.WriteCacheFile("example/example", []byte("This is a test cache file."))
	cacheData, err := cache.ReadCacheFile("example/example")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(cacheData))

	err := keyring.Set("vpmobil", "username", "password")
	if err != nil {
		panic(err)
	}

	password, err = keyring.Get("vpmobil_app", "lehrer")
	if err != nil {
		panic(err)
	}*/

	go func() {
		app.Title("VPMobil")
		app.Size(200, 450)
		window := new(app.Window)

		window.Option(app.Title("VPMobil"))
		window.Option(app.Size(400, 600))

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

				if AppState.LoginRequested && !AppState.LoginInProgress {
					AppState.LoginRequested = false
					AppState.LoginInProgress = true
					AppState.LoginNote = localizer.MustLocalize(&i18n.LocalizeConfig{
						MessageID: "fetch_plans_progress",
					})

					go func(school, user, pass string) {
						res, err := fetchTimetable(school, user, pass)

						AppState.ClassesResponse = res

						AppState.LoginInProgress = false
						if err != nil {
							AppState.LoginNote = localizer.MustLocalize(&i18n.LocalizeConfig{
								MessageID: "fetch_plans_error_reason_any",
							})
						} else {
							AppState.LoginNote = ""
						}

						AppState.LoginSuccess = err == nil
						if AppState.LoginSuccess {
							AppState.LoginNote = localizer.MustLocalize(&i18n.LocalizeConfig{
								MessageID: "fetch_plans_success",
							})

							AppState.ActiveUI = "dayview"
						}

						AppState.LoginInProgress = false
						AppState.LoginRequested = false

						AppState.SelectedDate = res.Kopf.DatumPlan
					}(
						AppState.SelectedSchool,
						AppState.SelectedUsername,
						AppState.SelectedPassword,
					)
				}

				// This graphics context is used for managing the rendering state.
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Fg = color.NRGBA{R: 161, G: 161, B: 161, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.DrawLoginUI(gtx, theme, &AppState,
					map[string]string{
						"title":                            localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_title"}),
						"schoolnumber":                     localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "schoolnumber"}),
						"username":                         localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "username"}),
						"password":                         localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "password"}),
						"class":                            localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "class"}),
						"login_btn":                        localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "login_btn"}),
						"fetch_plans_progress":             localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "fetch_plans_progress"}),
						"fetch_plans_success":              localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "fetch_plans_success"}),
						"fetch_plans_error_reason_any":     localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "fetch_plans_error_reason_any"}),
						"fetch_plans_error_reason_auth":    localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "fetch_plans_error_reason_auth"}),
						"fetch_plans_error_reason_network": localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "fetch_plans_error_reason_network"}),
					})
				// Pass the drawing operations to the GPU.
				e.Frame(gtx.Ops)

			case "dayview":
				gtx := app.NewContext(&ops, e)
				theme := material.NewTheme()
				theme.Palette.ContrastBg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				theme.Palette.ContrastFg = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
				theme.Fg = color.NRGBA{R: 161, G: 161, B: 161, A: 255}
				theme.Bg = color.NRGBA{R: 30, G: 30, B: 30, A: 255}

				ui.DrawDayViewUI(gtx, theme, &AppState)

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

// still in development
/*func main() {
	var UITableData types.TableViewData

	UITableData = types.TableViewData{

		ValuesPerRow: int(response.Kopf.TageProWoche),
		HorHeader:    getLocalizedStrings([]string{"weekday_1", "weekday_2", "weekday_3", "weekday_4", "weekday_5", "weekday_6", "weekday_7"}[0:response.Kopf.TageProWoche]),
		VerHeader: getLocalizedStrings([]string{
			"class_1", "class_2", "class_3", "class_4", "class_5",
			"class_6", "class_7", "class_8", "class_9"}),
	}

	fmt.Print(UITableData.HorHeader)
}*/
