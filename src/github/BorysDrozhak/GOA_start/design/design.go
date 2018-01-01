package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Description("random service")
	Host("localhost:8080")
})

var BottlePaylod = Type("BottlePayload", func() {
	Description("bottle is the type used to create a bottle")

	Attribute("name", String, "name of bottle", func() {
		MinLength(2)
	})

	Attribute("vintage", Integer, "Vintage of bottle", func() {
		Minimum(1900)
	})

	Attribute("rating", Integer, "Rating of bottle", func() {
		Minimum(1)
		Maximum(5)
	})

	Required("name", "vintage", "rating")
})


var BottleMedia = MediaType("application/vnd.gophercon.goa.bottle", func() {
	TypeName("bottle")
	Reference(BottlePaylod)

	Attributes(func() {
		Attribute("ID", Integer, "Unique bottle ID")
		Attribute("name")
		Attribute("vintage")
		Attribute("rating")
		Required("ID", "name", "vintage", "rating")
	})

	View("default", func() {
		Attribute("ID")
		Attribute("name")
		Attribute("vintage")
		Attribute("rating")
	})
})

var _ = Resource("bottle", func() {
	Description("A wine bottle")
	BasePath("/bottles")
	Action("create", func() {
		Description("creates a bottle")
		Routing(POST("/"))
		Payload(BottlePaylod)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a bottle")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, BottlePaylod)
	})
})