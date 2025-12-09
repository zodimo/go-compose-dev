package modifier

type ModifierAwareComposer interface {
	Modifier(func(modifier Modifier) Modifier)
}
