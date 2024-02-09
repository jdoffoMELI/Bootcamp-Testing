package prey_test

import (
	"testdoubles02/positioner"
	"testdoubles02/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTuna_GetSpeed(t *testing.T) {
	t.Run("tuna with default values", func(t *testing.T) {
		/* Preparing dependencies */
		tuna := prey.NewTuna(0, nil)
		spectedSpeed := 0.0

		/* Executing the test */
		speed := tuna.GetSpeed()

		/* Validating the result */
		require.Equal(t, spectedSpeed, speed)
	})

	t.Run("tuna with custom values", func(t *testing.T) {
		/* Preparing dependencies */
		tuna := prey.NewTuna(100.0, nil)
		spectedSpeed := 100.0

		/* Executing the test */
		speed := tuna.GetSpeed()

		/* Validating the result */
		require.Equal(t, spectedSpeed, speed)
	})
}

func TestTuna_GetPosition(t *testing.T) {
	t.Run("tuna with default values", func(t *testing.T) {
		/* Preparing dependencies */
		tuna := prey.NewTuna(0, &positioner.Position{})
		spectedPosition := positioner.Position{X: 0, Y: 0, Z: 0}

		/* Executing the test */
		position := tuna.GetPosition()

		/* Validating the result */
		require.Equal(t, spectedPosition.X, position.X)
		require.Equal(t, spectedPosition.Y, position.Y)
		require.Equal(t, spectedPosition.Z, position.Z)
	})

	t.Run("tuna with custom values", func(t *testing.T) {
		/* Preparing dependencies */
		tuna := prey.NewTuna(0, &positioner.Position{X: 100, Y: 200, Z: 300})
		spectedPosition := positioner.Position{X: 100, Y: 200, Z: 300}

		/* Executing the test */
		position := tuna.GetPosition()

		/* Validating the result */
		require.Equal(t, spectedPosition.X, position.X)
		require.Equal(t, spectedPosition.Y, position.Y)
		require.Equal(t, spectedPosition.Z, position.Z)
	})
}
