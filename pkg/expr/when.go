package expr

func When[T comparable, U any](value T, conditions map[T]U, otherwise U) U {
	if v, ok := conditions[value]; ok {
		return v
	}
	return otherwise
}
