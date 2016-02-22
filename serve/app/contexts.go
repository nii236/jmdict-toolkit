//************************************************************************//
// API "jmdict": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/jmdict-toolkit/serve
// --design=github.com/nii236/jmdict-toolkit/serve/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// TranslateWordContext provides the Word translate action context.
type TranslateWordContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *TranslateWordPayload
}

// NewTranslateWordContext parses the incoming request URL and body, performs validations and creates the
// context used by the Word controller translate action.
func NewTranslateWordContext(ctx context.Context) (*TranslateWordContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := TranslateWordContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	return &rctx, err
}

// TranslateWordPayload is the Word translate action payload.
type TranslateWordPayload struct {
	// Word to be translated
	Word string `json:"word" xml:"word"`
}

// Validate runs the validation rules defined in the design.
func (payload *TranslateWordPayload) Validate() (err error) {
	if payload.Word == "" {
		err = goa.MissingAttributeError(`raw`, "word", err)
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *TranslateWordContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "plain/text")
	ctx.ResponseData.WriteHeader(200)
	ctx.ResponseData.Write(resp)
	return nil
}

// NotFound sends a HTTP response with status code 404.
func (ctx *TranslateWordContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
