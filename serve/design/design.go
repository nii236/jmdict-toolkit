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

var _ = Resource("J->E", func() {
	BasePath("/je")
	Action("translate", func() {
		Description("Translate a Japanese word to English")
		Routing(POST)
	})
	Response(OK)
	Response(NotFound)
})

var _ = Resource("E->J", func() {
	BasePath("/ej")
	Action("translate", func() {
		Description("Translate an English word to Japanese")
		Routing(POST)
	})
	Response(OK)
	Response(NotFound)
})
