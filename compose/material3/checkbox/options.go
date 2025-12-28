package checkbox

type CheckboxOptions struct {
	Modifier Modifier
}

type CheckboxOption func(*CheckboxOptions)

func DefaultCheckboxOptions() CheckboxOptions {
	return CheckboxOptions{
		Modifier: EmptyModifier,
	}
}

func WithModifier(m Modifier) CheckboxOption {
	return func(o *CheckboxOptions) {
		o.Modifier = m
	}
}
