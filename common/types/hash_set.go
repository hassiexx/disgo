package types

type setValue struct{}

// StringHashSet is a type for a string hash set.
type StringHashSet struct {
	m map[string]setValue
}

// Add adds a value to the hash set.
func (s *StringHashSet) Add(v string) {
	s.m[v] = setValue{}
}

// Contains checks whether the specified value is in the hash set.
func (s *StringHashSet) Contains(v string) bool {
	_, exists := s.m[v]
	return exists
}

// Remove removes a value from the hash set.
func (s *StringHashSet) Remove(v string) {
	delete(s.m, v)
}
