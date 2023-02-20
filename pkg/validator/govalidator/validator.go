package govalidator

import (
	"encoding/json"
	"errors"

	govalidator "github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"gitlab.com/indie-developers/go-api-echo-template/pkg/validator"
)

const (
	MessageError = "request failed with validations errors"
)

type ExclusiveErrorResponse struct {
	Message string      `json:"message"`
	Field   string      `json:"field"`
	Tag     string      `json:"tag"`
	Value   interface{} `json:"value"`
}

type ErrorResponse struct {
	Errors  []ExclusiveErrorResponse `json:"errors"`
	Message string                   `json:"message"`
}

func (e ErrorResponse) Error() string {
	jsonBody, _ := json.Marshal(e.Errors)
	return string(jsonBody)
}

type RequestValidator struct {
	validator *govalidator.Validate
}

func NewRequestValidator() validator.Validator {
	validate := govalidator.New()
	validate.RegisterValidation("notblank", validators.NotBlank)

	return &RequestValidator{validate}
}

func (cv *RequestValidator) Validate(i interface{}) error {
	if errs := cv.validator.Struct(i); errs != nil {
		var validationErrors govalidator.ValidationErrors
		errors.As(errs, &validationErrors)
		return createErrorResponse(validationErrors)
	}
	return nil
}

func createErrorResponse(validationErrors govalidator.ValidationErrors) ErrorResponse {
	var errorResponses []ExclusiveErrorResponse
	for _, err := range validationErrors {
		errorResponses = append(errorResponses, ExclusiveErrorResponse{
			Message: "Field validation failed",
			Field:   err.StructNamespace(),
			Tag:     err.ActualTag(),
			Value:   err.Value(),
		})
	}

	return ErrorResponse{
		Errors:  errorResponses,
		Message: MessageError,
	}
}
