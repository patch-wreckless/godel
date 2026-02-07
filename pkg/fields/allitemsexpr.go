package fields

// AllItemsToken is a special index expression referring to every item in an
// ordered list.
const AllItemsToken string = "[*]"

// AllItemsExpr is a [PathSegment] representing the [AllItemsToken].
type AllItemsExpr struct{}

// String returns the [AllItemsToken].
func (a AllItemsExpr) String() string {
	return AllItemsToken
}
