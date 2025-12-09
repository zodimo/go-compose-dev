package state

import (
	"go-compose-dev/internal/immap"
)

// Immutable MAP, not related to the Slots in tree/map.go
var EmptyMemo = immap.EmptyImmutableMapAny

type Memo = immap.ImmutableMap[any]
