// Code generated by options-gen. DO NOT EDIT.
package testcase

import (
	"fmt"
	"net/http"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptOptionsSetter[KeyT int | string, TT any] func(o *Options[KeyT, TT])

func NewOptions[KeyT int | string, TT any](
	requiredHandler http.Handler,
	requiredKey KeyT,
	handler http.Handler,
	key KeyT,
	options ...OptOptionsSetter[KeyT, TT],
) Options[KeyT, TT] {
	o := Options[KeyT, TT]{}
	o.requiredHandler = requiredHandler
	o.requiredKey = requiredKey
	o.handler = handler
	o.key = key

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func WithOptHandler[KeyT int | string, TT any](opt http.Handler) OptOptionsSetter[KeyT, TT] {
	return func(o *Options[KeyT, TT]) {
		o.optHandler = opt
	}
}

func WithOptKey[KeyT int | string, TT any](opt KeyT) OptOptionsSetter[KeyT, TT] {
	return func(o *Options[KeyT, TT]) {
		o.optKey = opt
	}
}

func WithAnyOpt[KeyT int | string, TT any](opt TT) OptOptionsSetter[KeyT, TT] {
	return func(o *Options[KeyT, TT]) {
		o.anyOpt = opt
	}
}

func (o *Options[KeyT, TT]) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("requiredHandler", _validate_Options_requiredHandler[KeyT, TT](o)))
	errs.Add(errors461e464ebed9.NewValidationError("requiredKey", _validate_Options_requiredKey[KeyT, TT](o)))
	return errs.AsError()
}

func _validate_Options_requiredHandler[KeyT int | string, TT any](o *Options[KeyT, TT]) error {
	if err := validator461e464ebed9.GetProvidedValidatorOrDefault(o).Var(o.requiredHandler, "required"); err != nil {
		return fmt.Errorf("field `requiredHandler` did not pass the test: %w", err)
	}
	return nil
}

func _validate_Options_requiredKey[KeyT int | string, TT any](o *Options[KeyT, TT]) error {
	if err := validator461e464ebed9.GetProvidedValidatorOrDefault(o).Var(o.requiredKey, "required"); err != nil {
		return fmt.Errorf("field `requiredKey` did not pass the test: %w", err)
	}
	return nil
}
