package libzone

import "strconv"

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

// Function to return the human readable version of a brand
func (s state) String() string {
	//Convert brand byte to string
	switch s {
	case Installed:
		return "Installed"
	case Ready:
		return "Ready"
	case Running:
		return "Running"
	case Configured:
		return "Configured"
	case Incomplete:
		return "Incomplete"
	case ShuttingDown:
		return "ShuttingDown"
	case Down:
		return "Down"
	}

	return "Unknown: " + strconv.Itoa(int(s))
}

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
