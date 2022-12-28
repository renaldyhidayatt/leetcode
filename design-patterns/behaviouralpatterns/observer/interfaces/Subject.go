package interfaces

type Subject interface {
	Register()
	DeRegister()
	NotifyAll()
}
