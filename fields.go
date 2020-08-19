package forms

type Field interface {
	Name() string
	Value() interface{}
	Validate(val interface{}) []string
	Assign(val interface{}) error
}

type field struct {
	name     string
	required bool
	ftype    string
	value    interface{}
	vs       []Validator
}

func (f *field) Value() interface{} {
	return f.value
}

func (f *field) Name() string {
	return f.name
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
