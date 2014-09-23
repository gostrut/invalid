package invalid

import "testing"
import "github.com/nowk/assert"

type tField struct {
	name      string
	validator string
}

func (f tField) Name() string {
	return f.name
}

func (f tField) Validator() string {
	return f.validator
}

func (f tField) Error() string {
	return ""
}

func TestFields(t *testing.T) {
	a := tField{name: "one"}
	b := tField{name: "two"}
	c := tField{name: "two"}
	d := tField{name: "three"}

	fields := NewFields()
	fields.Add("one", a)
	fields.Add("two", b, c)
	fields.Add("three", d)

	for _, v := range []struct {
		s string
		f []Field
	}{
		{"one", []Field{a}},
		{"two", []Field{b, c}},
		{"three", []Field{d}},
		{"four", nil},
	} {
		field := fields.Get(v.s)
		assert.Equal(t, field, v.f)
	}
}
