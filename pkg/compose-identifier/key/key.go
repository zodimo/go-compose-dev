package key

import (
	"fmt"
	"sync/atomic"

	"github.com/zodimo/go-zero-hash/hasher"
)

// What are the properties of the key
// - It is a unique identifier given based on when it was created
// - It is a 32 bit integer

var keyManagerMap map[string]KeyManager

type Key struct {
	value uint32
}

func (k Key) String() string {
	return fmt.Sprintf("%d", k.value)
}

func (k Key) Value() uint32 {
	return k.value
}

type KeyManager interface {
	private()
	GenerateKey() Key
	ResetKeyCounter()
	EmptyKey() Key
	CreateKey(seed string) Key
}

var _ KeyManager = (*keyManager)(nil)

type keyManager struct {
	counter atomic.Uint32
}

func (km *keyManager) private() {}

func (km *keyManager) GenerateKey() Key {
	nextKeyValue := km.counter.Add(1)
	return Key{
		value: nextKeyValue,
	}
}

func (km *keyManager) ResetKeyCounter() {
	km.counter.Store(0)
}

func (km *keyManager) EmptyKey() Key {
	return Key{}
}

func (km *keyManager) CreateKey(seed string) Key {
	return Key{
		value: hasher.HashStringZero(seed),
	}
}

func GetOrCreateKeyManager(id string) KeyManager {
	km, ok := keyManagerMap[id]
	if !ok {
		km = &keyManager{}
		keyManagerMap[id] = km
	}
	return km
}

func init() {
	keyManagerMap = map[string]KeyManager{}
}
