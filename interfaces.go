package main

type Component interface {
	GetType() ComponentType
	GetID() EntityID
	SetID(EntityID)
	IsVisible() bool
	SetVisible(bool)
	GetStyle() *Style
	SetStyle(*Style)
	Clone() Component
	Validate() error
}

type Composite interface {
	Component
	GetChildren() []EntityID
	AddChild(EntityID) error
	RemoveChild(EntityID) error
	HasChild(EntityID) bool
	GetChildCount() int
	ClearChildren()
	GetCompositeType() CompositeType
}

type CompositeType uint8

const (
	CompositeTypeUnknown CompositeType = iota
	CompositeTypeBox
	CompositeTypeScroll
	CompositeTypeTabs
	CompositeTypePopup
	CompositeTypeCollapsible
	CompositeTypeTree
)

type Focusable interface {
	IsFocused() bool
	SetFocused(bool)
	CanFocus() bool
}

type Updatable interface {
	Update(msg any) (Component, error)
}

type Renderable interface {
	Render() string
}

type Style struct {
	Width           int
	Height          int
	MarginTop       int
	MarginRight     int
	MarginBottom    int
	MarginLeft      int
	PaddingTop      int
	PaddingRight    int
	PaddingBottom   int
	PaddingLeft     int
	BackgroundColor string
	TextColor       string
	BorderStyle     string
	BorderColor     string
	Visible         bool
}

func NewStyle() *Style {
	return &Style{
		Width:           -1, // -1 means auto-size
		Height:          -1,
		BackgroundColor: "",
		TextColor:       "",
		BorderStyle:     "",
		BorderColor:     "",
		Visible:         true,
	}
}
