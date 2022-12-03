package booking_request

type ClientBookingRequest struct {
	Slot     int32  `json:"slot" binding:"required"`
	ClientId string `json:"clientId" binding:"required"`
	FlightId string `json:"flightId" binding:"required"`
}

type GuestBookingRequest struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	PhoneNumber    string `json:"phoneNumber" binding:"required"`
	DateOfBith     string `json:"dateOfBith" binding:"required"`
	IdentityCard   string `json:"identityCard" binding:"required"`
	Address        string `json:"address" binding:"max=256,min=6"`
	MembershipCard string `json:"membershipCard"`
	FlightId       string `json:"flightId" binding:"required"`
	Slot           int32  `json:"slot" binding:"required"`
}

type CancelBookingRequest struct {
	Id string `json:"id" binding:"required"`
}

type ViewBookingRequest struct {
	ClientId string `json:"clientId" binding:"required"`
}

type SearchBookingRequest struct {
	ClientId string `json:"clientId"`
	FlightId string `json:"flightId"`
	Code     string `json:"code"`
	Status   string `json:"status"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}
