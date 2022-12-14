package booking_handler

import (
	"context"
	"database/sql"
	booking_model "flight-booking/grpc/booking-grpc/model"
	booking_repo "flight-booking/grpc/booking-grpc/repo"
	booking_request "flight-booking/grpc/booking-grpc/request"
	"flight-booking/pb"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookingHandler struct {
	pb.UnimplementedRPCBookingServer
	bookingRepository booking_repo.BookingRepository
	mu                *sync.Mutex
}

func (h *BookingHandler) FindById(ctx context.Context, in *pb.BookingParamId) (*pb.Booking, error) {
	out, err := h.bookingRepository.FindById(ctx, uuid.MustParse(in.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponse(), nil
}

func NewBookingHandler(bookingRepository booking_repo.BookingRepository) (*BookingHandler, error) {
	return &BookingHandler{
		bookingRepository: bookingRepository,
		mu:                &sync.Mutex{},
	}, nil
}

func (h *BookingHandler) CreateBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	req := &booking_model.Booking{
		Id:         uuid.New(),
		ClientId:   in.ClientId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		BookedSlot: in.BookedSlot,
		BookedDate: time.Now(),
		Status:     in.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	out, err := h.bookingRepository.CreateBooking(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponseForCreate(), nil
}

func (h *BookingHandler) UpdateBooking(ctx context.Context, in *pb.Booking) (*pb.Booking, error) {
	req, err := h.bookingRepository.FindById(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Code != "" {
		req.Code = in.Code
	}

	if in.BookedSlot > 0 {
		req.BookedSlot = in.BookedSlot
	}

	if in.Status != "" {
		req.Status = in.Status
	}

	req.UpdatedAt = time.Now()

	out, err := h.bookingRepository.UpdateBooking(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponseForCreate(), nil
}

func (h *BookingHandler) SearchBooking(ctx context.Context, in *pb.SearchBookingRequest) (*pb.SearchBookingResponse, error) {
	params := &booking_request.SearchBookingRequest{
		Id:       in.Id,
		ClientId: in.ClientId,
		FlightId: in.FlightId,
		Code:     in.Code,
		Status:   in.Status,
		// FromDate:   in.FromDate.AsTime(),
		// ToDate:     in.ToDate.AsTime(),
	}

	bookings, err := h.bookingRepository.SearchBooking(ctx, params)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.SearchBookingResponse{
		Booking: []*pb.Booking{},
	}

	for _, bk := range bookings {
		pRes.Booking = append(pRes.Booking, bk.ToResponse())
	}

	if err != nil {
		return nil, err
	}

	return pRes, nil
}
