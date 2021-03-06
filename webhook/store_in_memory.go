package webhook

import "fmt"

type inMemoryStore struct {
	kv map[string][]byte
}

func (store *inMemoryStore) Put(key string, val []byte) error {
	store.kv[key] = val
	return nil
}

func (store *inMemoryStore) Get(key string) ([]byte, error) {
	val, ok := store.kv[key]
	if !ok {
		return nil, fmt.Errorf("inMemoryStoreError: key %s is not found", key)
	}
	return val, nil
}

func (store *inMemoryStore) Delete(key string) error {
	// delete(store.kv[key])
	return nil
}

func (store *inMemoryStore) List(key string) ([]string, error) {
	return nil, nil
}

// NewInMemoryStore returns a new in-memory store
func NewInMemoryStore() Store {
	return &inMemoryStore{
		kv: make(map[string][]byte),
	}
}
