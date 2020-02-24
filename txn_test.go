package intelligence

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimpleQuery(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cli, err := connect(ctx, "127.0.0.1:8230", true)

	defer cancel()

	require.NoError(t, err)

	txn := cli.NewTxn("knowledge")

	require.NotNil(t, txn)

	query := txn.Query()

	query.Select("given-name", "family-name")
	query.From("person")

	require.NotNil(t, query)
	require.NoError(t, txn.Commit(context.Background()))
}
