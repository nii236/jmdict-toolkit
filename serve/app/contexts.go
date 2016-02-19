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

import "github.com/goadesign/goa"

// TranslateWordContext provides the Word translate action context.
type TranslateWordContext struct {
	*goa.Context
	Payload *TranslateWordPayload
}

// NewTranslateWordContext parses the incoming request URL and body, performs validations and creates the
// context used by the Word controller translate action.
func NewTranslateWordContext(c *goa.Context) (*TranslateWordContext, error) {
	var err error
	ctx := TranslateWordContext{Context: c}
	return &ctx, err
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
	ctx.Header().Set("Content-Type", "plain/text")
	return ctx.RespondBytes(200, resp)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *TranslateWordContext) NotFound() error {
	return ctx.RespondBytes(404, nil)
}
