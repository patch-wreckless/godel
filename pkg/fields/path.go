package fields

import (
	"fmt"
)

// Ensure Path is comparable.
var _ = Path{} == Path{}

// A Path represents a field access path. It can include named fields and array
// index references, e.g., "users[0].name".
type Path struct {
	str string
}

func (p Path) Split() []Identifier {

	if len(p.str) == 0 {
		return nil
	}

	type identifierRange struct {
		start int
		end   int
	}

	identifierRanges := make([]identifierRange, 0, len(p.str))

	start := 0
	for i := 1; i < len(p.str); i++ {

		if p.str[i] == '.' {
			r := identifierRange{start: start, end: i}
			identifierRanges = append(identifierRanges, r)
			start = i + 1
			continue
		}

		if p.str[i] == '[' {
			r := identifierRange{start: start, end: i}
			identifierRanges = append(identifierRanges, r)
			start = i
			continue
		}
	}

	lastRange := identifierRange{start: start, end: len(p.str)}
	identifierRanges = append(identifierRanges, lastRange)

	identifiers := make([]Identifier, 0, len(identifierRanges))
	for _, r := range identifierRanges {
		identifierStr := p.str[r.start:r.end]
		identifier, err := ParseIdentifier(identifierStr)
		if err != nil {
			panic(
				fmt.Sprintf("invalid identifier at position %d in %q: %s",
					r.start,
					p.str,
					err.Error()))
		}
		identifiers = append(identifiers, identifier)
	}

	return identifiers
}
