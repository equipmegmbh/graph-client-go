package graph

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/connectivity"
)

func TestConnect(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := connect(ctx, "passthrough:///127.0.0.1:8230", false)
	require.NoError(t, err)
	require.NotNil(t, conn)
	require.Equal(t, connectivity.Ready, conn.GetState())
}

func TestWatch(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := connect(ctx, "passthrough:///127.0.0.1:8230", false)
	require.NoError(t, err)
	go watch(ctx, conn)

	time.Sleep(200 * time.Millisecond)
	require.Equal(t, connectivity.Ready, conn.GetState())
}
