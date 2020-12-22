package forms

type Field interface {
	Name() string
	Value() interface{}
	Validate(val interface{}) []string
	Assign(val interface{}) error
	Bound() bool
}

type field struct {
	name     string
	required bool
	ftype    string
	bound    bool
	value    interface{}
	vs       []Validator
}

func (f *field) Value() interface{} {
	return f.value
}

func (f *field) Name() string {
	return f.name
}

func (f *field) Bound() bool {
	return f.bound
}

func (f *field) runValidators(errors []string) []string {
	for _, validator := range f.vs {
		var err = validator(f.value)
		if err != nil {
			errors = append(errors, err.Error())
		}
	}
	return errors
}
