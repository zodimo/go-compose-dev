package compose

import (
	"go-compose-dev/internal/composer/zipper"
	"go-compose-dev/internal/state"
	"go-compose-dev/pkg/api"
)

type Composable = api.Composable
type Composer = api.Composer

func NewComposer(store state.PersistentState) Composer {
	return zipper.NewComposer(store)
}

type composableSequence struct {
	contents []Composable
}

func NewComposableSequence(composables []Composable) *composableSequence {
	return &composableSequence{contents: composables}
}

func (s *composableSequence) Compose(c Composer) Composer {
	for _, content := range s.contents {
		c = content(c)
	}
	return c
}

func Sequence(contents ...Composable) Composable {
	return func(c Composer) Composer {
		return NewComposableSequence(contents).Compose(c)
	}
}
