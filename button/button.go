package button

import "fmt"

type ButtonState int

const (
	StateOn  ButtonState = iota
	StateOff ButtonState = iota
)

func New(f int) Button {
	res := Button{
		f, StateOff,
	}

	return res
}

type Button struct {
	floor int
	state ButtonState
}

func (s *Button) Floor() int {
	return s.floor
}

func (s *Button) Pressed() bool {
	return s.state == StateOn
}

func (s *Button) Press() {
	s.state = StateOn
}

func (s *Button) Unpress() {
	s.state = StateOff
}

func NewPanel(num int) *Panel {
	res := make([]Button, num, num)
	for i := 0; i < num; i++ {
		res[i] = New(i + 1)
	}

	return &Panel{res}
}

type Panel struct {
	btns []Button
}

func (s *Panel) PressButton(num int) {
	for _, btn := range s.btns {
		if btn.Pressed() {
			//return error, instead of
			fmt.Printf("One button already pressed\n")
			return
		}
	}
	for i := range s.btns {
		if s.btns[i].Floor() == num {
			s.btns[i].Press()
			return
		}
	}
}

func (s *Panel) Reset() {
	for i := range s.btns {
		if s.btns[i].Pressed() {
			s.btns[i].Unpress()
		}
	}
}

func (s *Panel) HasPressedButton() bool {
	for i := range s.btns {
		if s.btns[i].Pressed() {
			return true
		}
	}

	return false
}
