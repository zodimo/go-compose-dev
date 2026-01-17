package flow

// BufferOverflow represents the strategy to handle buffer overflow in SharedFlow.
type BufferOverflow int

const (
	// BufferOverflowSuspend indicates that the sender should suspend on buffer overflow.
	BufferOverflowSuspend BufferOverflow = iota
	// BufferOverflowDropOldest indicates that the sender should drop the oldest value on buffer overflow.
	BufferOverflowDropOldest
	// BufferOverflowDropLatest indicates that the sender should drop the latest value on buffer overflow.
	BufferOverflowDropLatest
)
