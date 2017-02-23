package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var TimeWindow = Type("TimeWindow", func() {
	Description("Time window for serving an order")
	Attribute("start", Integer, "Start time window")
	Attribute("end", Integer, "End time window")

})

var OrderPayload = Type("OrderPayload", func() {
	Description("Order info")
	Attribute("id", String, "ID of customer")
	Attribute("address", String, "Address of customer")
	Attribute("time", TimeWindow, "Time for receiving")

	Required("id", "address")
})

var DepotPayload = Type("DepotPayload", func() {
	Description("Depot location")
	Attribute("address", String, "Address of the depot")
	Attribute("latitude", Number, "Latitude address")
	Attribute("longitude", Number, "Longitude address")

	Required("latitude", "longitude")
})

var EmployeePayload = Type("EmployeePayload", func() {
	Description("Employee infomation")
	Attribute("id", String, "Employee ID")
	Attribute("area", String, "Working area of employee")

	Required("id", "area")
})

// var AreaPayload = Type("AreaPayload", func() {
// 	Description("Area info")
// 	Attribute("name", String, "Name of area")
// 	Attribute("nbshippers", Integer, "Number of shippers")
// 	Attribute("depot", DepotPayload, "Depot address")
//
// 	Required("name", "nbshippers", "depot")
// })

var RouteNode = Type("RouteNode", func() {
	Description("Một node (order) trên tuyến đường phục vụ")
	Attribute("stt", String, "Thứ tự phục vụ")
	Attribute("id_order", String, "id cua don hang")
	Attribute("latitude", Number)
	Attribute("longitude", Number)

	// Required("stt", "id_order", "latitude", "longitude")
})

var RouteType = Type("RouteType", func() {
	Description("Route")
	Attribute("id_nhanvien", String, "id của nhân viên thực hiện route nay")
	Attribute("orders", ArrayOf(RouteNode), "danh sach cac order duoc phuc vu")

	Required("id_nhanvien")
	// Attribute("duration", Number, "Total duration")
})

var ExceptionType = Type("ExceptionType", func() {
	Description("order mà có địa chỉ không xác định đượng trên googe maps")
	Attribute("id", String, "order id")
	Required("id")
})

var SystemInputPayload = Type("InputPayload", func() {
	Description("System Imput Payload")
	Attribute("session_key", String, "public key")
	Attribute("orders", ArrayOf(OrderPayload))
	Attribute("employees", ArrayOf(EmployeePayload))
	Attribute("dungsai", Number, "dung sai bán kính giao hàng ngoài khu vực của nhân viên")
	Attribute("depot", DepotPayload, "Kho hàng")
})

var NhanVienInputPayload = Type("NhanVienInputPayload", func() {
	Description("Nhan vien Imput Payload")
	Attribute("session_key", String, "public key")
	Attribute("orders", ArrayOf(OrderPayload))
	Attribute("employees", String, "id of employee")
	Attribute("depot", DepotPayload, "Kho hàng")

})

var SystemOutputMedia = MediaType("application/ongvang.system", func() {
	Description("Solution for vrp")
	ContentType("application/json")
	Attributes(func() {
		Attribute("solution", ArrayOf(RouteType), "List of routes")
		Attribute("exception", ArrayOf(ExceptionType), "list of exception order address")
	})

	View("default", func() {
		Attribute("solution")
		Attribute("exception")
	})
})

var NhanVienOutputMedia = MediaType("application/ongvang.nhanvien", func() {
	Description("TSP solution for 1 vehicle")
	ContentType("application/json")
	Attributes(func() {
		Attribute("solution", RouteType, "solution")
		Attribute("exception", ArrayOf(ExceptionType), "list of exception order address")
	})

	View("default", func() {
		Attribute("solution")
		Attribute("exception")
	})
})

var _ = Resource("ghovtsp", func() {
	Description("Ongvang TSP API")
	BasePath("/api/v1/tsp")

	Action("system_solve", func() {
		Description("solve TSP problem for all orders")
		Routing(POST("/ongvang_system"))
		Payload(SystemInputPayload)
		Response(OK, SystemOutputMedia)

	})

	Action("nhanvien_solve", func() {
		Description("solve TSP problem for all orders")
		Routing(POST("/ongvang_nhanvien"))
		Payload(NhanVienInputPayload)
		Response(OK, NhanVienOutputMedia)

	})

	// Consumes("application/json")
})
