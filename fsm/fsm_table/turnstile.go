package fsm_table

type Turnstile struct {
	Transitions []Transition
	State       int
}

func (turnstile *Turnstile) AddTransition(transition Transition) {
	turnstile.Transitions = append(turnstile.Transitions, transition)
}

func (turnstile *Turnstile) Event(event int) {
	for _, transition := range turnstile.Transitions {
		if transition.CurrentState == turnstile.State && transition.Event == event {
			turnstile.State = transition.NewState
			transition.Action()
		}
	}
}
