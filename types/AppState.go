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

	Login LoginState

	ClassesResponse ClassesResponse

	TableViewData TableViewData
	DayViewState  DayViewState
}

// Dayview related states
type DayViewState struct {
	Lessons []LessonDisplayData
}

type LessonDisplayData struct {
	Beginn string
	Ende   string
	Fa     ValueWithNote
	Le     ValueWithNote
}

type ValueWithNote struct {
	Value string
	Note  string
}

// Login related states

type LoginState struct {
	LoginRequested               bool
	LoginInProgress              bool
	LoginSuccess                 bool
	LoginNote                    string
	RememberLogin                bool
	RecentLogin                  LoginCredentials
	RecentLoginDeletionRequested bool
}

type LoginCredentials struct {
	School   string
	Username string
	Password string
}
