package fields

import (
	"strconv"
)

// An IndexExpr is a [PathSegment] referencing an item in an ordered list.
type IndexExpr uint64

// String formats the [IndexExpr] into the bracket-notation used for indexing.
func (i IndexExpr) String() string {
	return "[" + strconv.FormatUint(uint64(i), 10) + "]"
}
