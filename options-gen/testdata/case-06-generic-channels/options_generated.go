// Code generated by options-gen. DO NOT EDIT.
package testcase

import (
	"fmt"

	goplvalidator "github.com/go-playground/validator/v10"
	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
)

var _validator461e464ebed9 = goplvalidator.New()

type OptOptionsSetter[T any] func(o *Options[T])

func NewOptions[T any](
	ch1 chan T,
	ch2 <-chan T,
	options ...OptOptionsSetter[T],
) Options[T] {
	o := Options[T]{}
	o.ch1 = ch1
	o.ch2 = ch2

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithCh3[T any](opt chan T) OptOptionsSetter[T] {
	return func(o *Options[T]) {
		o.ch3 = opt
	}
}

func WithCh4[T any](opt <-chan T) OptOptionsSetter[T] {
	return func(o *Options[T]) {
		o.ch4 = opt
	}
}

func (o *Options[T]) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("Ch1", _validate_Options_ch1[T](o)))
	return errs.AsError()
}

func _validate_Options_ch1[T any](o *Options[T]) error {
	if err := _getOptsValidatorOrDefault(o).Var(o.ch1, "required"); err != nil {
		return fmt.Errorf("field `ch1` did not pass the test: %w", err)
	}
	return nil
}

func _getOptsValidatorOrDefault(opts any) *goplvalidator.Validate {
	if v, ok := opts.(interface {
		Validator() *goplvalidator.Validate
	}); ok {
		return v.Validator()
	}
	return _validator461e464ebed9
}
