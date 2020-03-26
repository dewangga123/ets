package state

var (
	lockState   = &LockState{}
	unlockState = &UnlockState{}
)

type Turnstile struct {
	State      TurnstileState
	Controller TurnstileController
}

func (turnstile *Turnstile) Coin() {
	turnstile.State.Coin(turnstile)
}

func (turnstile *Turnstile) Pass() {
	turnstile.State.Pass(turnstile)
}

func (turnstile *Turnstile) SetLocked() {
	turnstile.State = lockState
}

func (turnstile *Turnstile) SetUnlocked() {
	turnstile.State = unlockState
}

func (turnstile *Turnstile) IsLocked() bool {
	return turnstile.State == lockState
}

func (turnstile *Turnstile) IsUnlocked() bool {
	return turnstile.State == unlockState
}

func (turnstile *Turnstile) Lock() {
	turnstile.Controller.Lock()
}
func (turnstile *Turnstile) Unlock() {
	turnstile.Controller.Unlock()
}
func (turnstile *Turnstile) Thankyou() {
	turnstile.Controller.Thankyou()
}
func (turnstile *Turnstile) Alarm() {
	turnstile.Controller.Alarm()
}
