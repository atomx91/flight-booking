package main

import (
	"flag"
	booking_handler "flight-booking/grpc/booking-grpc/handler"
	booking_repo "flight-booking/grpc/booking-grpc/repo"
	client_handler "flight-booking/grpc/client-grpc/handler"
	client_repo "flight-booking/grpc/client-grpc/repo"
	flight_handler "flight-booking/grpc/flight-grpc/handler"
	flight_repo "flight-booking/grpc/flight-grpc/repo"
	"flight-booking/helper"
	"flight-booking/intercepter"
	"flight-booking/pb"
	"fmt"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	configFile = flag.String("config-file", "../helper/config.yml", "Location of config file")
	port       = flag.Int("port", 4000, "Port of grpc")
)

func init() {
	flag.Parse()
}

func main() {
	err := helper.AutoBindConfig(*configFile)
	if err != nil {
		panic(err)
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			intercepter.UnaryServerLoggingIntercepter(logger),
		)),
	)

	reflection.Register(s)

	// Initial client repo START
	clientRepository, err := client_repo.NewDBManager()
	if err != nil {
		panic(err)
	}

	h, err := client_handler.NewClientHandler(clientRepository)
	if err != nil {
		panic(err)
	}
	pb.RegisterRPCClientServer(s, h)
	// Initial client repo END

	// Initial Flight repo START
	flightRepository, errFlight := flight_repo.NewDBManager()
	if errFlight != nil {
		panic(errFlight)
	}

	hFlight, errFlight := flight_handler.NewFlightHandler(flightRepository)
	if errFlight != nil {
		panic(errFlight)
	}
	pb.RegisterRPCFlightServer(s, hFlight)
	// Initial Flight repo END

	// Initial Booking repo START
	bookingRepository, errBooking := booking_repo.NewDBManager()
	if errBooking != nil {
		panic(errBooking)
	}

	hBooking, errBooking := booking_handler.NewBookingHandler(bookingRepository)
	if errBooking != nil {
		panic(errBooking)
	}
	pb.RegisterRPCBookingServer(s, hBooking)
	// Initial Booking repo END

	fmt.Printf("Listen at port: %v\n", *port)

	s.Serve(listen)
}
