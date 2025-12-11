package mswitch

type HandlerWrapper struct {
	Func func(bool)
}
