package types

type setValue struct{}

// StringSet is a type for a string set.
type StringSet map[string]setValue

// Add adds a value into the set.
func (s StringSet) Add(v string) {
	s[v] = setValue{}
}

// Contains checks whether the value is in the set.
func (s StringSet) Contains(v string) bool {
	_, exists := s[v]
	return exists
}

// Remove removes a value from the set.
func (s StringSet) Remove(v string) {
	delete(s, v)
}
