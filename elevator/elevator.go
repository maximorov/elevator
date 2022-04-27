// elevator
package elevator

type ElevatorState int

const (
	StateWaiting ElevatorState = iota
	StateMoveing ElevatorState = iota
)

type Signal struct {
	In   bool
	Dest int
}

func (s *Signal) IN() bool {
	return s.In == true
}

func (s *Signal) OUT() bool {
	return s.In == false
}

func (s *Signal) DEST() int {
	return s.Dest
}

type Mechanismer interface {
	Opens(floor int)
	Closes(floor int)
	Moves(dest string, floor int)
}

type ButtonsPanel interface {
	PressButton(num int)
	Reset()
	HasPressedButton() bool
}

type ElevatorAction interface {
	Do(*Elevator)
}

type ActionMovesTo struct {
	To int
}

func (s ActionMovesTo) Do(e *Elevator) {
	e.MovesTo(s.To)
}

type ActionOpens struct{}

func (s ActionOpens) Do(e *Elevator) {
	e.Opens(e.floor)
}

type ActionCloses struct{}

func (s ActionCloses) Do(e *Elevator) {
	e.Closes(e.floor)
}

type Elevator struct {
	mechanism    Mechanismer
	floor        int
	buttonsPanel ButtonsPanel
	state        ElevatorState
	actions      []ElevatorAction
}

func New(m Mechanismer, bp ButtonsPanel) *Elevator {
	res := &Elevator{
		m,
		1,
		bp,
		StateWaiting,
		[]ElevatorAction{},
	}

	return res
}

func (s *Elevator) PressButton(f int) {
	s.buttonsPanel.PressButton(f)
}

func (s *Elevator) ResetButton() {
	s.buttonsPanel.Reset()
}

func (s *Elevator) ProcessButtonPress(to chan Signal) {
	for {
		dest := <-to
		switch {
		case dest.IN():
			if dest.DEST() == s.floor {
				s.Do(
					ActionOpens{},
					ActionCloses{},
				)
			} else {
				s.Do(
					ActionMovesTo{dest.DEST()},
					ActionOpens{},
					ActionCloses{},
				)
			}
		case dest.OUT():
			if dest.DEST() == s.floor {
				s.Do(
					ActionOpens{},
					ActionCloses{},
				)
			} else {
				s.Do(
					ActionMovesTo{dest.DEST()},
					ActionOpens{},
					ActionCloses{},
				)
			}
		}
	}
}

func (s *Elevator) Do(ea ...ElevatorAction) {
	s.state = StateMoveing
	for _, a := range ea {
		a.Do(s)
	}
	s.state = StateWaiting
}

func (s *Elevator) Opens(f int) {
	s.mechanism.Opens(f)
}

func (s *Elevator) Closes(f int) {
	s.mechanism.Closes(f)
}

func (s *Elevator) MovesTo(to int) {
	switch {
	case to == s.floor:
	case to < s.floor:
		for to < s.floor {
			s.floor--
			s.mechanism.Moves(`down`, s.floor)
		}
	case to > s.floor:
		for to > s.floor {
			s.floor++
			s.mechanism.Moves(`up`, s.floor)
		}
	}
}
