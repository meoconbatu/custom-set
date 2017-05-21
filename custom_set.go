package stringset

const testVersion = 4

type Set []string

func (s Set) String() (output string) {
	output = "{"
	for i, ss := range s {
		output += "\"" + ss + "\""
		if i != len(s)-1 {
			output += ", "
		}
	}
	output += "}"
	return
}
func New() Set {
	s := make(Set, 0)
	return s
}
func NewFromSlice(input Set) (output Set) {
	output = New()
	for _, s := range input {
		output.Add(s)
	}
	return
}
func (s Set) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}
func (s Set) Has(input string) bool {
	for _, ss := range s {
		if ss == input {
			return true
		}
	}
	return false
}
func Subset(s1, s2 Set) bool {
	for _, ss1 := range s1 {
		if !s2.Has(ss1) {
			return false
		}
	}
	return true
}
func Disjoint(s1, s2 Set) bool {
	if len(s1) >= len(s2) {
		for _, ss1 := range s1 {
			if s2.Has(ss1) {
				return false
			}
		}
	} else {
		for _, ss2 := range s2 {
			if s1.Has(ss2) {
				return false
			}
		}
	}
	return true
}
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for _, ss1 := range s1 {
		if !s2.Has(ss1) {
			return false
		}
	}
	return true
}
func (s *Set) Add(input string) {
	if !s.Has(input) {
		*s = append(*s, input)
	}
}
func Intersection(s1, s2 Set) (output Set) {
	for _, ss1 := range s1 {
		if s2.Has(ss1) {
			output = append(output, ss1)
		}
	}
	return output
}
func Difference(s1, s2 Set) (output Set) {
	for _, ss1 := range s1 {
		if !s2.Has(ss1) {
			output = append(output, ss1)
		}
	}
	return output
}
func Union(s1, s2 Set) (output Set) {
	output = Intersection(s1, s2)
	difference1s := Difference(s1, s2)
	difference2s := Difference(s2, s1)
	output = append(output, []string(difference1s)...)
	output = append(output, []string(difference2s)...)
	return
}
