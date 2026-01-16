// Package runtime provides core composition runtime interfaces ported from Jetpack Compose.
package runtime

// ComposeNodeLifecycleCallback observes lifecycle of nodes emitted with ComposeNode
// inside ReusableContentHost and ReusableContent.
//
// The ReusableContentHost introduces the concept of reusing (or recycling) nodes,
// as well as deactivating parts of composition, while keeping the nodes around to
// reuse common structures in the next iteration.
//
// In this state, RememberObserver is not sufficient to track lifetime of data
// associated with reused node, as deactivated or reused parts of composition is disposed.
//
// These callbacks track intermediate states of the node in reusable groups for
// managing data contained inside reusable nodes or associated with them.
//
// Important: the runtime only supports node implementation of this interface.
type ComposeNodeLifecycleCallback interface {
	// OnReuse is invoked when the node was reused in the composition.
	// Consumers might use this callback to reset data associated with the
	// previous content, as it is no longer valid.
	OnReuse()

	// OnDeactivate is invoked when the group containing the node was deactivated.
	// This happens when the content of ReusableContentHost is deactivated.
	//
	// The node will not be reused in this recompose cycle, but might be reused
	// or released in the future. Consumers might use this callback to release
	// expensive resources or stop continuous processes that were dependent on
	// the node being used in composition.
	//
	// If the node is reused, OnReuse will be called again to prepare the node
	// for reuse. Similarly, OnRelease will indicate that deactivated node will
	// never be reused again.
	OnDeactivate()

	// OnRelease is invoked when the node exits the composition entirely and
	// won't be reused again. All intermediate data related to the node can
	// be safely disposed.
	OnRelease()
}
