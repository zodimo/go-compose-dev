package modifier

// Then is a helper function to chain modifiers.
func Then(modifier Modifier, other Modifier) Modifier {
	if modifier == nil || modifier == EmptyModifier {
		return other
	}
	if other == nil || other == EmptyModifier {
		return modifier
	}
	return modifier.Then(other)
}
