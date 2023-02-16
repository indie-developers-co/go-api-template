package validator

type Validator interface {
	Validate(i interface{}) error
}
