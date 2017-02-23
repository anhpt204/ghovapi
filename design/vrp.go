package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var VRPDepotPayload = Type("VRPDepotPayload", func() {
	Description("Depot location")
	// Attribute("address", String, "Address of the depot")
	Attribute("latitude", Number, "Latitude address")
	Attribute("longitude", Number, "Longitude address")
	Attribute("nbvehicles", Integer, "Number of vehicles")
	Attribute("capacity", Number, "Vehicles capacity")

	Required("latitude", "longitude", "nbvehicles", "capacity")
})

var VRPOrderPayload = Type("VRPOrderPayload", func() {
	Description("Order info for vrp problem")
	Attribute("id", String, "ID of customer")
	// Attribute("name", String, "Customer name")
	// Attribute("address", String, "Address of customer")
	Attribute("latitude", Number)
	Attribute("longitude", Number)
	// Attribute("time", TimeWindow, "Time for receiving")
	// Attribute("demand", Number, "Demand")

	Required("id", "latitude", "longitude")
})

var VRPRouteNode = Type("VRPRouteNode", func() {
	Description("Một node (order) trên tuyến đường phục vụ")
	Attribute("latitude", Number)
	Attribute("longitude", Number)
})

var VRPTour = Type("VRPTour", func() {
	Description("Tour")
	Attribute("depot", VRPDepotPayload, "depot")
	Attribute("nodes", ArrayOf(VRPRouteNode), "danh sach cac order duoc phuc vu")
	Attribute("duration", Number, "Tour duration")
})

var CVRPInputPayload = Type("CVRPInputPayload", func() {
	Description("Input for CVRP Solver")
	Attribute("session_key", String, "public key")
	Attribute("depot", VRPDepotPayload, "depot")
	Attribute("requests", ArrayOf(VRPOrderPayload), "requests")
	Attribute("orderbalancing", Boolean, "So order moi xe gan bang nhau")
	Attribute("maxnodesperroute", Integer, "Maximum number of nodes each route")
	Attribute("maxroutelength", Number, "Maximum route length")
	Attribute("submittime", Integer, "Submited datetime in for of epoch")

})

var CVRPOutputMedia = MediaType("application/cvrp.solution", func() {
	Description("Solution for cvrp")
	ContentType("application/json")
	Attributes(func() {
		Attribute("tours", ArrayOf(VRPTour), "List of routes")
		Attribute("duration", Number, "Solution duration")
	})

	View("default", func() {
		Attribute("tours")
		Attribute("duration")
	})
})

var _ = Resource("VRP", func() {
	Description("VRP System")
	BasePath("/api/v1/vrp")

	Action("gasolver", func() {
		Description("solve VRP problem for all orders")
		Routing(POST("/ga"))
		Payload(CVRPInputPayload)
		Response(OK, CVRPOutputMedia)

	})

	Action("greedysolver", func() {
		Description("solve VRP problem for all orders")
		Routing(POST("/greedy"))
		Payload(CVRPInputPayload)
		Response(OK, CVRPOutputMedia)

	})

})
