package main

import (
	"elevator/building"
	"elevator/button"
	"elevator/elevator"
	"fmt"
)

const buildingFloors = 9

func main() {
	el := elevator.New(elevator.NewMechanism(), button.NewPanel(buildingFloors))
	bld := building.New(buildingFloors, el, button.NewPanel(buildingFloors))
	bld.Live()

	fmt.Println(`Press From and To floors: `)

	for {
		var from int
		var to int
		_, _ = fmt.Scan(&from, &to)

		bld.PressButton(from)
		bld.PressElevatorButton(to)
	}
}
