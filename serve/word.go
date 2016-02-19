package main

import (
	"fmt"
	"regexp"

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
	fmt.Println(ctx.Value("Name"))
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
