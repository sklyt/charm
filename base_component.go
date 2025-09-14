package main

import (
	"fmt"
)

// BaseComponent provides common functionality for all components
type BaseComponent struct {
	id      EntityID
	cType   ComponentType
	visible bool
	style   *Style
}

// NewBaseComponent creates a new base component
func NewBaseComponent(id EntityID, cType ComponentType) *BaseComponent {
	return &BaseComponent{
		id:      id,
		cType:   cType,
		visible: true,
		style:   NewStyle(),
	}
}

// GetID returns the component's entity ID
func (bc *BaseComponent) GetID() EntityID {
	return bc.id
}

// SetID sets the component's entity ID
func (bc *BaseComponent) SetID(id EntityID) {
	bc.id = id
}

// GetType returns the component type
func (bc *BaseComponent) GetType() ComponentType {
	return bc.cType
}

// IsVisible returns whether the component is visible
func (bc *BaseComponent) IsVisible() bool {
	return bc.visible && bc.style.Visible
}

// SetVisible sets the component's visibility
func (bc *BaseComponent) SetVisible(visible bool) {
	bc.visible = visible
}

// GetStyle returns the component's style
func (bc *BaseComponent) GetStyle() *Style {
	return bc.style
}

// SetStyle sets the component's style
func (bc *BaseComponent) SetStyle(style *Style) {
	if style != nil {
		bc.style = style
	}
}

// Validate checks if the component is in a valid state
func (bc *BaseComponent) Validate() error {
	if bc.id == 0 {
		return fmt.Errorf("component has invalid entity ID")
	}
	if bc.cType == ComponentTypeUnknown {
		return fmt.Errorf("component has unknown type")
	}
	return nil
}
