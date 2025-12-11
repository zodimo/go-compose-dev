package models

import (
	"go-compose-dev/pkg/compose-identifier/key"
)

type Identifier struct {
	key key.Key
}

func (i Identifier) Value() uint32 {
	return i.key.Value()
}
func (i Identifier) String() string {
	return i.key.String()
}

func NewIdentifier() *Identifier {
	keyManager := key.GetOrCreateKeyManager("default")
	return &Identifier{
		key: keyManager.GenerateKey(),
	}
}

func NewScopedIdentifier(scope string) *Identifier {
	keyManager := key.GetOrCreateKeyManager(scope)
	return &Identifier{
		key: keyManager.GenerateKey(),
	}
}

type IdentityManager struct {
	scope string
}

func GetOrCreateIdentityManager(scope string) *IdentityManager {
	return &IdentityManager{
		scope: scope,
	}
}

func (im IdentityManager) GenerateKey() *Identifier {
	return NewScopedIdentifier(im.scope)
}
func (im IdentityManager) ResetKeyCounter() {
	im.getKeyManager().ResetKeyCounter()
}
func (im IdentityManager) EmptyIdentifier() *Identifier {
	return &Identifier{key: im.getKeyManager().EmptyKey()}
}

func (im IdentityManager) getKeyManager() key.KeyManager {
	return key.GetOrCreateKeyManager(im.scope)
}
