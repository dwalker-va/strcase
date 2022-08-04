package strcase

import "sync"

var defaultConverter = NewConverter()

// NewConverter returns a Converter object that can be configured without relying on global package state.
// This is most useful if you need to convert multiple strings in the same program with different customizations for the same acronym.
// This object is thread-safe.
func NewConverter() *Converter {
	return &Converter{
		uppercaseAcronym: map[string]string{
			"ID": "id",
		},
	}
}

type Converter struct {
	lock             sync.RWMutex
	uppercaseAcronym map[string]string
}
