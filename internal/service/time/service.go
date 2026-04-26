package timesvc

import (
	"time"
)

// Service provides current time.
type Service interface {
	Now() time.Time
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Now() time.Time {
	return time.Now().UTC()
}
