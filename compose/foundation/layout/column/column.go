package column

func DefaultColumnOptions() ColumnOptions {
	return ColumnOptions{
		Modifier:  EmptyModifier,
		Spacing:   SpaceEnd, // 0
		Alignment: Start,    // 0
	}
}

func Column(content Composable, options ...ColumnOption) Composable {
	opts := DefaultColumnOptions()
	for _, option := range options {
		option(&opts)
	}
	return func(c Composer) Composer {
		c.StartBlock("Column")
		c.Modifier(func(modifier Modifier) Modifier {
			return modifier.Then(opts.Modifier)
		})
		return c.EndBlock()
	}
}
