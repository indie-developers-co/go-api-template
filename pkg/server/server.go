package server

import "gitlab.com/indie-developers/go-api-echo-template/pkg/validator"

type Server interface {
	Run(address string)
	Debug(debug bool)
	SetRequestValidator(validator validator.Validator)
}
