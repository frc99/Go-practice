package main

import (
	"fmt"
	"test/mario/state"
)

func main() {
	fsm := state.InitialMarioStateMachine()
	fmt.Printf("%v %v\n", fsm.State(), fsm.Score())
	fsm.MeetMonster()
	fmt.Printf("%v %v\n", fsm.State(), fsm.Score())

	fsm = state.InitialMarioStateMachine()
	fmt.Printf("%v %v\n", fsm.State(), fsm.Score())
	fsm.ObtainMushRoom()
	fmt.Printf("%v %v\n", fsm.State(), fsm.Score())
	fsm.ObtainFireFlower()
	fmt.Printf("%v %v\n", fsm.State(), fsm.Score())
	fsm.MeetMonster()
	fmt.Printf("%v %v\n", fsm.State(), fsm.Score())
}
