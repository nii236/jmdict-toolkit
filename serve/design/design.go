package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("jmdict", func() {
	Title("JMDict REST API")
	Description("This is a REST API hosting a J-E Dictionary")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("Translate", func() {
	BasePath("/translate")
	Action("translate", func() {
		Description("Translate a word between Japanese and English")
		Routing(POST(""))
		Payload(Word)
	})
	Response(OK)
	Response(NotFound)
})
