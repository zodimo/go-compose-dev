package zipper

// // Tree Builder -> Node operations
// type Composer interface {
// 	//Tree Builder operations
// 	StartBlock(id string) Composer
// 	EndBlock() Composer
// 	Build() RootNode

// 	EmotSlot(k string, v any) Composer // slot is a property on the layoutNode

// 	// Tree navigation
// 	up() Composer

// 	GenerateID() Identifier
// 	GetID() Identifier
// 	GetPath() NodePath
// 	Modifier(modifier Modifier) Composer
// 	ModifierThen(modifier Modifier) Composer

// 	// --  state
// 	state.StatefulComposer
// }

type Composer interface {
	ApiComposer
}
