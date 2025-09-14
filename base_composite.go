package main

import (
	"fmt"
)

// BaseComposite provides common functionality for composite components
type BaseComposite struct {
	*BaseComponent
	children      []EntityID
	compositeType CompositeType
}

// NewBaseComposite creates a new base composite
func NewBaseComposite(id EntityID, compositeType CompositeType) *BaseComposite {
	return &BaseComposite{
		BaseComponent: NewBaseComponent(id, ComponentTypeComposite),
		children:      make([]EntityID, 0),
		compositeType: compositeType,
	}
}

// GetChildren returns a copy of the children list
func (bc *BaseComposite) GetChildren() []EntityID {
	children := make([]EntityID, len(bc.children))
	copy(children, bc.children)
	return children
}

// AddChild adds a child entity to the composite
func (bc *BaseComposite) AddChild(childID EntityID) error {
	if childID == 0 {
		return fmt.Errorf("cannot add invalid child entity ID")
	}

	if childID == bc.id {
		return fmt.Errorf("cannot add composite to itself")
	}

	// Check for circular references would go here in a more complete implementation

	// Don't add duplicates
	for _, existing := range bc.children {
		if existing == childID {
			return fmt.Errorf("child %d already exists", childID)
		}
	}

	bc.children = append(bc.children, childID)
	return nil
}

// RemoveChild removes a child entity from the composite
func (bc *BaseComposite) RemoveChild(childID EntityID) error {
	for i, child := range bc.children {
		if child == childID {
			// Remove by swapping with last element
			bc.children[i] = bc.children[len(bc.children)-1]
			bc.children = bc.children[:len(bc.children)-1]
			return nil
		}
	}
	return fmt.Errorf("child %d not found", childID)
}

// HasChild checks if a child entity exists in the composite
func (bc *BaseComposite) HasChild(childID EntityID) bool {
	for _, child := range bc.children {
		if child == childID {
			return true
		}
	}
	return false
}

// GetChildCount returns the number of children
func (bc *BaseComposite) GetChildCount() int {
	return len(bc.children)
}

// ClearChildren removes all children
func (bc *BaseComposite) ClearChildren() {
	bc.children = bc.children[:0]
}

// GetCompositeType returns the composite type
func (bc *BaseComposite) GetCompositeType() CompositeType {
	return bc.compositeType
}

// Clone creates a copy of the composite (shallow copy of children list)
func (bc *BaseComposite) Clone() Component {
	clone := &BaseComposite{
		BaseComponent: &BaseComponent{
			id:      bc.id,
			cType:   bc.cType,
			visible: bc.visible,
			style:   bc.style, // Shared reference - deep copy if needed
		},
		compositeType: bc.compositeType,
		children:      make([]EntityID, len(bc.children)),
	}
	copy(clone.children, bc.children)
	return clone
}
