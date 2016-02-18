package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Word describes a single operand with a name and an integer value.
var Word = Type("Word", func() {
	Attribute("word", String, "Word to be translated")
	Required("word")
})
