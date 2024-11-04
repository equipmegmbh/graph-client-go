package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/golang/glog"
	"github.com/stretchr/testify/require"
)

const (
	kind = "person"
)

type PersonHex struct {
	Uid        *HexUid `json:"uid,omitempty"`
	GivenName  string  `json:"given-name,omitempty"`
	FamilyName string  `json:"family-name,omitempty"`
}

type Person struct {
	Uid        uint64 `json:"uid,omitempty"`
	Type       string `json:"dgraph.type,omitempty"`
	GivenName  string `json:"given-name,omitempty"`
	FamilyName string `json:"family-name,omitempty"`
}

type Link struct {
	Uid uint64 `json:"uid"`
}

func prepare(t *testing.T, ctx context.Context) Client {
	cli, err := NewClient(ctx, "passthrough:///127.0.0.1:8230", false, "knowledge")
	require.NoError(t, err)
	require.NotNil(t, cli)
	return cli
}

func TestSubscribe(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli := prepare(t, ctx)

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
		require.NoError(t, cli.Update(ctx, kind, entity, reply))
	}()

	glog.Infof("subscribe to events")

	go func() {
		require.NoError(t, cli.Subscribe(ctx, kind, events))
	}()

	glog.Infof("update an entity")

	entities <- &Person{Uid: 10001, GivenName: "John", FamilyName: "Doe"}

	glog.Infof("wait for goroutines to finish...")

	wg.Wait()

	glog.Infof("clean up")
	glog.Infof("close the context to cancel the subscription")
}

// TestExecute initiate some data and test different cases with Execute method
func TestExecute(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cli := prepare(t, ctx)

	var person *Person
	var err error
	// Create person node for further tests
	err = cli.Create(ctx, kind, &Person{GivenName: "John", FamilyName: "Doe"}, &person)
	require.NoError(t, err)
	require.NotNil(t, person)

	query := fmt.Sprintf("{ persons(func: uid(%d)) @filter(type(person)) { uid given-name family-name }}", person.Uid)
	var queryResult struct {
		Persons []*PersonHex `json:"persons"`
	}

	t.Run("SelectQuery", func(t *testing.T) {
		result, transactionOut, err := cli.Execute(ctx, "", query, nil, nil, true)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotEmpty(t, transactionOut)
		require.NoError(t, json.Unmarshal(result, &queryResult))
		require.Len(t, queryResult.Persons, 1)
		require.Equal(t, person.Uid, queryResult.Persons[0].Uid.Uint64)
		require.Equal(t, person.GivenName, queryResult.Persons[0].GivenName)
		require.Equal(t, person.FamilyName, queryResult.Persons[0].FamilyName)
	})

	t.Run("UpdateMutation", func(t *testing.T) {
		person.GivenName = "Jane"
		setMutations := []interface{}{person}
		result, transactionOut, err := cli.Execute(ctx, "", "", setMutations, nil, true)
		require.NoError(t, err)
		require.Nil(t, result)

		result, transactionOut, err = cli.Execute(ctx, "", query, nil, nil, true)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotEmpty(t, transactionOut)

		require.NoError(t, json.Unmarshal(result, &queryResult))
		require.Len(t, queryResult.Persons, 1)
		require.Equal(t, person.Uid, queryResult.Persons[0].Uid.Uint64)
		require.Equal(t, person.GivenName, queryResult.Persons[0].GivenName)
		require.Equal(t, person.FamilyName, queryResult.Persons[0].FamilyName)
	})

	t.Run("DeleteMutation", func(t *testing.T) {
		delMutations := []interface{}{&Link{Uid: person.Uid}}
		result, transactionOut, err := cli.Execute(ctx, "", "", nil, delMutations, true)
		require.NoError(t, err)
		require.Nil(t, result)
		require.NotEmpty(t, transactionOut)

		result, transactionOut, err = cli.Execute(ctx, "", query, nil, nil, true)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NoError(t, json.Unmarshal(result, &queryResult))
		require.Len(t, queryResult.Persons, 0)
	})
}
