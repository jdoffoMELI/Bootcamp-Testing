package prey

import "testdoubles02/positioner"

// PreyStub is a stub for Prey
type PreyStub struct {
	GetSpeedFunc    func() float64
	GetPositionFunc func() *positioner.Position
}

// NewPreyStub creates a new PreyStub
func NewPreyStub() *PreyStub {
	return &PreyStub{}
}

// GetSpeed returns the speed of the prey
func (p *PreyStub) GetSpeed() (speed float64) {
	return p.GetSpeedFunc()
}

// GetPosition returns the position of the prey
func (p *PreyStub) GetPosition() (position *positioner.Position) {
	return p.GetPositionFunc()
}
