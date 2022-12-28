package services

import (
	"behaviouralpatterns/mediator/interfaces"
	"sync"
)

type StationManager struct {
	IsPlatformFree bool
	Lock           *sync.Mutex
	TrainQueue     []interfaces.Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
		Lock:           &sync.Mutex{},
	}

}
func (s *StationManager) CanLand(t interfaces.Train) bool {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, t)
	return false
}

func (s *StationManager) NotifyFree() {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrainInQueue := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}
