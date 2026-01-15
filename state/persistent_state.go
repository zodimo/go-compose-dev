package state

func WithCompare(compare func(any, any) bool) StateOption {
	if compare == nil {
		panic("WithCompare: compare cannot be nil")
	}
	return func(opts *StateOptions) {
		opts.Compare = compare
	}
}

type PersistentState interface {
	GetState(key string, initial func() any, options ...StateOption) MutableValue
	SetOnStateChange(callback func())
}
