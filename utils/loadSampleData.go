package utils

import (
	"encoding/xml"
	"vpdesktop/types"
)

func LoadSampleData(state *types.AppState) {

	state.ClassesResponse = types.ClassesResponse{
		XMLName: xml.Name{Local: "VpMobil"},
		Kopf: types.Kopf{
			Planart:      "KL",
			Zeitstempel:  "2026-01-15T08:30:00",
			DatumPlan:    "2026-01-15",
			Datei:        "klassen.xml",
			Nativ:        1,
			Woche:        3,
			TageProWoche: 5,
			Schulnummer:  "123456",
		},
		FreieTage: types.FreieTage{
			FT: []int{6, 7},
		},
		Klassen: types.Klassen{
			Klassen: []types.Klasse{

				// ---------- 10A ----------
				{
					Kurz:      "10A",
					Hash:      "hash-10a",
					KlStunden: fullDayStunden(),
					Plan: types.Plan{
						Stunden: []types.Stunde{
							st(1, "08:00", "08:45", "MA", "ME", "R101", "normal"),
							st(2, "08:45", "09:30", "EN", "SM", "R102", "normal"),
							st(3, "09:45", "10:30", "PH", "WE", "R201", "normal"),
							st(4, "10:30", "11:15", "PH", "WE", "R201", "normal"),
							st(5, "11:30", "12:15", "GE", "KR", "R105", "normal"),
							st(6, "12:15", "13:00", "GE", "KR", "R105", "normal"),
							st(7, "13:45", "14:30", "SP", "HN", "GYM", "normal"),
							st(8, "14:30", "15:15", "SP", "HN", "GYM", "normal"),
							st(9, "15:30", "16:15", "IT", "AL", "LAB2", "normal"),
						},
					},
				},

				// ---------- 9B ----------
				{
					Kurz:      "9B",
					Hash:      "hash-9b",
					KlStunden: fullDayStunden(),
					Plan: types.Plan{
						Stunden: []types.Stunde{
							st(1, "08:00", "08:45", "BI", "KR", "LAB1", "normal"),
							st(2, "08:45", "09:30", "MA", "ME", "R103", "normal"),
							st(3, "09:45", "10:30", "SP", "HN", "GYM", "vertretung"),
							st(4, "10:30", "11:15", "EN", "SM", "R104", "normal"),
						},
					},
				},

				// ---------- 8C ----------
				{
					Kurz:      "8C",
					Hash:      "hash-8c",
					KlStunden: fullDayStunden(),
					Plan: types.Plan{
						Stunden: []types.Stunde{
							st(1, "08:00", "08:45", "DE", "MU", "R110", "normal"),
							st(2, "08:45", "09:30", "DE", "MU", "R110", "normal"),
							st(3, "09:45", "10:30", "MA", "ME", "R111", "normal"),
							st(4, "10:30", "11:15", "KU", "BA", "ART", "normal"),
							st(5, "11:30", "12:15", "MU", "KL", "MUS", "normal"),
						},
					},
				},

				// ---------- 7A ----------
				{
					Kurz:      "7A",
					Hash:      "hash-7a",
					KlStunden: fullDayStunden(),
					Plan: types.Plan{
						Stunden: []types.Stunde{
							st(1, "08:00", "08:45", "MA", "ME", "R101", "normal"),
							st(2, "08:45", "09:30", "MA", "ME", "R101", "normal"),
							st(3, "09:45", "10:30", "EN", "SM", "R102", "normal"),
							st(4, "10:30", "11:15", "GE", "KR", "R105", "normal"),
						},
					},
				},
			},
		},
		ZusatzInfo: types.ZusatzInfo{
			Zeilen: []string{
				"Stand: 15.01.2026",
				"Beispieldaten – nicht produktiv",
			},
		},
	}
}

// ---------- helpers (local only) ----------

func fullDayStunden() types.KlStunden {
	return types.KlStunden{
		Stunden: []types.KlStunde{
			{Value: 1, ZeitVon: "08:00", ZeitBis: "08:45"},
			{Value: 2, ZeitVon: "08:45", ZeitBis: "09:30"},
			{Value: 3, ZeitVon: "09:45", ZeitBis: "10:30"},
			{Value: 4, ZeitVon: "10:30", ZeitBis: "11:15"},
			{Value: 5, ZeitVon: "11:30", ZeitBis: "12:15"},
			{Value: 6, ZeitVon: "12:15", ZeitBis: "13:00"},
			{Value: 7, ZeitVon: "13:45", ZeitBis: "14:30"},
			{Value: 8, ZeitVon: "14:30", ZeitBis: "15:15"},
			{Value: 9, ZeitVon: "15:30", ZeitBis: "16:15"},
		},
	}
}

func st(n int8, von, bis, fa, le, ra, info string) types.Stunde {
	return types.Stunde{
		St:     n,
		Beginn: von,
		Ende:   bis,
		Fa:     types.Fach{Value: fa},
		Le:     types.Lehrer{Value: le},
		Ra:     types.Raum{Value: ra},
		If:     info,
	}
}

func int16Ptr(v int16) *int16 {
	return &v
}
