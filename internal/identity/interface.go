package identity

type IdentityManager interface {
	GenerateID() Identifier
	ResetKeyCounter()
	EmptyIdentifier() Identifier
	private()
}
