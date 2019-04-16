package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var err error

type Store struct {
	Client  *clientv3.Client
	Address string
}

func (store *Store) Access() {
	store.Client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{store.Address},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
}

func (store *Store) Put(key, val string, opts ...clientv3.OpOption) error {
	if _, err := store.Client.Put(context.Background(), key, val, opts...); err != nil {
		return err
	} else {
		return nil
	}
}

func (store *Store) Get(key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	if resp, err := store.Client.Get(context.Background(), key, opts...); err != nil {
		return nil, err
	} else {
		return resp, err
	}
}

func (store *Store) Watch(key string, opts ...clientv3.OpOption) {
	rch := store.Client.Watch(context.Background(), key, opts...)
	for resp := range rch {
		for _, ev := range resp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func (store *Store) Delete(key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	if resp, err := store.Client.Delete(context.Background(), key, opts...); err != nil {
		return nil, err
	} else {
		return resp, err
	}
}

func NewStore(address string) *Store {
	store := &Store{
		Address: address,
	}
	store.Access()
	return store
}
