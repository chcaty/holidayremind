package scheduler

type UnitType int

const (
	Second UnitType = iota
	Minute
	Hour
	Day
)
