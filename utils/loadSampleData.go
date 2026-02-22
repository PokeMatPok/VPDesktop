package utils

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"vpdesktop/types"
)

var subjects = []string{"MA", "EN", "DE", "PH", "BI", "CH", "GE", "GK", "ET", "SP", "KU", "MU", "IT", "RE", "FR", "LA", "WI"}
var teachers = []string{"ME", "SM", "WE", "KR", "HN", "AL", "MU", "BA", "KL", "GR", "BR", "SC", "WI", "ZI", "VO", "PF"}
var rooms = []string{"R101", "R102", "R103", "R104", "R105", "R110", "R111", "R201", "R202", "R203", "GYM", "LAB1", "LAB2", "ART", "MUS", "LIB"}
var changeTypes = []string{"normal", "normal", "normal", "normal", "vertretung", "entfall", "raumänderung", "verlegt"}
var classNames = []string{"5A", "5B", "5C", "6A", "6B", "7A", "7B", "8A", "8B", "8C", "9A", "9B", "10A", "10B", "11", "12"}

var periodTimes = map[int8][2]string{
	1: {"08:00", "08:45"},
	2: {"08:45", "09:30"},
	3: {"09:45", "10:30"},
	4: {"10:30", "11:15"},
	5: {"11:30", "12:15"},
	6: {"12:15", "13:00"},
	7: {"13:45", "14:30"},
	8: {"14:30", "15:15"},
	9: {"15:30", "16:15"},
}

func pick(pool []string) string { return pool[rand.Intn(len(pool))] }
func pickN(pool []string, n int) []string {
	cp := make([]string, len(pool))
	copy(cp, pool)
	rand.Shuffle(len(cp), func(i, j int) { cp[i], cp[j] = cp[j], cp[i] })
	if n > len(cp) {
		n = len(cp)
	}
	return cp[:n]
}

func randomLesson(period int8) types.Stunde {
	times := periodTimes[period]
	ae := ""
	info := pick(changeTypes)
	if info != "normal" {
		ae = "1"
	}
	fa := types.Fach{Value: pick(subjects), FaAe: ae}
	le := types.Lehrer{Value: pick(teachers), LeAe: ae}
	ra := types.Raum{Value: pick(rooms), RaAe: ae}
	s := types.Stunde{
		St:     period,
		Beginn: times[0],
		Ende:   times[1],
		Fa:     fa,
		Le:     le,
		Ra:     ra,
		If:     info,
	}
	if rand.Intn(5) == 0 {
		v := int16(rand.Intn(200) + 1)
		s.Nr = &v
	}
	return s
}

func randomLessons(minP, maxP int8) []types.Stunde {
	count := int(minP) + rand.Intn(int(maxP-minP)+1)
	var lessons []types.Stunde
	used := map[int8]bool{}
	for p := minP; p <= int8(count) && p <= maxP; p++ {
		if !used[p] {
			lessons = append(lessons, randomLesson(p))
			used[p] = true
		}
	}
	return lessons
}

func fullPeriodSlots() types.KlStunden {
	slots := make([]types.KlStunde, 0, len(periodTimes))
	for p := int8(1); p <= 9; p++ {
		t := periodTimes[p]
		slots = append(slots, types.KlStunde{Value: p, ZeitVon: t[0], ZeitBis: t[1]})
	}
	return types.KlStunden{Stunden: slots}
}

func randomCourses(subjectPool []string) types.Kurse {
	n := rand.Intn(4) + 1
	selected := pickN(subjectPool, n)
	courses := make([]types.Kurs, len(selected))
	for i, s := range selected {
		courses[i] = types.Kurs{KKz: types.KKz{Value: s, KLe: pick(teachers)}}
	}
	return types.Kurse{Kurse: courses}
}

func randomUnits() types.Unterricht {
	n := rand.Intn(5) + 1
	units := make([]types.Unterrichtseinheit, n)
	for i := range units {
		units[i] = types.Unterrichtseinheit{
			Nummer: types.UeNr{
				Value: int16(rand.Intn(300) + 1),
				UeLe:  pick(teachers),
				UeFa:  pick(subjects),
				UeGr:  fmt.Sprintf("G%d", rand.Intn(3)+1),
			},
		}
	}
	return types.Unterricht{Einheiten: units}
}

