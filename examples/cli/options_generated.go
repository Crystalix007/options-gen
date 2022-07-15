// Code generated by options-gen. DO NOT EDIT.
package cli

import (
	"net/http"

	goplvalidator "github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type optOptionsMeta struct {
	setter    func(o *Options)
	validator func(o *Options) error
}

func NewOptions(
	httpClient *http.Client,
	token string,

	options ...optOptionsMeta,
) Options {
	o := Options{}
	o.httpClient = httpClient
	o.token = token

	for i := range options {
		options[i].setter(&o)
	}

	return o
}

func (o *Options) Validate() error {
	g := new(errgroup.Group)

	g.Go(func() error {
		err := _Options_httpClientValidator(o)

		return errors.Wrap(err, "invalid value for option WithHttpClient")
	})
	g.Go(func() error {
		err := _Options_tokenValidator(o)

		return errors.Wrap(err, "invalid value for option WithToken")
	})
	return g.Wait()
}

func _Options_httpClientValidator(o *Options) error {

	if err := goplvalidator.New().Var(o.httpClient, "required"); err != nil {
		return errors.Wrap(err, "field `httpClient` did not pass the test")
	}

	return nil
}

func _Options_tokenValidator(o *Options) error {

	return nil
}
