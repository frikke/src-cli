package output

import "time"

// StatusBar is a sub-element of a progress bar that displays the current status
// of a process.
type StatusBar struct {
	completed bool

	label  string
	format string
	args   []interface{}

	startedAt  time.Time
	finishedAt time.Time
}

// Completef sets the StatusBar to completed and updates its text.
func (sb *StatusBar) Completef(format string, args ...interface{}) {
	sb.completed = true
	sb.format = format
	sb.args = args
	sb.finishedAt = time.Now()
}

// Resetf sets the status of the StatusBar to incomplete and updates its label and text.
func (sb *StatusBar) Resetf(label, format string, args ...interface{}) {
	sb.completed = false
	sb.label = label
	sb.format = format
	sb.args = args
	sb.startedAt = time.Now()
}

// Updatef updates the StatusBar's text.
func (sb *StatusBar) Updatef(format string, args ...interface{}) {
	sb.format = format
	sb.args = args
}

func (sb *StatusBar) runtime() time.Duration {
	if sb.startedAt.IsZero() {
		return 0
	}
	if sb.finishedAt.IsZero() {
		return time.Since(sb.startedAt).Truncate(time.Second)
	}

	return sb.finishedAt.Sub(sb.startedAt).Truncate(time.Second)
}

func NewStatusBarWithLabel(label string) *StatusBar {
	return &StatusBar{label: label, startedAt: time.Now()}
}
