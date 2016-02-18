//************************************************************************//
// API "jmdict": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/nii236/jmdict
// --design=github.com/nii236/jmdict/serve/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// TranslateTranslateContext provides the Translate translate action context.
type TranslateTranslateContext struct {
	*goa.Context
	Payload *TranslateTranslatePayload
}

// NewTranslateTranslateContext parses the incoming request URL and body, performs validations and creates the
// context used by the Translate controller translate action.
func NewTranslateTranslateContext(c *goa.Context) (*TranslateTranslateContext, error) {
	var err error
	ctx := TranslateTranslateContext{Context: c}
	return &ctx, err
}

// TranslateTranslatePayload is the Translate translate action payload.
type TranslateTranslatePayload struct {
	// Word to be translated
	Word string `json:"word" xml:"word"`
}

// Validate runs the validation rules defined in the design.
func (payload *TranslateTranslatePayload) Validate() (err error) {
	if payload.Word == "" {
		err = goa.MissingAttributeError(`raw`, "word", err)
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *TranslateTranslateContext) OK(resp []byte) error {
	ctx.Header().Set("Content-Type", "plain/text")
	return ctx.RespondBytes(200, resp)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *TranslateTranslateContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}
