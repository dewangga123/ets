package nested_fsm

const (
	LOCKED   = 1
	UNLOCKED = 2
	COIN     = 0
	PASS     = 1
)

type Turnstile struct {
	State      int
	Controller TurnstileController
}

func (turnstile *Turnstile) Event(event int) {
	switch turnstile.State {
	case LOCKED:
		switch event {
		case COIN:
			turnstile.State = UNLOCKED
			turnstile.Controller.Unlock()
			break
		case PASS:
			turnstile.Controller.Alarm()
			break
		default:
			break
		}
		break
	case UNLOCKED:
		switch event {
		case COIN:
			turnstile.Controller.Thankyou()
			break
		case PASS:
			turnstile.State = LOCKED
			turnstile.Controller.Lock()
			break
		default:
			break
		}
		break
	default:
		break
	}
}
