// Code generated by options-gen. DO NOT EDIT.
package testcase

import (
	fmt461e464ebed9 "fmt"
	"net/http"
	"time"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptOptionsSetter func(o *Options)

func NewOptions(
	options ...OptOptionsSetter,
) Options {
	o := Options{}

	// Setting defaults from variable
	o.name = defaultOptions.name
	o.timeout = defaultOptions.timeout
	o.maxAttempts = defaultOptions.maxAttempts
	o.httpClient = defaultOptions.httpClient

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithName(opt string) OptOptionsSetter {
	return func(o *Options) {
		o.name = opt
	}
}

func WithTimeout(opt time.Duration) OptOptionsSetter {
	return func(o *Options) {
		o.timeout = opt
	}
}

func WithMaxAttempts(opt int) OptOptionsSetter {
	return func(o *Options) {
		o.maxAttempts = opt
	}
}

func WithHttpClient(opt *http.Client) OptOptionsSetter {
	return func(o *Options) {
		o.httpClient = opt
	}
}

func (o *Options) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("name", _validate_Options_name(o)))
	errs.Add(errors461e464ebed9.NewValidationError("timeout", _validate_Options_timeout(o)))
	errs.Add(errors461e464ebed9.NewValidationError("maxAttempts", _validate_Options_maxAttempts(o)))
	errs.Add(errors461e464ebed9.NewValidationError("httpClient", _validate_Options_httpClient(o)))
	return errs.AsError()
}

func _validate_Options_name(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.name, "required"); err != nil {
		return fmt461e464ebed9.Errorf("field `name` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_timeout(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.timeout, "min=100ms,max=30s"); err != nil {
		return fmt461e464ebed9.Errorf("field `timeout` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_maxAttempts(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.maxAttempts, "min=1,max=10"); err != nil {
		return fmt461e464ebed9.Errorf("field `maxAttempts` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_httpClient(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.httpClient, "gt=0"); err != nil {
		return fmt461e464ebed9.Errorf("field `httpClient` did not pass the test: %w", err)
	}
	return nil
}