package libzone

type state byte

const (
	Installed state = iota
	Ready
	Running
	Configured
	Incomplete
	ShuttingDown
	Down
)

func (s state) Installed() state {
	return Installed
}

func (s state) Ready() state {
	return Ready
}

func (s state) Running() state {
	return Running
}

func (s state) Configured() state {
	return Configured
}

func (s state) Incomplete() state {
	return Incomplete
}

func (s state) ShuttingDown() state {
	return ShuttingDown
}

func (s state) Down() state {
	return Down
}

const State state = 255
