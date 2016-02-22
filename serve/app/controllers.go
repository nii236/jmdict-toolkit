//************************************************************************//
// API "jmdict": Application Controllers
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
	"net/http"
)

// WordController is the controller interface for the Word actions.
type WordController interface {
	goa.Controller
	Translate(*TranslateWordContext) error
}

// MountWordController "mounts" a Word resource controller on the given service.
func MountWordController(service goa.Service, ctrl WordController) {
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
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewTranslateWordContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		if rawPayload := goa.Request(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TranslateWordPayload)
		}
		return ctrl.Translate(rctx)
	}
	mux.Handle("POST", "/translate", ctrl.HandleFunc("Translate", h, unmarshalTranslateWordPayload))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Word"}, goa.KV{"action", "Translate"}, goa.KV{"route", "POST /translate"})
}

// unmarshalTranslateWordPayload unmarshals the request body into the context request data Payload field.
func unmarshalTranslateWordPayload(ctx context.Context, req *http.Request) error {
	var payload TranslateWordPayload
	if err := goa.RequestService(ctx).DecodeRequest(req, &payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.Request(ctx).Payload = &payload
	return nil
}