func randomClass(name string) types.Klasse {
	maxPeriod := int8(5 + rand.Intn(5))
	return types.Klasse{
		Kurz:       name,
		Hash:       fmt.Sprintf("hash-%s-%04x", name, rand.Intn(0xFFFF)),
		KlStunden:  fullPeriodSlots(),
		Kurse:      randomCourses(subjects),
		Unterricht: randomUnits(),
		Plan:       types.Plan{Stunden: randomLessons(1, maxPeriod)},
	}
}

func addDays(year, month, day, n int) (int, int, int) {
	daysInMonth := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	day += n
	for day > daysInMonth[month] {
		day -= daysInMonth[month]
		month++
		if month > 12 {
			month = 1
			year++
		}
	}
	return year, month, day
}

func formatDate(year, month, day int) string {
	return fmt.Sprintf("%04d-%02d-%02d", year, month, day)
}

func randomSchoolDate() (dateStr, tsStr string) {
	weekOffset := rand.Intn(40)
	totalDays := weekOffset * 7
	year, month, day := addDays(2025, 9, 1, totalDays)
	dateStr = formatDate(year, month, day)
	tsStr = fmt.Sprintf("%sT%02d:%02d:00", dateStr, rand.Intn(3)+6, rand.Intn(60))
	return
}

func randomDayResponse(year, month, day int, schoolNum string, weekNum int8, selectedClasses []string) types.ClassesResponse {
	date := formatDate(year, month, day)
	ts := fmt.Sprintf("%sT%02d:%02d:00", date, rand.Intn(3)+6, rand.Intn(60))

	classes := make([]types.Klasse, len(selectedClasses))
	for i, name := range selectedClasses {
		classes[i] = randomClass(name)
	}

	numFreeDays := rand.Intn(3)
	freeDays := make([]int, numFreeDays)
	for i := range freeDays {
		freeDays[i] = rand.Intn(365) + 1
	}

	infoLines := []string{
		fmt.Sprintf("Stand: %s", date),
		fmt.Sprintf("Week %d", weekNum),
		"No lessons in marked periods",
		"Room changes apply",
		"Sample data only",
	}
	numLines := rand.Intn(3) + 1
	rand.Shuffle(len(infoLines), func(i, j int) { infoLines[i], infoLines[j] = infoLines[j], infoLines[i] })

	return types.ClassesResponse{
		XMLName: xml.Name{Local: "VpMobil"},
		Kopf: types.Kopf{
			Planart:      "KL",
			Zeitstempel:  ts,
			DatumPlan:    date,
			Datei:        fmt.Sprintf("klassen_%s.xml", date),
			Nativ:        1,
			Woche:        weekNum,
			TageProWoche: 5,
			Schulnummer:  schoolNum,
		},
		FreieTage:  types.FreieTage{FT: freeDays},
		Klassen:    types.Klassen{Klassen: classes},
		ZusatzInfo: types.ZusatzInfo{Zeilen: infoLines[:numLines]},
	}
}

func LoadSampleData(state *types.AppState) {
	mondayDate, _ := randomSchoolDate()
	weekNum := int8(rand.Intn(52) + 1)
	schoolNum := fmt.Sprintf("%06d", rand.Intn(900000)+100000)

	numClasses := rand.Intn(5) + 4
	selectedClasses := pickN(classNames, numClasses)

	var mondayYear, mondayMonth, mondayDay int
	fmt.Sscanf(mondayDate, "%d-%d-%d", &mondayYear, &mondayMonth, &mondayDay)

	weekDays := make([]types.ClassesResponse, 5)
	for i := 0; i < 5; i++ {
		y, m, d := addDays(mondayYear, mondayMonth, mondayDay, i)
		weekDays[i] = randomDayResponse(y, m, d, schoolNum, weekNum, selectedClasses)
	}

	state.ClassesResponse = weekDays[0]
	state.NextDayText = weekDays[1].Kopf.DatumPlan

	lastDay := weekDays[4].Kopf.DatumPlan
	state.WeekClassesResponse = types.WeeklyClassesResponse{
		FetchStart: mondayDate,
		FetchEnd:   lastDay,
		Classes:    weekDays,
	}
}

func int16Ptr(v int16) *int16 { return &v }
