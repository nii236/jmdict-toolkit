package serve

import (
	"fmt"
	"regexp"

	"github.com/goadesign/goa"
	"github.com/nii236/jmdict-toolkit/serve/app"
)

type language int

const (
	english language = 1 + iota
	japanese
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
	fmt.Println(ctx.Value("Name"))

	// detectEJ(ctx.AllParams().Get("word"))
	return nil
}

// detectEJ returns English if English, Japanese if Japanese
func detectEJ(input string) language {
	fmt.Println("Input:", input)
	r, _ := regexp.Compile("[a-z]+")
	if r.MatchString(input) {
		fmt.Println("matched with English")
		return english
	}
	fmt.Println("matched with Japanese")
	return japanese
}
