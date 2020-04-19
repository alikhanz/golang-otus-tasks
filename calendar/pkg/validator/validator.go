package validator

type Validator struct {
	errors Errors
}

func NewValidator() Validator {
	return Validator{}
}

func (v *Validator) HasErrors() bool {
	return len(v.errors.errors) > 0
}

func (v *Validator) AddError(err string) {
	v.errors.Add(err)
}

func (v *Validator) RenderErrors() string {
	return v.errors.Render()
}

func (v *Validator) ValidateString(value string, rule StringRule) {
	if !rule.Validate(value) {
		v.errors.Add(rule.ErrorMessage())
	}
}
