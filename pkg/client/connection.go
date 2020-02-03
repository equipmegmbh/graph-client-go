package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	"github.com/verticalgmbh/intelligence-go/pkg/transport"
)

type Client interface {
}

func Connect(ctx context.Context, url string, ssl bool) (Client, error) {
	return connect(ctx, url, !ssl)
}

func connect(ctx context.Context, url string, insecure bool) (Client, error) {
	client := new(defaultClient)
	opts := make([]grpc.DialOption, 0)

	cred := credentials.NewTLS(&tls.Config{})
	tcl := grpc.WithTransportCredentials(cred)

	if opt := grpc.WithInsecure(); insecure {
		tcl = opt
	}

	opts = append(opts, grpc.WithBlock())
	opts = append(opts, tcl)

	conn, err := grpc.DialContext(ctx, url, opts...)

	if f := "dial failed: %v"; err != nil {
		return nil, fmt.Errorf(f, err)
	}

	glog.Info("intelligence client connected")
	glog.Info("start watching intelligence connection")

	go watch(ctx, conn)

	client.cli = transport.NewApiClient(conn)

	return client, nil
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

		if f := "intelligence connection lost"; chg {
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
		glog.Info("intelligence client reconnected")
	}
}
