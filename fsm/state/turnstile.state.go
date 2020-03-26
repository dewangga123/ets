package state

type TurnstileState interface {
	Coin(t *Turnstile)
	Pass(t *Turnstile)
}
