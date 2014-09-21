package invalid

// Field is an interface for an invalidated field
type Field interface {
	Name() string
	Validator() string
	Error() string
}

// Fields is a collection of Field interfaces
type Fields []Field

// Len returns len of f
func (f Fields) Len() int {
	return len(f)
}

// Valid returns true of len is 0
func (f Fields) Valid() bool {
	return f.Len() == 0
}
