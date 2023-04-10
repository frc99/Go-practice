package state

type MarioStateMachine struct {
	score      int
	marioState IMarioState
}

func InitialMarioStateMachine() *MarioStateMachine {
	return &MarioStateMachine{
		score:      0,
		marioState: &SmallMario{},
	}
}

func (m *MarioStateMachine) ObtainMushRoom() {
	m.marioState.obtainMushRoom(m)
}

func (m *MarioStateMachine) ObtainFireFlower() {
	m.marioState.obtainFireFlower(m)
}

func (m *MarioStateMachine) MeetMonster() {
	m.marioState.meetMonster(m)
}

func (m *MarioStateMachine) Score() int {
	return m.score
}

func (m *MarioStateMachine) State() State {
	return m.marioState.getState()
}
