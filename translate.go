package main

import (
	"github.com/goadesign/goa"
	"github.com/nii236/jmdict/app"
)

// TranslateController implements theTranslate resource.
type TranslateController struct {
	goa.Controller
}

// NewTranslateController creates a Translate controller.
func NewTranslateController(service goa.Service) app.TranslateController {
	return &TranslateController{Controller: service.NewController("Translate")}
}

// Translate runs the translate action.
func (c *TranslateController) Translate(ctx *app.TranslateTranslateContext) error {
	return nil
}
