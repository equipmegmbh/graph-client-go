package graph

import (
	"context"
	"encoding/json"
	"fmt"
	v2 "github.com/equipmegmbh/graph-client-go/pb/v2"
	"github.com/golang/glog"
	"io"

	"github.com/equipmegmbh/graph-client-go/pb/v1"
)

type Client interface {
	Subscribe(ctx context.Context, nodeType string, out chan *Event) error
	SubscribeInSet(ctx context.Context, set, nodeType string, out chan *Event) error

	Query(ctx context.Context, nodeType, query string, out interface{}) error
	QueryFromSet(ctx context.Context, set, nodeType, query string, out interface{}) error

	Create(ctx context.Context, nodeType string, data interface{}, out interface{}) error
	CreateInSet(ctx context.Context, set, nodeType string, data interface{}, out interface{}) error

	Update(ctx context.Context, nodeType string, data interface{}, out interface{}) error
	UpdateInSet(ctx context.Context, set, nodeType string, data interface{}, out interface{}) error

	Delete(ctx context.Context, nodeType string, data interface{}, out interface{}) error
	DeleteInSet(ctx context.Context, set, nodeType string, data interface{}, out interface{}) error

	Execute(ctx context.Context, transactionIn, query string, setMutations, delMutations []interface{}, commit bool) (result []byte, transactionOut string, err error)
	CommitTransaction(ctx context.Context, transaction string) error
	DiscardTransaction(ctx context.Context, transaction string) error
}

type Event struct {
	Set    string
	Type   string
	Action string
	Data   []byte
}

type QueryResponse struct {
	Transaction string
	Data        []byte
}

func NewClient(ctx context.Context, url string, ssl bool, defaultSet string) (Client, error) {
	client := new(defaultClient)
	client.defaultSet = defaultSet

	conn, err := connect(ctx, url, ssl)
	if err != nil {
		return nil, err
	}

	client.v1 = v1.NewApiClient(conn)
	client.v2 = v2.NewDataClient(conn)

	// Start monitoring the connection
	glog.Info("Start watching graph connection")
	go watch(ctx, conn)

	return client, nil
}

type defaultClient struct {
	Client // defaultClient implements the Client interface

	v1         v1.ApiClient
	v2         v2.DataClient
	defaultSet string
}

func (dc *defaultClient) Subscribe(ctx context.Context, nodeType string, out chan *Event) error {
	return dc.SubscribeInSet(ctx, dc.defaultSet, nodeType, out)
}

func (dc *defaultClient) SubscribeInSet(ctx context.Context, set, nodeType string, out chan *Event) error {
	stream, err := dc.v1.Subscribe(ctx, &v1.Subscription{Set: set, Type: nodeType})

	defer close(out)

	if err != nil {
		return err
	}

	for {
		event, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return fmt.Errorf("subscribe failed: %w", err)
		}

		message := new(Event)

		message.Action = ""
		message.Data = event.Response.Data

		switch event.Kind {
		case v1.Event_NOTIFY:
			message.Action = "notify"
		case v1.Event_INSERT:
			message.Action = "insert"
		case v1.Event_UPDATE:
			message.Action = "update"
		case v1.Event_DELETE:
			message.Action = "delete"
		default:
			return fmt.Errorf("unknown event type received")
		}

		out <- message
	}
}

// #region v1

func (dc *defaultClient) Query(ctx context.Context, t, query string, out interface{}) error {
	return dc.QueryFromSet(ctx, dc.defaultSet, t, query, out)
}

