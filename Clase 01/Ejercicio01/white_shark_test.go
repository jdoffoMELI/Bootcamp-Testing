package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		/* Preparing dependencies */
		whiteShark := hunt.NewWhiteShark(true, false, 100)
		tuna := hunt.NewTuna("Phillip", 50)

		/* Executing the method */
		err := whiteShark.Hunt(tuna)

		/* Validating the result */
		require.NoError(t, err)
		require.Equal(t, whiteShark.Hungry, false)
		require.Equal(t, whiteShark.Tired, true)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		/* Preparing dependencies */
		whiteShark := hunt.NewWhiteShark(false, false, 100)
		tuna := hunt.NewTuna("Phillip", 50)

		/* Executing the method */
		err := whiteShark.Hunt(tuna)

		/* Validating the result */
		require.Equal(t, err, hunt.ErrSharkIsNotHungry)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		/* Preparing dependencies */
		whiteShark := hunt.NewWhiteShark(true, true, 100)
		tuna := hunt.NewTuna("Phillip", 50)

		/* Executing the method */
		err := whiteShark.Hunt(tuna)

		/* Validating the result */
		require.Equal(t, err, hunt.ErrSharkIsTired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		/* Preparing dependencies */
		whiteShark := hunt.NewWhiteShark(true, false, 50)
		tuna := hunt.NewTuna("Phillip", 100)

		/* Executing the method */
		err := whiteShark.Hunt(tuna)

		/* Validating the result */
		require.Equal(t, err, hunt.ErrSharkIsSlower)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		/* Preparing dependencies */
		whiteShark := hunt.NewWhiteShark(true, false, 100)

		/* Executing the method */
		err := whiteShark.Hunt(nil)

		/* Validating the result */
		require.Equal(t, err, hunt.ErrTunaIsNil)
	})
}
