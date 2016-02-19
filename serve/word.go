package serve

import (
	"github.com/goadesign/goa"
	"github.com/nii236/jmdict-toolkit/serve/app"
)

// WordController implements theWord resource.
type WordController struct {
	goa.Controller
}

// NewWordController creates a Word controller.
func NewWordController(service goa.Service) app.WordController {
	return &WordController{Controller: service.NewController("Word")}
}

// Translate runs the translate action.
func (c *WordController) Translate(ctx *app.TranslateWordContext) error {
	return nil
}
