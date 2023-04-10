package state

type State string

const (
	StateSmall     = State("Small")
	StateSuper     = State("Super")
	StateSuperFire = State("SuperFire")
	StateDead      = State("Dead")
)

var (
	SmallMarioInstance     = &SmallMario{}
	SuperMarioInstance     = &SuperMario{}
	SuperFireMarioInstance = &SuperFireMario{}
	DeadMarioInstance      = &DeadMario{}
)

type IMarioState interface {
	getState() State
	obtainMushRoom(*MarioStateMachine)
	obtainFireFlower(*MarioStateMachine)
	meetMonster(*MarioStateMachine)
}
// Small
type SmallMario struct {
}

func (m *SmallMario) getState() State {
	return StateSmall
}

func (m *SmallMario) obtainMushRoom(sm *MarioStateMachine) {
	sm.score += 100
	sm.marioState = SuperMarioInstance
}

func (m *SmallMario) obtainFireFlower(sm *MarioStateMachine) {
	// will not happen
}

func (m *SmallMario) meetMonster(sm *MarioStateMachine) {
	sm.marioState = DeadMarioInstance
}

// Super
type SuperMario struct {
}

func (m *SuperMario) getState() State {
	return StateSuper
}

func (m *SuperMario) obtainMushRoom(sm *MarioStateMachine) {
	// will not happen
}

func (m *SuperMario) obtainFireFlower(sm *MarioStateMachine) {
	sm.score += 200
	sm.marioState = SuperFireMarioInstance
}

func (m *SuperMario) meetMonster(sm *MarioStateMachine) {
	sm.marioState = SmallMarioInstance
}

// SuperFire
type SuperFireMario struct {
}

func (m *SuperFireMario) getState() State {
	return StateSuperFire
}

func (m *SuperFireMario) obtainMushRoom(sm *MarioStateMachine) {
	// will not happen
}

func (m *SuperFireMario) obtainFireFlower(sm *MarioStateMachine) {
	sm.score += 200
}

func (m *SuperFireMario) meetMonster(sm *MarioStateMachine) {
	sm.marioState = SmallMarioInstance
}

// Dead
type DeadMario struct {
}

func (m *DeadMario) getState() State {
	return StateDead
}

func (m *DeadMario) obtainMushRoom(sm *MarioStateMachine) {
}

func (m *DeadMario) obtainFireFlower(sm *MarioStateMachine) {
}

func (m *DeadMario) meetMonster(sm *MarioStateMachine) {

}