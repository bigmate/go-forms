package forms

type ftype int

const (
	char ftype = 1 << iota
	number
	boolean
	datetime
	array
	charOrNumber = char | number
)

var table = map[ftype]string{
	char:         "String",
	datetime:     "Datetime",
	number:       "Number",
	boolean:      "Boolean",
	array:        "Array",
	charOrNumber: "String or Number",
}

func (ft ftype) String() string {
	return table[ft]
}

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
