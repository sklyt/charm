package main

type SparseSet[T any] struct {
	sparse []uint32   // Maps entity IDs to dense array indices
	dense  []EntityID // Packed array of entity IDs
	data   []T        // Component data corresponding to dense entities
	size   int        // Current number of entities
}

func NewSparseSet[T any](capacity int) *SparseSet[T] {
	return &SparseSet[T]{
		sparse: make([]uint32, capacity),
		dense:  make([]EntityID, 0, capacity),
		data:   make([]T, 0, capacity),
		size:   0,
	}
}

func (s *SparseSet[T]) Has(entityID EntityID) bool {
	id := uint32(entityID)
	if id >= uint32(len(s.sparse)) {
		return false
	}

	denseIndex := s.sparse[id]
	return denseIndex < uint32(s.size) && s.dense[denseIndex] == entityID
}

func (s *SparseSet[T]) Get(entityID EntityID) (T, bool) {
	var zero T

	if !s.Has(entityID) {
		return zero, false
	}

	denseIndex := s.sparse[entityID]
	return s.data[denseIndex], true
}

func (s *SparseSet[T]) Set(entityID EntityID, component T) {
	id := uint32(entityID)

	// Grow sparse array if needed
	if id >= uint32(len(s.sparse)) {
		newSize := max(uint32(len(s.sparse)*2), id+1)
		newSparse := make([]uint32, newSize)
		copy(newSparse, s.sparse)
		s.sparse = newSparse
	}

	if s.Has(entityID) {
		// Update existing component
		denseIndex := s.sparse[id]
		s.data[denseIndex] = component
	} else {
		// Add new component
		s.sparse[id] = uint32(s.size)
		s.dense = append(s.dense, entityID)
		s.data = append(s.data, component)
		s.size++
	}
}

func (s *SparseSet[T]) Remove(entityID EntityID) bool {
	if !s.Has(entityID) {
		return false
	}

	id := uint32(entityID)
	denseIndex := s.sparse[id]
	lastIndex := uint32(s.size - 1)

	// Swap with last element
	if denseIndex != lastIndex {
		lastEntity := s.dense[lastIndex]
		s.dense[denseIndex] = lastEntity
		s.data[denseIndex] = s.data[lastIndex]
		s.sparse[lastEntity] = denseIndex
	}

	// Shrink arrays
	s.dense = s.dense[:s.size-1]
	s.data = s.data[:s.size-1]
	s.size--

	return true
}

func (s *SparseSet[T]) Clear() {
	s.dense = s.dense[:0]
	s.data = s.data[:0]
	s.size = 0
}

func (s *SparseSet[T]) Size() int {
	return s.size
}

func (s *SparseSet[T]) Iterate(fn func(EntityID, T) bool) {
	for i := 0; i < s.size; i++ {
		if !fn(s.dense[i], s.data[i]) {
			break
		}
	}
}

func (s *SparseSet[T]) GetEntities() []EntityID {
	entities := make([]EntityID, s.size)
	copy(entities, s.dense)
	return entities
}
