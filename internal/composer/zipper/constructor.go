package zipper

func NewComposer(store PersistentState) Composer {

	idManager := GetScopedIdentityManager("composer")
	idManager.ResetKeyCounter()

	return &composer{
		focus:          nil,
		path:           []pathItem{},
		memo:           EmptyMemo,
		store:          store,
		idManager:      idManager,
		locals:         make(map[interface{}]interface{}),
		providersStack: []map[interface{}]interface{}{},
	}
}
