package main

type StationManeger struct {
	isPlatformFree bool
	trainQueue     []Train
}

func newStationManger() *StationManeger {
	return &StationManeger{
		isPlatformFree: true,
	}
}

func (s *StationManeger) canArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManeger) notifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}

	if len(s.trainQueue) > 0 {
		firstTrainQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainQueue.permitArrival()
	}
}
