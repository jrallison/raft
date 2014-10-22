package raft

// StateMachine is the interface for allowing the host application to save and
// recovery the state machine. This makes it possible to make snapshots
// and compact the log.
type StateMachine interface {
	Save() ([]byte, error)
	Recovery([]byte) error
}

// PITStateMachine is the interface for allowing the host application to save and
// recovery the state machine from a specific point (index/term) in time.
type PITStateMachine interface {
	SaveAt(index, term uint64) ([]byte, error)
}
