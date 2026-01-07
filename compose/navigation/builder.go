package navigation

import "github.com/zodimo/go-compose/pkg/api"

// ComposableWithArgs is a function that receives the back stack entry (with arguments)
type ComposableWithArgs func(entry *BackStackEntry) api.Composable

type NavGraphBuilder struct {
	destinations map[string]ComposableWithArgs         // key is the route pattern
	argSpecs     map[string]map[string]NavArgumentSpec // pattern → argName → spec
	patterns     []string                              // ordered list of patterns for matching
}

func NewNavGraphBuilder() *NavGraphBuilder {
	return &NavGraphBuilder{
		destinations: make(map[string]ComposableWithArgs),
		argSpecs:     make(map[string]map[string]NavArgumentSpec),
		patterns:     []string{},
	}
}

// Composable adds a destination without arguments (backward compatible)
func (b *NavGraphBuilder) Composable(route string, content api.Composable) {
	b.ComposableWithArgs(route, func(_ *BackStackEntry) api.Composable {
		return content
	})
}

// ComposableWithArgs adds a destination that receives the back stack entry
func (b *NavGraphBuilder) ComposableWithArgs(route string, content ComposableWithArgs) {
	b.destinations[route] = content
	b.patterns = append(b.patterns, route)
}

// Argument registers an argument spec for a route pattern
// Example: b.Argument("details/{itemId}", "itemId", NewNavArgumentSpec(NavTypeString))
func (b *NavGraphBuilder) Argument(route, name string, spec NavArgumentSpec) {
	if b.argSpecs[route] == nil {
		b.argSpecs[route] = make(map[string]NavArgumentSpec)
	}
	b.argSpecs[route][name] = spec
}

// findDestination matches a route against registered patterns
// Returns the composable, extracted arguments, and whether a match was found
func (b *NavGraphBuilder) findDestination(route string) (ComposableWithArgs, NavArguments, bool) {
	for _, pattern := range b.patterns {
		args, matched := matchRoute(pattern, route)
		if matched {
			composable := b.destinations[pattern]

			// Apply default values from arg specs for missing optional args
			if specs, ok := b.argSpecs[pattern]; ok {
				for argName, spec := range specs {
					if _, exists := args[argName]; !exists && spec.IsDefaultValuePresent {
						if spec.DefaultValue.IsSome() {
							args[argName] = spec.DefaultValue.UnwrapUnsafe()
						}
					}
				}
			}

			return composable, args, true
		}
	}
	return nil, nil, false
}
