package graph

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func connect(ctx context.Context, url string, ssl bool) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithBlock(),
	}
	if ssl {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.DialContext(ctx, url, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", url, err)
	}

	glog.Info("Connected to graph server")

	return conn, nil
}

func watch(ctx context.Context, conn *grpc.ClientConn) {
	//noinspection GoUnhandledErrorResult
	defer conn.Close()

	for st, chg := connectivity.Ready, false; ; {
		if !conn.WaitForStateChange(ctx, st) {
			break
		}

		chg = st == connectivity.Ready
		st = conn.GetState()

		if f := "graph connection lost"; chg {
			glog.Info(f)
		}

		switch st {
		case connectivity.Idle:
		case connectivity.Connecting:
		case connectivity.TransientFailure:
		case connectivity.Ready:
			goto reconnected
		case connectivity.Shutdown:
			return
		}

		continue

	reconnected:
		glog.Info("graph client reconnected")
	}
}
