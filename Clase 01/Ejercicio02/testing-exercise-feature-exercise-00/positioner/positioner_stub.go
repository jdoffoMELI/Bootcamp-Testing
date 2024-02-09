package positioner

// PositionerStub is a stub for Positioner
type PositionerStub struct {
	GetLinearDistanceFunc func(from, to *Position) (linearDistance float64)
}

// NewPositionerStub creates a new PositionerStub
func NewPositionerStub(distance float64) *PositionerStub {
	return &PositionerStub{}
}

// GetLinearDistance returns the linear distance between 2 positions (in meters)
func (p *PositionerStub) GetLinearDistance(from, to *Position) (linearDistance float64) {
	return p.GetLinearDistanceFunc(from, to)
}
