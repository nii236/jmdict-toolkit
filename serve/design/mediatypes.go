package design

import (
	. "github.com/goadesign/goa/design"
	apidsl "github.com/goadesign/goa/design/apidsl"
)

// Word describes a single operand with a name and an integer value.
var Word = apidsl.Type("Word", func() {
	apidsl.Attribute("word", String, "Word to be translated")
	apidsl.Required("word")
})
