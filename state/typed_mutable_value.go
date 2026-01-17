package state

/*
// Broken implementation - commented out to fix build

type mutableValueTyped[T any] struct {
	cell           T
	changeNotifier func(T)
	compare        func(T, T) bool
}

func NewMutableValueTyped[T any](initial T, changeNotifier func(T), compare func(T, T) bool) TypedMutableValue[T] {
	return &mutableValueTyped[T]{
		cell:           initial,
		changeNotifier: changeNotifier,
		compare:        compare,
	}
}

type MutableValueTypedWrapper[T any] struct {
	mv MutableValue
}

func (w *MutableValueTypedWrapper[T]) Get() T {
	return w.mv.Get().(T)
}

func (w *MutableValueTypedWrapper[T]) Set(value T) {
	w.mv.Set(value)
}

func (w *MutableValueTypedWrapper[T]) Subscribe(callback func()) Subscription {
	return w.mv.Subscribe(callback)
}

func WrapMutableValue[T any](mv MutableValue) (TypedMutableValue[T], error) {

	mvTyped, ok := mv.(*mutableValueTyped[T])
	if !ok {
		return nil, fmt.Errorf("mutable value is not of type %T", mvTyped)
	}

	_, ok = mvTyped.cell.(T)
	if !ok {
		var zero T
		return nil, fmt.Errorf("cell is not of type %T, got %T", zero, mv.cell)
	}

	return &MutableValueTypedWrapper[T]{
		mv: mv,
	}, nil
}

func (w *MutableValueTypedWrapper[T]) Unwrap() MutableValueInterface {
	return w.mv
}
*/
