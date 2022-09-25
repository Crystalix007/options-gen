// Code generated by options-gen. DO NOT EDIT.
package cli

import (
	"fmt"
	"net/http"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
	validator461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/validator"
)

type OptOptionsSetter func(o *Options)

func NewOptions(
	httpClient *http.Client,
	token string,
	options ...OptOptionsSetter,
) Options {
	o := Options{}
	o.httpClient = httpClient
	o.token = token

	for _, opt := range options {
		opt(&o)
	}
	return o
}

func (o *Options) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)
	errs.Add(errors461e464ebed9.NewValidationError("httpClient", _validate_Options_httpClient(o)))
	return errs.AsError()
}

func _validate_Options_httpClient(o *Options) error {
	if err := validator461e464ebed9.GetProvidedValidatorOrDefault(o).Var(o.httpClient, "required"); err != nil {
		return fmt.Errorf("field `httpClient` did not pass the test: %w", err)
	}
	return nil
}
