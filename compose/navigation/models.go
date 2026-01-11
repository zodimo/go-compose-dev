package navigation

import "github.com/zodimo/go-compose/pkg/api"

type RouteableComposable[T any] interface {
	Pattern() string
	HasArgs() bool
	Content() T
	isRoutable()
}

type SimpleRouteabeComposable interface {
	RouteableComposable[api.Composable]
}

type RouteableWithArgsComposable interface {
	RouteableComposable[ComposableWithArgs]
}

var _ SimpleRouteabeComposable = (*simpleRouteable)(nil)

type simpleRouteable struct {
	pattern string
	content api.Composable
}

func newSimpleRouteable(pattern string, content api.Composable) RouteableComposable[api.Composable] {
	return &simpleRouteable{
		pattern: pattern,
		content: content,
	}
}
func (r *simpleRouteable) isRoutable() {}
func (r *simpleRouteable) Pattern() string {
	return r.pattern
}
func (r *simpleRouteable) Content() api.Composable {
	return r.content
}
func (r *simpleRouteable) HasArgs() bool {
	return false
}

func newRouteableWithArguments(pattern string, content ComposableWithArgs) RouteableComposable[ComposableWithArgs] {
	return &routeableWithArgs{
		pattern: pattern,
		content: content,
	}
}

type routeableWithArgs struct {
	pattern string
	content ComposableWithArgs
}

func (r *routeableWithArgs) isRoutable() {}

func (r *routeableWithArgs) Pattern() string {
	return r.pattern
}

func (r *routeableWithArgs) Content() ComposableWithArgs {
	return r.content
}

func (r *routeableWithArgs) HasArgs() bool {
	return true
}
