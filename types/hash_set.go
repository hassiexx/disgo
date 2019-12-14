package types

// HashSetValue is an empty hash set value struct.
type hashSetValue struct{}

// StringHashSet is a type for a string hash set.
type StringHashSet struct {
	m map[string]hashSetValue
}

// NewStringHashSet creates a new string hash set.
func NewStringHashSet() *StringHashSet {
	return &StringHashSet{
		m: make(map[string]hashSetValue),
	}
}

// Add adds a value to the hash set.
func (s *StringHashSet) Add(v string) {
	s.m[v] = hashSetValue{}
}

// Contains checks whether the specified value is in the hash set.
func (s *StringHashSet) Contains(v string) bool {
	_, exists := s.m[v]
	return exists
}

// IsEmpty checks whether the hash set is empty.
func (s *StringHashSet) IsEmpty() bool {
	return len(s.m) == 0
}

// Remove removes a value from the hash set.
func (s *StringHashSet) Remove(v string) {
	delete(s.m, v)
}

// Values gets the values from the hash set.
func (s *StringHashSet) Values() []string {
	var values []string
	for k, _ := range s.m {
		values = append(values, k)
	}
	return values
}
