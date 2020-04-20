package dom

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuerySelector(t *testing.T) {
	n, err := QuerySelector(nil, "div")
	require.NoError(t, err)
	require.Nil(t, n)
}
