package intelligence

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	"github.com/verticalgmbh/intelligence-go/pb"
)

type ConnectionHandler func()

type Client interface {
	Subscribe(ctx context.Context, set, t string, out chan *Event) error
	Select(ctx context.Context, set, t string, request interface{}, out interface{}) error
	Query(ctx context.Context, set, t, query string, out interface{}) error
	Create(ctx context.Context, set, t string, data interface{}, out interface{}) error
	Update(ctx context.Context, set, t string, data interface{}, out interface{}) error
	Delete(ctx context.Context, set, t string, data interface{}, out interface{}) error
}

type Options struct {
	Secure bool

	ReconnectHandler  ConnectionHandler
	DisconnectHandler ConnectionHandler
}

type Event struct {
	Set    string
	Type   string
	Action string
	Data   []byte
}

func Connect(ctx context.Context, url string, options *Options) (Client, error) {
	return connect(ctx, url, options)
}

func connect(ctx context.Context, url string, options *Options) (Client, error) {
	client := new(defaultClient)
	opts := make([]grpc.DialOption, 0)

	cred := credentials.NewTLS(&tls.Config{})
	tcl := grpc.WithTransportCredentials(cred)

	if opt := grpc.WithInsecure(); !options.Secure {
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

	go watch(ctx, conn, options)

	client.cli = pb.NewApiClient(conn)

	return client, nil
}

func watch(ctx context.Context, conn *grpc.ClientConn, options *Options) {
	if handler := func() {}; options.ReconnectHandler == nil {
		options.ReconnectHandler = handler
	}

	if handler := func() {}; options.DisconnectHandler == nil {
		options.DisconnectHandler = handler
	}

	//noinspection GoUnhandledErrorResult
	defer conn.Close()

	for st, chg := connectivity.Ready, false; ; {
		if !conn.WaitForStateChange(ctx, st) {
			break
		}

		chg = st == connectivity.Ready
		st = conn.GetState()

		if chg {
			glog.Info("intelligence connection lost")
			glog.Info("reconnecting...")

			options.DisconnectHandler()
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
		glog.Info("call reconnect handler")

		options.ReconnectHandler()
	}
}

func reconnected() {
	// default handler, do nothing
}

func disconnected() {
	// default handler, do nothing
}
