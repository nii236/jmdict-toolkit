//************************************************************************//
// API "jmdict": Application Controllers
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

// TranslateController is the controller interface for the Translate actions.
type TranslateController interface {
	goa.Controller
	Translate(*TranslateTranslateContext) error
}

// MountTranslateController "mounts" a Translate resource controller on the given service.
func MountTranslateController(service goa.Service, ctrl TranslateController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(c *goa.Context) error {
		ctx, err := NewTranslateTranslateContext(c)
		ctx.Payload = ctx.RawPayload().(*TranslateTranslatePayload)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Translate(ctx)
	}
	mux.Handle("POST", "/translate", ctrl.HandleFunc("Translate", h, unmarshalTranslateTranslatePayload))
	service.Info("mount", "ctrl", "Translate", "action", "Translate", "route", "POST /translate")
}

// unmarshalTranslateTranslatePayload unmarshals the request body.
func unmarshalTranslateTranslatePayload(ctx *goa.Context) error {
	payload := &TranslateTranslatePayload{}
	if err := ctx.Service().DecodeRequest(ctx, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	ctx.SetPayload(payload)
	return nil
}
