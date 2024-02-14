package hunter_test

import (
	"testdoubles02/hunter"
	"testdoubles02/positioner"
	"testdoubles02/prey"
	"testdoubles02/simulator"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestWhiteShark_hunt(t *testing.T) {
	t.Run("hunt should return nil when the shark can hunt it's prey", func(t *testing.T) {
		/* Preparing dependencies */
		// simulator mock
		simulator := simulator.NewSimulatorMock()
		simulator.On("CanCatch", mock.Anything, mock.Anything).Return(true)
		// prey stub
		prey := prey.NewPreyStub()
		prey.GetSpeedFunc = func() float64 { return 10.0 }
		prey.GetPositionFunc = func() *positioner.Position { return &positioner.Position{X: 10.0, Y: 10.0, Z: 10.0} }
		// hunter
		hunter := hunter.CreateWhiteShark(simulator)

		/* Executing the test */
		err := hunter.Hunt(prey)

		/* Validating the result */
		require.Nil(t, err)
		simulator.AssertExpectations(t)
	})

	t.Run("hunt should return an error when the shark can't hunt it's prey", func(t *testing.T) {
		/* Preparing dependencies */
		// simulator mock
		simulator := simulator.NewSimulatorMock()
		simulator.On("CanCatch", mock.Anything, mock.Anything).Return(false)
		// prey stub
		prey := prey.NewPreyStub()
		prey.GetSpeedFunc = func() float64 { return 10.0 }
		prey.GetPositionFunc = func() *positioner.Position { return &positioner.Position{X: 10.0, Y: 10.0, Z: 10.0} }
		// hunter
		hunter := hunter.CreateWhiteShark(simulator)

		/* Executing the test */
		err := hunter.Hunt(prey)

		/* Validating the result */
		require.Error(t, err)
		simulator.AssertExpectations(t)
	})
}
