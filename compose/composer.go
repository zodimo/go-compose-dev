package compose

import (
	"github.com/zodimo/go-compose/internal/composer/zipper"
	"github.com/zodimo/go-compose/internal/sequence"
	"github.com/zodimo/go-compose/pkg/api"
	"github.com/zodimo/go-compose/state"
)

type Composable = api.Composable
type Composer = api.Composer

func NewComposer(store state.PersistentState) Composer {
	return zipper.NewComposer(store)
}

var Sequence = sequence.Sequence

func Id() Composable {
	return func(c Composer) Composer {
		return c
	}
}
