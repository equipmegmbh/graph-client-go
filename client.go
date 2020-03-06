package intelligence

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/verticalgmbh/intelligence-go/pb"
)

type defaultClient struct {
	cli pb.ApiClient
}

func (dc *defaultClient) Subscribe(ctx context.Context, set, t string, out chan *Event) error {
	stream, err := dc.cli.Subscribe(ctx, &pb.Subscription{Set: set, Type: t})

	defer close(out)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	for {
		event, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if f := "subscribe failed: %v"; err != nil {
			return fmt.Errorf(f, err)
		}

		message := new(Event)

		message.Action = ""
		message.Data = event.Response.Data

		switch event.Kind {
		case pb.Event_NOTIFY:
			message.Action = "notify"
		case pb.Event_INSERT:
			message.Action = "insert"
		case pb.Event_UPDATE:
			message.Action = "update"
		case pb.Event_DELETE:
			message.Action = "delete"
		default:
			return fmt.Errorf("unknown event type received")
		}

		out <- message
	}
}

func (dc *defaultClient) Select(ctx context.Context, set, t string, request interface{}, out interface{}) error {
	d, err := json.Marshal(request)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	stream, err := dc.cli.Select(ctx, &pb.Request{Set: set, Type: t, Data: d})

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	result, buf := make([]byte, 0), make([]byte, 0)
	response := make([][]byte, 0)

	for {
		reply, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if f := "select failed: %v"; err != nil {
			return fmt.Errorf(f, err)
		}

		response = append(response, reply.Data)
	}

	if lb, rb := byte(0x5B), byte(0x5D); len(response) < 1 {
		result = append(result, lb)
		result = append(result, rb)

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

func (dc *defaultClient) Query(ctx context.Context, set, t, query string, out interface{}) error {
	stream, err := dc.cli.Query(ctx, &pb.Request{Set: set, Type: t, Data: []byte(query)})

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	result, buf := make([]byte, 0), make([]byte, 0)
	response := make([][]byte, 0)

	for {
		reply, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if f := "query failed: %v"; err != nil {
			return fmt.Errorf(f, err)
		}

		response = append(response, reply.Data)
	}

	if lb, rb := byte(0x5B), byte(0x5D); len(response) < 1 {
		result = append(result, lb)
		result = append(result, rb)

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

func (dc *defaultClient) Create(ctx context.Context, set, t string, data interface{}, out interface{}) error {
	d, err := json.Marshal(data)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	request := &pb.Request{
		Set:  set,
		Type: t,
		Data: d,
	}

	response, err := dc.cli.Create(ctx, request)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	return json.Unmarshal(response.Data, out)
}

func (dc *defaultClient) Update(ctx context.Context, set, t string, data interface{}, out interface{}) error {
	d, err := json.Marshal(data)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	request := &pb.Request{
		Set:  set,
		Type: t,
		Data: d,
	}

	response, err := dc.cli.Update(ctx, request)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	return json.Unmarshal(response.Data, out)
}

func (dc *defaultClient) Delete(ctx context.Context, set, t string, data interface{}, out interface{}) error {
	d, err := json.Marshal(data)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	request := &pb.Request{
		Set:  set,
		Type: t,
		Data: d,
	}

	response, err := dc.cli.Delete(ctx, request)

	if f := "%v"; err != nil {
		return fmt.Errorf(f, err)
	}

	return json.Unmarshal(response.Data, out)
}