func (dc *defaultClient) QueryFromSet(ctx context.Context, set, t, query string, out interface{}) error {
	stream, err := dc.v1.Query(ctx, &v1.Request{Set: set, Type: t, Data: []byte(query)})
	if err != nil {
		return err
	}

	result, buf := make([]byte, 0), make([]byte, 0)
	response := make([][]byte, 0)

	for {
		reply, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("query failed: %w", err)
		}

		response = append(response, reply.Data)
	}

	if len(response) < 1 {
		result = append(result, 0x5B)
		result = append(result, 0x5D)

		goto end
	}

	for i, sep := 0, byte(0x2C); i < len(response); i++ {
		buf = append(buf, response[i]...)
		buf = append(buf, sep)
	}

	result = append(result, 0x5B)
	result = append(result, buf[:len(buf)-1]...)
	result = append(result, 0x5D)

end:
	return json.Unmarshal(result, out)
}

func (dc *defaultClient) Create(ctx context.Context, nodeType string, data interface{}, out interface{}) error {
	return dc.CreateInSet(ctx, dc.defaultSet, nodeType, data, out)
}

func (dc *defaultClient) CreateInSet(ctx context.Context, set, nodeType string, data interface{}, out interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request := &v1.Request{
		Set:  set,
		Type: nodeType,
		Data: d,
	}

	response, err := dc.v1.Create(ctx, request)
	if err != nil {
		return err
	}

	return json.Unmarshal(response.Data, out)
}

func (dc *defaultClient) Update(ctx context.Context, nodeType string, data interface{}, out interface{}) error {
	return dc.UpdateInSet(ctx, dc.defaultSet, nodeType, data, out)
}

func (dc *defaultClient) UpdateInSet(ctx context.Context, set, nodeType string, data interface{}, out interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request := &v1.Request{
		Set:  set,
		Type: nodeType,
		Data: d,
	}

	response, err := dc.v1.Update(ctx, request)
	if err != nil {
		return err
	}

	return json.Unmarshal(response.Data, out)
}

func (dc *defaultClient) Delete(ctx context.Context, nodeType string, data interface{}, out interface{}) error {
	return dc.DeleteInSet(ctx, dc.defaultSet, nodeType, data, out)
}

func (dc *defaultClient) DeleteInSet(ctx context.Context, set, nodeType string, data interface{}, out interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	request := &v1.Request{
		Set:  set,
		Type: nodeType,
		Data: d,
	}

	response, err := dc.v1.Delete(ctx, request)
	if err != nil {
		return err
	}

	return json.Unmarshal(response.Data, out)
}

// #endregion

// #region v2

// Execute a query and/or mutations in a transaction. Set commit to true to commit the transaction with this execution. Otherwise, the transaction will remain open
func (dc *defaultClient) Execute(ctx context.Context, transactionIn, query string, setMutations, delMutations []interface{}, commit bool) (result []byte, transactionOut string, err error) {
	request := &v2.Request{
		Ns:    dc.defaultSet,
		Ttx:   transactionIn,
		Query: query,
	}

	for _, s := range setMutations {
		data, err := json.Marshal(s)
		if err != nil {
			return nil, "", err
		}
		request.Mutations = append(request.Mutations, &v2.Mutation{Set: data})
	}

	for _, d := range delMutations {
		data, err := json.Marshal(d)
		if err != nil {
			return nil, "", err
		}
		request.Mutations = append(request.Mutations, &v2.Mutation{Del: data})
	}

	var response *v2.Response
	response, err = dc.v2.Query(ctx, request)
	if err != nil {
		return
	}

	if commit {
		if _, err = dc.v2.Commit(ctx, &v2.Ttx{Ns: dc.defaultSet, Id: response.Ttx}); err != nil {
			return
		}
	}

	return response.Payload, response.Ttx, nil
}

func (dc *defaultClient) CommitTransaction(ctx context.Context, transaction string) error {
	_, err := dc.v2.Commit(ctx, &v2.Ttx{
		Ns: dc.defaultSet,
		Id: transaction,
	})
	return err
}

func (dc *defaultClient) DiscardTransaction(ctx context.Context, transaction string) error {
	_, err := dc.v2.Discard(ctx, &v2.Ttx{
		Ns: dc.defaultSet,
		Id: transaction,
	})
	return err
}

// #endregion
