package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("api", func() {
	Description("api for vrp app")
	Host("137.74.174.141:16005")
	Scheme("http")
})

var _ = Resource("swagger", func() {
	Description("The API Swagger specification")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui")
})
