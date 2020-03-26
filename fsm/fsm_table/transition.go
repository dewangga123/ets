package fsm_table

type Transition struct {
	CurrentState int
	Event        int
	NewState     int
	Action       func()
}
