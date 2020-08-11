package intelligence

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/golang/glog"
	"github.com/stretchr/testify/require"
)

const (
	set  = "knowledge"
	kind = "person"
)

type Person struct {
	Uid      uint64 `json:"uid,omitempty"`
	Forename string `json:"given-name,omitempty"`
	Surname  string `json:"family-name,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Age      uint64 `json:"age,omitempty"`
}

func TestSubscribe(t *testing.T) {
	options := new(Options)

	options.Secure = false

	ctx, cancel := context.WithCancel(context.Background())
	cli, err := Connect(ctx, "passthrough:///127.0.0.1:8230", options)

	defer cancel()

	require.NoError(t, err)
	require.NotNil(t, cli)

	events := make(chan *Event)
	entities := make(chan interface{})

	defer close(entities)

	person := new(Person)
	reply := new(Person)
	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		defer wg.Done()

		event, open := <-events

		glog.Infof("event received: %v", event)

		require.True(t, open)
		require.NoError(t, json.Unmarshal(event.Data, person))
		require.Equal(t, uint64(10001), person.Uid)
	}()

	go func() {
		defer wg.Done()

		entity, open := <-entities

		require.True(t, open)
		require.NoError(t, cli.Update(ctx, set, kind, entity, reply))
	}()

	glog.Infof("subscribe to events")

	go func() {
		require.NoError(t, cli.Subscribe(ctx, set, kind, events))
	}()

	glog.Infof("update an entity")

	entities <- &Person{Uid: 10001, Age: 99}

	glog.Infof("wait for goroutines to finish...")

	wg.Wait()

	glog.Infof("clean up")
	glog.Infof("close the context to cancel the subscription")
}
