// elevator
package elevator

import (
	"fmt"
	"time"
)

type Mechanism struct {
}

func NewMechanism() *Mechanism {
	res := &Mechanism{}

	return res
}

func (s *Mechanism) Opens(floor int) {
	fmt.Printf("Elevator is opening door on %d\n", floor)
	time.Sleep(1000 * time.Millisecond)
}

func (s *Mechanism) Closes(floor int) {
	fmt.Printf("Elevator is closing door on %d\n", floor)
	time.Sleep(1000 * time.Millisecond)
}

func (s *Mechanism) Moves(dest string, floor int) {
	fmt.Printf("Elevator is moveing %s: %d\n", dest, floor)
	time.Sleep(1000 * time.Millisecond)
}
