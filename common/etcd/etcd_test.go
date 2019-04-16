package etcd

import (
	"testing"
)

var client = NewStore("localhost:2379")

func TestStore_Put(t *testing.T) {
	if err := client.Put("test", "test etcd put function"); err == nil {
		t.Logf("put ok")
	} else {
		t.Logf("put error: %v", err)
	}
}

func TestStore_Get(t *testing.T) {
	if resp, err := client.Get("test"); err == nil {
		for _, v := range resp.Kvs {
			t.Logf("key: %s", v.Key)
			t.Logf("value: %s", v.Value)
		}
	} else {
		t.Logf("get error: %v", err)
	}

}
