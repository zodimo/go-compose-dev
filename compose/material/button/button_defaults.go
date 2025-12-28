package button

func DefaultButtonOptions() ButtonOptions {
	return ButtonOptions{
		Modifier: EmptyModifier,
		Theme:    GetThemeManager().MaterialTheme(),
	}
}
