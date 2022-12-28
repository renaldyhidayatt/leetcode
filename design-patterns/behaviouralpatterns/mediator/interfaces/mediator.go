package interfaces

type Mediator interface {
	CanLand(Train) bool
	NotifyFree()
}
