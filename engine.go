package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type EntityID uint32

type ComponentType uint8

const (
	ComponentTypeUnknown ComponentType = iota
	ComponentTypeList
	ComponentTypeInput
	ComponentTypeButton
	ComponentTypeText
	ComponentTypeComposite
)

type Engine struct {
	mu           sync.RWMutex
	sparseSet    *SparseSet[any]
	nextEntityID uint32
	activeRoot   EntityID
	lastError    string
	initialized  bool
}

var (
	engineInstance *Engine
	engineOnce     sync.Once
)

func GetEngineInstance() *Engine {
	engineOnce.Do(func() {
		engineInstance = &Engine{
			sparseSet:    NewSparseSet[any](1024), // Start with capacity for 1024 entities
			nextEntityID: 1,                       // 0 is reserved for "null entity"
		}
	})
	return engineInstance
}

func (e *Engine) Initialize() int {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.initialized {
		e.setError("Engine already initialized")
		return -1
	}

	e.initialized = true
	e.lastError = ""
	return 0
}

func (e *Engine) CreateEntity() EntityID {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.initialized {
		e.setError("Engine not initialized")
		return 0
	}

	entityID := EntityID(atomic.AddUint32(&e.nextEntityID, 1))
	return entityID
}

func (e *Engine) DestroyEntity(entityID EntityID) int {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.initialized {
		e.setError("Engine not initialized")
		return -1
	}

	if !e.sparseSet.Has(entityID) {
		e.setError(fmt.Sprintf("Entity %d does not exist", entityID))
		return -1
	}

	e.sparseSet.Remove(entityID)

	// If this was the active root, clear it
	if e.activeRoot == entityID {
		e.activeRoot = 0
	}

	return 0
}

func (e *Engine) SetActiveRoot(entityID EntityID) int {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.initialized {
		e.setError("Engine not initialized")
		return -1
	}

	if entityID != 0 && !e.sparseSet.Has(entityID) {
		e.setError(fmt.Sprintf("Entity %d does not exist", entityID))
		return -1
	}

	e.activeRoot = entityID
	return 0
}

func (e *Engine) GetActiveRoot() EntityID {
	e.mu.RLock()
	defer e.mu.RUnlock()
	return e.activeRoot
}

func (e *Engine) AddComponent(entityID EntityID, component any) int {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.initialized {
		e.setError("Engine not initialized")
		return -1
	}

	e.sparseSet.Set(entityID, component)
	return 0
}

func (e *Engine) GetComponent(entityID EntityID) (any, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if !e.initialized {
		return nil, false
	}

	return e.sparseSet.Get(entityID)
}

func (e *Engine) HasComponent(entityID EntityID) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if !e.initialized {
		return false
	}

	return e.sparseSet.Has(entityID)
}

func (e *Engine) Shutdown() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.sparseSet.Clear()
	e.activeRoot = 0
	e.lastError = ""
	e.initialized = false
}

func (e *Engine) setError(msg string) {
	e.lastError = msg
}
