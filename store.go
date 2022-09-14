package d7024e

type Store struct {
	kv map[string][]byte //key-value store
	//owner *Node
}

func NewStore() *Store {
	store := &Store{}
	store.kv = make(map[string][]byte)
	return store
}

func (store *Store) get(key string) ([]byte, bool) {
	if val, ok := store.kv[key]; ok {
		return val, true
	}
	return nil, false
}
func (store *Store) add(key string, val []byte) {
	store.kv[key] = val
}
