package state

type UnlockState struct{}

func (state *UnlockState) Coin(t *Turnstile) {
	t.Alarm()
}
func (state *UnlockState) Pass(t *Turnstile) {
	t.Lock()
	t.SetLocked()
}
