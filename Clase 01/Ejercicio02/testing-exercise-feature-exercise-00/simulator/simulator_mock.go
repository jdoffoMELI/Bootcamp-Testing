package simulator

import "github.com/stretchr/testify/mock"

// SimulatorMock is a mock for Simulator
type SimulatorMock struct {
	mock.Mock
}

// NewSimulatorMock creates a new SimulatorMock
func NewSimulatorMock() *SimulatorMock {
	return &SimulatorMock{}
}

// CanCatch returns true if the hunter can catch the prey
func (s *SimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	args := s.Called(hunter, prey)
	return args.Bool(0)
}
