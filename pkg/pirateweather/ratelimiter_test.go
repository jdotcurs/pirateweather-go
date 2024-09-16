package pirateweather_test

import (
	"testing"
	"time"

	"github.com/jdotcurs/pirateweather-go/pkg/pirateweather"
	"github.com/stretchr/testify/require"
)

func TestRateLimiter(t *testing.T) {
	rl := pirateweather.NewRateLimiter(5)

	for i := 0; i < 5; i++ {
		require.True(t, rl.Allow())
	}

	require.False(t, rl.Allow())
}

func TestRateLimiterReset(t *testing.T) {
	rl := pirateweather.NewRateLimiter(5)

	for i := 0; i < 5; i++ {
		require.True(t, rl.Allow())
	}

	require.False(t, rl.Allow())

	rl.UpdateFromHeaders(5, 5, time.Now().Add(time.Second))
	time.Sleep(time.Second * 2)

	require.True(t, rl.Allow())
}

func TestRateLimiterUpdateFromHeaders(t *testing.T) {
	rl := pirateweather.NewRateLimiter(5)

	rl.UpdateFromHeaders(10, 8, time.Now().Add(time.Hour))

	for i := 0; i < 8; i++ {
		require.True(t, rl.Allow())
	}

	require.False(t, rl.Allow())
}
