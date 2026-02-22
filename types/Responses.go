package types

import "encoding/xml"

/*
Root
*/

type WeeklyClassesResponse struct {
	FetchStart string            `xml:"fetchStart"`
	FetchEnd   string            `xml:"fetchEnd"`
	Classes    []ClassesResponse `xml:"classes>class"`
}

type ClassesResponse struct {
	XMLName    xml.Name   `xml:"VpMobil"`
	Kopf       Kopf       `xml:"Kopf"`
	FreieTage  FreieTage  `xml:"FreieTage"`
	Klassen    Klassen    `xml:"Klassen"`
	ZusatzInfo ZusatzInfo `xml:"ZusatzInfo"`
}

type TeachersResponse struct {
	XMLName    xml.Name                  `xml:"VpMobil"`
	Kopf       Kopf                      `xml:"Kopf"`
	FreieTage  FreieTage                 `xml:"FreieTage"`
	Klassen    KlassenInTeachersResponse `xml:"Klassen"`
	ZusatzInfo ZusatzInfo                `xml:"ZusatzInfo"`
}

type RoomsResponse struct {
	XMLName    xml.Name   `xml:"VpMobil"`
	Kopf       Kopf       `xml:"Kopf"`
	FreieTage  FreieTage  `xml:"FreieTage"`
	Klassen    Klassen    `xml:"Klassen"`
	ZusatzInfo ZusatzInfo `xml:"ZusatzInfo"`
}

/*
Kopf
*/

type Kopf struct {
	Planart      string `xml:"planart"`
	Zeitstempel  string `xml:"zeitstempel"`
	DatumPlan    string `xml:"DatumPlan"`
	Datei        string `xml:"datei"`
	Nativ        int8   `xml:"nativ"`
	Woche        int8   `xml:"woche"`
	TageProWoche int8   `xml:"tageprowoche"`
	Schulnummer  string `xml:"schulnummer"`
}

/*
FreieTage
*/

type FreieTage struct {
	FT []int `xml:"ft"`
}

/*
Aufsichten
*/

type Aufsichten struct {
	Aufsicht *Aufsicht `xml:"Aufsicht,omitempty"`
}

type Aufsicht struct {
	AuTag       int8   `xml:"AuTag"`
	AuVorStunde int8   `xml:"AuVorStunde"`
	AuUhrzeit   string `xml:"AuUhrzeit"`
	AuZeit      string `xml:"AuZeit"`
	AuOrt       string `xml:"AuOrt"`
	AuFuer      string `xml:"AuFuer,omitempty"`
	AuInfo      string `xml:"AuInfo,omitempty"`
	AuAe        string `xml:"AuAe,attr,omitempty"`
}

/*
Klassen
*/

type Klassen struct {
	Klassen []Klasse `xml:"Kl"`
}

type Klasse struct {
	Kurz       string     `xml:"Kurz"`
	Hash       string     `xml:"Hash"`
	KlStunden  KlStunden  `xml:"KlStunden"`
	Kurse      Kurse      `xml:"Kurse"`
	Unterricht Unterricht `xml:"Unterricht"`
	Plan       Plan       `xml:"Pl"`
}

type KlassenInTeachersResponse struct {
	Klassen []KlasseInTeachersResponse `xml:"Kl"`
}

type KlasseInTeachersResponse struct {
	Kurz       string     `xml:"Kurz"`
	Hash       string     `xml:"Hash"`
	KlStunden  KlStunden  `xml:"KlStunden"`
	Plan       Plan       `xml:"Pl"`
	Aufsichten Aufsichten `xml:"Aufsichten,omitempty"` // new
}

type KlassenInRoomsResponse struct {
	Klassen []KlasseInRoomsResponse `xml:"Kl"`
}

type KlasseInRoomsResponse struct {
	Hash      string    `xml:"Hash"`
	Kurz      string    `xml:"Kurz"`
	KlStunden KlStunden `xml:"KlStunden"`
	Plan      Plan      `xml:"Pl"`
}

/*
KlStunden
*/

type KlStunden struct {
	Stunden []KlStunde `xml:"KlSt"`
}

type KlStunde struct {
	Value   int8   `xml:",chardata"`
	ZeitVon string `xml:"ZeitVon,attr,omitempty"`
	ZeitBis string `xml:"ZeitBis,attr,omitempty"`
}

/*
Kurse
*/

type Kurse struct {
	Kurse []Kurs `xml:"Ku"`
}

type Kurs struct {
	KKz KKz `xml:"KKz"`
}

type KKz struct {
	Value string `xml:",chardata"`
	KLe   string `xml:"KLe,attr,omitempty"`
}

/*
Unterricht
*/

type Unterricht struct {
	Einheiten []Unterrichtseinheit `xml:"Ue"`
}

type Unterrichtseinheit struct {
	Nummer UeNr `xml:"UeNr"`
}

type UeNr struct {
	Value int16  `xml:",chardata"`
	UeLe  string `xml:"UeLe,attr,omitempty"`
	UeFa  string `xml:"UeFa,attr,omitempty"`
	UeGr  string `xml:"UeGr,attr,omitempty"`
}

/*
Plan (Pl)
*/

type Plan struct {
	Stunden []Stunde `xml:"Std"`
}

type Stunde struct {
	St     int8   `xml:"St"`
	Beginn string `xml:"Beginn"`
	Ende   string `xml:"Ende"`
	Fa     Fach   `xml:"Fa"`
	Ku2    string `xml:"Ku2,omitempty"`
	Le     Lehrer `xml:"Le"`
	Ra     Raum   `xml:"Ra"`
	Nr     *int16 `xml:"Nr,omitempty"`
	If     string `xml:"If"`
}

/*
Fach / Lehrer / Raum
*/

type Fach struct {
	Value string `xml:",chardata"`
	FaAe  string `xml:"FaAe,attr,omitempty"`
}

type Lehrer struct {
	Value string `xml:",chardata"`
	LeAe  string `xml:"LeAe,attr,omitempty"`
}

type Raum struct {
	Value string `xml:",chardata"`
	RaAe  string `xml:"RaAe,attr,omitempty"`
}

/*
ZusatzInfo
*/

type ZusatzInfo struct {
	Zeilen []string `xml:"ZiZeile"`
}
