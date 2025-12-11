package divider

func DefaultDividerOptions() DividerOptions {
	return DividerOptions{
		Modifier:  EmptyModifier,
		Thickness: 1,
	}
}
