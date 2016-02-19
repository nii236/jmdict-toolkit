package design

import (
	. "github.com/goadesign/goa/design"
	apidsl "github.com/goadesign/goa/design/apidsl"
)

var _ = apidsl.API("jmdict", func() {
	apidsl.Title("JMDict REST API")
	apidsl.Description("This is a REST API hosting a J-E Dictionary")
	apidsl.Scheme("http")
	apidsl.Host("localhost:8080")
})

var _ = apidsl.Resource("Word", func() {
	apidsl.BasePath("/translate")
	apidsl.Action("translate", func() {
		apidsl.Description("Translate a word between Japanese and English")
		apidsl.Routing(apidsl.POST(""))
		apidsl.Payload(Word)
	})
	apidsl.Response(apidsl.OK)
	apidsl.Response(apidsl.NotFound)
})
