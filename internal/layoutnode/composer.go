package layoutnode

type TreeBuilderComposer interface {
	StartBlock(id string) TreeBuilderComposer
	EndBlock() TreeBuilderComposer
	Build() TreeNode
}
