package navigationbar

type NavigationBarItemOptions struct {
	Modifier Modifier
}

type NavigationBarItemOption func(*NavigationBarItemOptions)

func ItemWithModifier(m Modifier) NavigationBarItemOption {
	return func(o *NavigationBarItemOptions) {
		o.Modifier = m
	}
}
