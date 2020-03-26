package state

type TurnstileController interface {
	Lock()
	Unlock()
	Thankyou()
	Alarm()
}
