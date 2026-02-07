package fields

import (
	"slices"
)

// A Path represents a field access path.
type Path struct {
	segments []PathSegment
}

// NewPath initializes a [Path] from the given [PathSegment] values.
func NewPath(segments []PathSegment) Path {
	return Path{segments: segments}
}

// Segments returns the individual [PathSegment] values of the [Path].
func (p Path) Segments() []PathSegment {
	if len(p.segments) == 0 {
		return nil
	}
	return slices.Clone(p.segments)
}
