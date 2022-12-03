package main

import (
	booking_handler "flight-booking/api/booking-api/handler"
	client_handler "flight-booking/api/client-api/handler"
	flight_handler "flight-booking/api/flight-api/handler"
	"flight-booking/middleware"
	"flight-booking/pb"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	//Create grpc client connect
	conn, err := grpc.Dial(":4000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	//Singleton
	clientClient := pb.NewRPCClientClient(conn)
	bookingClient := pb.NewRPCBookingClient(conn)
	flightClient := pb.NewRPCFlightClient(conn)

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//Handler for GIN Gonic
	hClient := client_handler.NewClientHandler(clientClient)
	hFlight := flight_handler.NewFlightHandler(flightClient)
	hBooking := booking_handler.NewBookingHandler(bookingClient, clientClient, flightClient)
	os.Setenv("GIN_MODE", "debug")
	g := gin.Default()
	g.Use(middleware.LoggingMiddleware(logger))

	//Create routes
	gr := g.Group("/v1/api")

	// API Client
	gr.POST("/client/create", hClient.CreateClient)
	gr.POST("/client/update", hClient.UpdateClient)
	gr.POST("/client/changePassword", hClient.ChangePassword)
	gr.POST("/client/viewBookingHistory", hBooking.BookingHistory)
	gr.POST("/client/searchBooking", hBooking.SearchBooking)

	// API Booking
	gr.POST("/booking/create", hBooking.ClientBooking)
	gr.POST("/booking/guest", hBooking.GuestBooking)
	gr.POST("/booking/cancel", hBooking.CancelBooking)

	// API Flight
	gr.POST("/flight/create", hFlight.CreateFlight)
	gr.POST("/flight/update", hFlight.UpdateFlight)
	gr.POST("/flight/search", hFlight.SearchFlight)

	//Listen and serve
	http.ListenAndServe(":5000", g)
}
