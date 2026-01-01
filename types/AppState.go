package types

type TableViewData struct {
	Title         string
	SelectedDate  string
	SelectedClass string

	ValuesPerRow int
	HorHeader    []string
	VerHeader    []string
	Data         [][]CellData
	Notes        []string
}

type CellData struct {
	Value    string
	Canceled bool
	Note     string
}

type AppState struct {
	ActiveUI string

	SelectedSchool   string
	SelectedUsername string
	SelectedPassword string

	SelectedClass string
	SelectedDate  string

	LoginRequested  bool
	LoginInProgress bool
	LoginSuccess    bool
	LoginNote       string
	ClassesResponse ClassesResponse
}
