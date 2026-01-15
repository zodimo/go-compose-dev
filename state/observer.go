package state

import (
	"fmt"
	"sync"
)

// singleton
var singletonObserverManager ObserverManager

type ReadObserver func(source StateChangeNotifier)

// NotifyRead notifies the active read observer that a state object has been read.
// Source is typically the state object itself (e.g. *MutableValue).
func NotifyRead(source StateChangeNotifier) {
	singletonObserverManager.NotifyRead(source)
}

// WithReadObserver executes the block with the given read observer.
// It restores the previous observer after the block finishes.
func WithReadObserver(observer ReadObserver, block func()) {
	singletonObserverManager.WithReadObserver(observer, block)
}

type ObserverManager interface {
	WithReadObserver(ReadObserver, func())
	NotifyRead(source StateChangeNotifier)
}

type observerManager struct {
	readObservers []ReadObserver
	mu            sync.RWMutex
	isLocked      bool
}

func (m *observerManager) lock() {
	if m.isLocked {
		fmt.Println("observerManager: is locked")
	}
	m.mu.Lock()
	m.isLocked = true
}
func (m *observerManager) unlock() {
	if !m.isLocked {
		fmt.Println("observerManager: not locked")
		return
	}
	m.mu.Unlock()
	m.isLocked = false
}

func (m *observerManager) pushObserver(observer ReadObserver) {
	m.lock()
	defer m.unlock()
	m.readObservers = append(m.readObservers, observer)
}

func (m *observerManager) popObserver() ReadObserver {
	if len(m.readObservers) == 0 {
		panic("observerManager: no observers")
	}
	observer := m.readObservers[len(m.readObservers)-1]
	m.readObservers = m.readObservers[:len(m.readObservers)-1]
	return observer
}

func (m *observerManager) WithReadObserver(observer ReadObserver, block func()) {
	if observer == nil {
		panic("observerManager: observer cannot be nil")
	}
	if block == nil {
		panic("observerManager: block cannot be nil")
	}
	m.pushObserver(observer)
	defer m.popObserver()
	block()
}

func (m *observerManager) NotifyRead(source StateChangeNotifier) {
	m.mu.RLock()
	if len(m.readObservers) > 0 {
		observer := m.readObservers[len(m.readObservers)-1]
		m.mu.RUnlock()
		observer(source)
		return
	}
	m.mu.RUnlock()
}

func init() {
	singletonObserverManager = &observerManager{}
}
