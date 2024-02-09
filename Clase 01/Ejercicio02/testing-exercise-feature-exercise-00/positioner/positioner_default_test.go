package positioner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearDistance(t *testing.T) {
	t.Run("success: both coordinates are negative", func(t *testing.T) {
		/* Preparing dependencies */
		positioner := NewPositionerDefault()
		c1 := Position{-1, -1, -2}
		c2 := Position{-1, -1, -4}
		expected_value := 2.0

		/* Executing the function */
		distance := positioner.GetLinearDistance(&c1, &c2)

		/* Checking the result */
		require.Equal(t, expected_value, distance, "The distance should be 2.0")
	})

	t.Run("success: both coordinates are positive", func(t *testing.T) {
		/* Preparing dependencies */
		positioner := NewPositionerDefault()
		c1 := Position{2, 1, 3}
		c2 := Position{4, 4, 4}
		expected_value := 3.7416573867739413

		/* Executing the function */
		distance := positioner.GetLinearDistance(&c1, &c2)

		/* Checking the result */
		require.Equal(t, expected_value, distance, "The distance should be 3.7416573867739413")
	})

	t.Run("success: both coordinates are positive and the distance is an integer number", func(t *testing.T) {
		/* Preparing dependencies */
		positioner := NewPositionerDefault()
		c1 := Position{2, 2, 3}
		c2 := Position{2, 2, 8}
		expected_value := 5.0

		/* Executing the function */
		distance := positioner.GetLinearDistance(&c1, &c2)

		/* Checking the result */
		require.Equal(t, expected_value, distance, "The distance should be 5.0")
	})

}
