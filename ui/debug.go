package ui

type Debugger struct {
	Enabled bool
	// if specified, run until here
	BreakPoint uint16

	Stepped bool
}
