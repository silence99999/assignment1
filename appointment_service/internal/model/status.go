package model

type Status string

const (
	StatusNew        = "NEW"
	StatusInProgress = "IN_PROGRESS"
	StatusDone       = "DONE"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return true

	default:
		return false
	}

}
