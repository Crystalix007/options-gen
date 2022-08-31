// Code generated by options-gen. DO NOT EDIT.
package testcase

import (
	"net/http"

	errors461e464ebed9 "github.com/kazhuravlev/options-gen/pkg/errors"
)

type optOptionsMeta struct {
	setter    func(o *Options)
	validator func(o *Options) error
}

func NewOptions(
	FnTypeParam FnType,
	FnParam func(server *http.Server) error,
	HandlerFunc http.HandlerFunc,
	Middleware func(next http.HandlerFunc) http.HandlerFunc,
	Local localFnType,

	options ...optOptionsMeta,
) Options {
	o := Options{}
	o.FnTypeParam = FnTypeParam
	o.FnParam = FnParam
	o.HandlerFunc = HandlerFunc
	o.Middleware = Middleware
	o.Local = Local

	for i := range options {
		options[i].setter(&o)
	}

	return o
}

func WithOptFnTypeParam(opt FnType) optOptionsMeta {
	return optOptionsMeta{
		setter:    func(o *Options) { o.OptFnTypeParam = opt },
		validator: _Options_OptFnTypeParamValidator,
	}
}

func WithOptFnParam(opt func(server *http.Server) error) optOptionsMeta {
	return optOptionsMeta{
		setter:    func(o *Options) { o.OptFnParam = opt },
		validator: _Options_OptFnParamValidator,
	}
}

func WithOptHandlerFunc(opt http.HandlerFunc) optOptionsMeta {
	return optOptionsMeta{
		setter:    func(o *Options) { o.OptHandlerFunc = opt },
		validator: _Options_OptHandlerFuncValidator,
	}
}

func WithOptMiddleware(opt func(next http.HandlerFunc) http.HandlerFunc) optOptionsMeta {
	return optOptionsMeta{
		setter:    func(o *Options) { o.OptMiddleware = opt },
		validator: _Options_OptMiddlewareValidator,
	}
}

func WithOptLocal(opt localFnType) optOptionsMeta {
	return optOptionsMeta{
		setter:    func(o *Options) { o.OptLocal = opt },
		validator: _Options_OptLocalValidator,
	}
}

func (o *Options) Validate() error {
	errs := new(errors461e464ebed9.ValidationErrors)

	errs.Add(errors461e464ebed9.NewValidationError("FnTypeParam", _Options_FnTypeParamValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("FnParam", _Options_FnParamValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("HandlerFunc", _Options_HandlerFuncValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("Middleware", _Options_MiddlewareValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("Local", _Options_LocalValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("OptFnTypeParam", _Options_OptFnTypeParamValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("OptFnParam", _Options_OptFnParamValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("OptHandlerFunc", _Options_OptHandlerFuncValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("OptMiddleware", _Options_OptMiddlewareValidator(o)))
	errs.Add(errors461e464ebed9.NewValidationError("OptLocal", _Options_OptLocalValidator(o)))
	return errs.AsError()
}

func _Options_FnTypeParamValidator(o *Options) error {

	return nil
}

func _Options_FnParamValidator(o *Options) error {

	return nil
}

func _Options_HandlerFuncValidator(o *Options) error {

	return nil
}

func _Options_MiddlewareValidator(o *Options) error {

	return nil
}

func _Options_LocalValidator(o *Options) error {

	return nil
}

func _Options_OptFnTypeParamValidator(o *Options) error {

	return nil
}

func _Options_OptFnParamValidator(o *Options) error {

	return nil
}

func _Options_OptHandlerFuncValidator(o *Options) error {

	return nil
}

func _Options_OptMiddlewareValidator(o *Options) error {

	return nil
}

func _Options_OptLocalValidator(o *Options) error {

	return nil
}