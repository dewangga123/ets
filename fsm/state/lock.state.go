package state

type LockState struct {
}

func (state *LockState) Coin(t *Turnstile) {
	t.Unlock()
	t.SetUnlocked()
}
func (state *LockState) Pass(t *Turnstile) {
	t.Thankyou()
}
