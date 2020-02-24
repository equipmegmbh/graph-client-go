package intelligence

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/golang/glog"
	"github.com/verticalgmbh/intelligence-core/pkg/pb"
)

type Txn struct {
	set string

	mux sync.RWMutex
	cli pb.ApiClient

	queries   []*pb.Query
	mutations []*pb.Mutation
}

func (txn *Txn) Query() *pb.Query {
	// Collect all errors and, if there are
	// any, return them as soon as commit is called

	query := new(pb.Query)

	txn.mux.Lock()
	txn.queries = append(txn.queries, query)
	txn.mux.Unlock()

	return query
}

func (txn *Txn) Commit(ctx context.Context) error {
	// Validate queries and mutations

	request := &pb.Request{Set: txn.set}

	request.Queries = append(request.Queries, txn.queries...)
	request.Mutations = append(request.Mutations, txn.mutations...)

	stream, err := txn.cli.Query(ctx, request)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	for {
		reply, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if f := "query failed: %v"; err != nil {
			return fmt.Errorf(f, err)
		}

		glog.Infof("response: %v", reply)
	}

	return nil
}

func (txn *Txn) Discard(ctx context.Context) error {
	return nil
}
