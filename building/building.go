package building

import (
	"elevator/elevator"
)

type ButtonsPanel interface {
	PressButton(num int)
	Reset()
}

type Elevator interface {
	ProcessButtonPress(sig chan elevator.Signal)
	Do(ea ...elevator.ElevatorAction)
	PressButton(btn int)
	ResetButton()
}

func New(bf int, e Elevator, eb ButtonsPanel) *Building {
	res := &Building{
		bf, e, eb, make(chan elevator.Signal),
	}

	return res
}

type Building struct {
	floors int
	elr    Elevator
	btns   ButtonsPanel
	sig    chan elevator.Signal
}

func (s *Building) Live() {
	go s.elr.ProcessButtonPress(s.sig)
}

func (s *Building) PressElevatorButton(f int) {
	s.sig <- elevator.Signal{true, f}
}

func (s *Building) PressButton(f int) {
	s.btns.PressButton(f)
	s.sig <- elevator.Signal{false, f}
	s.btns.Reset()
}
