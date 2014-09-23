package invalid

// Field is an interface for an invalidated field
type Field interface {
	Name() string
	Validator() string
	Error() string
}

// Fields type is a map of []Field to a string (the field name)
type Fields map[string][]Field

// NewFields returns an initialized Fields type
func NewFields() Fields {
	return make(map[string][]Field)
}

// Len returns len of f
func (f Fields) Len() int {
	return len(f)
}

// Valid returns true of len is 0
func (f Fields) Valid() bool {
	return f.Len() == 0
}

// Get returns invalid fields by it's mapped name
func (f Fields) Get(name string) []Field {
	return f[name]
}

// Add appends invalid fields to a mapped name
func (f Fields) Add(name string, field ...Field) {
	f[name] = append(f[name], field...)
}
