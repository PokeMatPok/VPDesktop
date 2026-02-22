package types

import (
	"time"

	"gioui.org/layout"
	"gioui.org/widget"
)

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

	ShowDatePicker bool

	SelectingFavorites bool
	ActiveFavoriteSlot int

	ViewMode string // "day", "week", "month"

	Login LoginState

	ClassesResponse     ClassesResponse
	WeekClassesResponse WeeklyClassesResponse

	NextDayText   string
	TableViewData TableViewData
	DayViewState  DayViewState
	WeekViewState WeekViewState

	ClassClickables map[string]*widget.Clickable
	ClassList       layout.List

	AnimationStates map[string]*AnimationState

	lastFrame time.Time
}

type AnimationState struct {
	Progress       float32
	AnimatingIn    bool
	AnimationState string
	StartTime      time.Time
	Duration       time.Duration
}

// Dayview related states
type DayViewState struct {
	Lessons []LessonDisplayData
}

type WeekViewState struct {
	Days []DayData
}

type DayData struct {
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
	LoginPhase                   string // "school_entry", "user_selection", "password_entry"
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
