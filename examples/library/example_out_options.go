// Code generated by options-gen. DO NOT EDIT.
package main

import (
	fmt461e464ebed9 "fmt"

	subpackage "github.com/kazhuravlev/options-gen/examples/library/sub-package"
	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptOptionsSetter func(o *Options)

func NewOptions(
	service1 *subpackage.Service1,
	s3Endpoint string,
	options ...OptOptionsSetter,
) Options {
	o := Options{}

	// Setting defaults from field tag (if present)

	o.service1 = service1
	o.s3Endpoint = s3Endpoint

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithPort(opt int) OptOptionsSetter {
	return func(o *Options) {
		o.port = opt
	}
}

func (o *Options) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("service1", _validate_Options_service1(o)))
	errs.Add(errors461e464ebed9.NewValidationError("s3Endpoint", _validate_Options_s3Endpoint(o)))
	errs.Add(errors461e464ebed9.NewValidationError("port", _validate_Options_port(o)))
	return errs.AsError()
}

func _validate_Options_service1(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.service1, "required"); err != nil {
		return fmt461e464ebed9.Errorf("field `service1` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_s3Endpoint(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.s3Endpoint, "required,url"); err != nil {
		return fmt461e464ebed9.Errorf("field `s3Endpoint` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_port(o *Options) error {
	if err := validator461e464ebed9.GetValidatorFor(o).Var(o.port, "required,min=10"); err != nil {
		return fmt461e464ebed9.Errorf("field `port` did not pass the test: %w", err)
	}
	return nil
}
