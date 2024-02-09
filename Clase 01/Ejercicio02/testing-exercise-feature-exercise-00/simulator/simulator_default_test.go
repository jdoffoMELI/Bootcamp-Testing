package simulator_test

import (
	"testdoubles02/positioner"
	"testdoubles02/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefualt_CanCatch(t *testing.T) {
	t.Run("Hunter hunts his prey", func(t *testing.T) {
		/* Arrange */
		// HunterSubject definition
		hunterSubject := &simulator.Subject{
			Position: &positioner.Position{
				X: 5,
				Y: 0,
				Z: 0,
			},
			Speed: 10,
		}
		// PreySubject definition
		preySubject := &simulator.Subject{
			Position: &positioner.Position{
				X: 5,
				Y: 0,
				Z: 5,
			},
			Speed: 5,
		}
		// CatchSimulatorDefault definition
		positioner_stub := positioner.NewPositionerStub()
		positioner_stub.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			return 5
		}
		simulator := simulator.NewCatchSimulatorDefault(3, positioner_stub)

		/* Act */
		canCatch := simulator.CanCatch(hunterSubject, preySubject)

		/* Assert */
		require.True(t, canCatch)
	})

	t.Run("Hunter doesn't hunt his prey, due its slowness", func(t *testing.T) {
		/* Arrange */
		// HunterSubject definition
		hunterSubject := &simulator.Subject{
			Position: &positioner.Position{
				X: 5,
				Y: 0,
				Z: 0,
			},
			Speed: 3,
		}
		// PreySubject definition
		preySubject := &simulator.Subject{
			Position: &positioner.Position{
				X: 5,
				Y: 0,
				Z: 5,
			},
			Speed: 5,
		}
		// CatchSimulatorDefault definition
		positioner_stub := positioner.NewPositionerStub()
		positioner_stub.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			return 5
		}
		simulator := simulator.NewCatchSimulatorDefault(3, positioner_stub)

		/* Act */
		canCatch := simulator.CanCatch(hunterSubject, preySubject)

		/* Assert */
		require.False(t, canCatch)
	})

	t.Run("Hunter doesn't hunt his prey, due its lack of time", func(t *testing.T) {
		/* Arrange */
		// HunterSubject definition
		hunterSubject := &simulator.Subject{
			Position: &positioner.Position{
				X: 5,
				Y: 0,
				Z: 0,
			},
			Speed: 8,
		}
		// PreySubject definition
		preySubject := &simulator.Subject{
			Position: &positioner.Position{
				X: 5,
				Y: 0,
				Z: 5,
			},
			Speed: 5,
		}
		// CatchSimulatorDefault definition
		positioner_stub := positioner.NewPositionerStub()
		positioner_stub.GetLinearDistanceFunc = func(from, to *positioner.Position) (linearDistance float64) {
			return 5
		}
		simulator := simulator.NewCatchSimulatorDefault(1, positioner_stub)

		/* Act */
		canCatch := simulator.CanCatch(hunterSubject, preySubject)

		/* Assert */
		require.False(t, canCatch)
	})
}
