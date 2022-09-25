// Code generated by options-gen. DO NOT EDIT.
package testcase

import (
	"fmt"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptOptionsSetter func(o *Options)

func NewOptions(
	amount int,
	age int,
	options ...OptOptionsSetter,
) Options {
	o := Options{}
	o.amount = amount
	o.age = age

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func (o *Options) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("age", _validate_Options_age(o)))
	return errs.AsError()
}

func _validate_Options_age(o *Options) error {
	if err := validator461e464ebed9.GetProvidedValidatorOrDefault(o).Var(o.age, "child"); err != nil {
		return fmt.Errorf("field `age` did not pass the test: %w", err)
	}
	return nil
}
