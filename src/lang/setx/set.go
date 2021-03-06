package setx

// Set is the primary interface provided by the mapset package.  It
// represents an unordered set of data and a large number of
// operations that can be applied to that set.
type Set interface {
	// Add Adds an element to the set. Returns whether the item was added.
	Add(i interface{}) bool

	// Cardinality Returns the number of elements in the set.
	Cardinality() int

	// Clear Removes all elements from the set, leaving the empty set.
	Clear()

	// Clone Returns a clone of the set using the same implementation, duplicating all keys.
	Clone() Set

	// Contains Returns whether the given items are all in the set.
	Contains(i ...interface{}) bool

	// Difference Returns the difference between this set
	// and other. The returned set will contain
	// all elements of this set that are not also elements of other.
	//
	// Note that the argument to Difference must be of the same type as the receiver
	// of the method. Otherwise, Difference will panic.
	Difference(other Set) Set

	// Equal Determines if two sets are equal to each other. If they have the same cardinality
	// and contain the same elements, they are considered equal. The order in which
	// the elements were added is irrelevant.
	//
	// Note that the argument to Equal must be of the same type as the receiver of the
	// method. Otherwise, Equal will panic.
	Equal(other Set) bool

	// Intersect Returns a new set containing only the elements that exist only in both sets.
	//
	// Note that the argument to Intersect must be of the same type as the receiver
	// of the method. Otherwise, Intersect will panic.
	Intersect(other Set) Set

	// IsProperSubset Determines if every element in this set is in
	// the other set but the two sets are not equal.
	//
	// Note that the argument to IsProperSubset must be of the same type as the receiver
	// of the method. Otherwise, IsProperSubset will panic.
	IsProperSubset(other Set) bool

	// IsProperSuperset Determines if every element in the other set
	// is in this set but the two sets are not equal.
	//
	// Note that the argument to IsSuperset  must be of the same type as the receiver
	// of the method. Otherwise, IsSuperset will panic.
	IsProperSuperset(other Set) bool

	// IsSubset Determines if every element in this set is in  the other set.
	//
	// Note that the argument to IsSubset must be of the same type as the receiver
	// of the method. Otherwise, IsSubset will panic.
	IsSubset(other Set) bool

	// IsSuperset Determines if every element in the other set is in this set.
	//
	// Note that the argument to IsSuperset must be of the same type as the receiver
	// of the method. Otherwise, IsSuperset will panic.
	IsSuperset(other Set) bool

	// Each Iterates over elements and executes the passed func against each element.
	// If passed func returns true, stop iteration at the time.
	Each(func(interface{}) bool)

	// Iter Returns a channel of elements that you can
	// range over.
	Iter() <-chan interface{}

	// Iterator Returns an Iterator object that you can
	// use to range over the set.
	Iterator() *Iterator

	// Remove a single element from the set.
	Remove(i interface{})

	// String Provides a convenient string representation
	// of the current state of the set.
	String() string

	// SymmetricDifference Returns a new set with all elements which are
	// in either this set or the other set but not in both.
	//
	// Note that the argument to SymmetricDifference
	// must be of the same type as the receiver
	// of the method. Otherwise, SymmetricDifference
	// will panic.
	SymmetricDifference(other Set) Set

	// Union Returns a new set with all elements in both sets.
	//
	// Note that the argument to Union must be of the same type as the receiver of the method.
	// Otherwise, IsSuperset will panic.
	Union(other Set) Set

	// Pop removes and returns an arbitrary item from the set.
	Pop() interface{}

	// PowerSet Returns all subsets of a given set (Power Set).
	PowerSet() Set

	// CartesianProduct Returns the Cartesian Product of two sets.
	CartesianProduct(other Set) Set

	// ToSlice Returns the members of the set as a slice.
	ToSlice() []interface{}
}
