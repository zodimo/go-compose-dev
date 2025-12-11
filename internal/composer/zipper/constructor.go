package zipper

func NewComposer(state PersistentState) Composer {

	idManager := GetScopedIdentityManager("composer")
	idManager.ResetKeyCounter()

	return &composer{
		focus:     nil,
		path:      []pathItem{},
		memo:      EmptyMemo,
		state:     state,
		idManager: idManager,
	}
}
