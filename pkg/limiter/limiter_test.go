package limiter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLimiterMax(t *testing.T) {
	r := require.New(t)
	limiter := NewLimiter(100)
	for i := 0; i < 100; i++ {
		err := limiter.AddTask()
		r.NoError(err)
	}
	err := limiter.AddTask()
	require.Error(t, err)
}

func TestLimiterDecreaseBelow(t *testing.T) {
	r := require.New(t)
	limiter := NewLimiter(100)
	r.Equal(int64(0), limiter.curEvents.Load())
	limiter.DoneTask()
	r.Equal(int64(0), limiter.curEvents.Load())
}
