package etcd

import (
	"testing"
	"time"

	"github.com/docker/libkv/store"
	"github.com/docker/libkv/testutils"
)

func makeEtcdClient(t *testing.T) store.Store {
	client := "localhost:4001"

	kv, err := New(
		[]string{client},
		&store.Config{
			ConnectionTimeout: 3 * time.Second,
		},
	)

	if err != nil {
		t.Fatalf("cannot create store: %v", err)
	}

	return kv
}

func TestEtcdStore(t *testing.T) {
	kv := makeEtcdClient(t)
	backup := makeEtcdClient(t)

	testutils.RunTestCommon(t, kv)
	testutils.RunTestAtomic(t, kv)
	testutils.RunTestWatch(t, kv)
	testutils.RunTestLock(t, kv)
	testutils.RunTestTTL(t, kv, backup)
	testutils.RunCleanup(t, kv)
}
