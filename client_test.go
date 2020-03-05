package intelligence

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubscribe(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cli, err := connect(ctx, "passthrough:///127.0.0.1:8230", true)

	defer cancel()

	require.NoError(t, err)
	require.NotNil(t, cli)
}
