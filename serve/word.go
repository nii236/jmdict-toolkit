package serve

import (
	"fmt"
	"regexp"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/nii236/jmdict-toolkit/serve/app"
)

var db gorm.DB

// WordController implements theWord resource.
type WordController struct {
	goa.Controller
}

// NewWordController creates a Word controller.
func NewWordController(service goa.Service) app.WordController {
	db, _ = gorm.Open("sqlite3", "data/gorm.db")

	return &WordController{Controller: service.NewController("Word")}
}

// Translate runs the translate action.
func (c *WordController) Translate(ctx *app.TranslateWordContext) error {
	english, err := detectJapanese(ctx.Payload.Word)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if !english {
		fmt.Println("Japanese word detected")
		ctx.WriteHeader(200)
		result := translateWord(ctx.Payload.Word)
		fmt.Println(result)
	} else {
		fmt.Println("English word detected")
	}
	return nil
}

func detectJapanese(word string) (bool, error) {
	match, err := regexp.MatchString("[a-z]+", word)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return match, nil
}

func translateWord(word string) string {
	result := db.Where("KanjiElementBase = ?", word)
	fmt.Println(result)
	return "PLACEHOLDER TRANSLATION"
	// return result.Value.(string)
}
