package booking_model

import (
	client_model "flight-booking/grpc/client-grpc/model"
	flight_model "flight-booking/grpc/flight-grpc/model"
	"time"

	"flight-booking/pb"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Booking struct {
	Id         uuid.UUID            `gorm:"type:uuid;primaryKey"`
	ClientId   string               `gorm:"column:client_id"`
	FlightId   string               `gorm:"column:flight_id"`
	Code       string               `gorm:"column:code"`
	BookedSlot int32                `gorm:"column:booked_slot"`
	BookedDate time.Time            `gorm:"column:booked_date"`
	Status     string               `gorm:"column:status"`
	CreatedAt  time.Time            `gorm:"column:created_at"`
	UpdatedAt  time.Time            `gorm:"column:updated_at"`
	Client     *client_model.Client `gorm:"foreignKey:client_id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Flight     *flight_model.Flight `gorm:"foreignKey:flight_id;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (in *Booking) ToResponse() *pb.Booking {
	res := &pb.Booking{
		Id:         in.Id.String(),
		ClientId:   in.ClientId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		BookedSlot: in.BookedSlot,
		BookedDate: timestamppb.New(in.BookedDate),
		Status:     in.Status,
		CreatedAt:  timestamppb.New(in.CreatedAt),
		UpdatedAt:  timestamppb.New(in.UpdatedAt),
		Client: &pb.ClientDTO{
			Id:             in.Client.Id.String(),
			Role:           in.Client.Role,
			Name:           in.Client.Name,
			Email:          in.Client.Email,
			PhoneNumber:    in.Client.PhoneNumber,
			DateOfBith:     in.Client.DateOfBith,
			IdentityCard:   in.Client.IdentityCard,
			Address:        in.Client.Address,
			MembershipCard: in.Client.MembershipCard,
			Password:       in.Client.Password,
			Status:         in.Client.Status,
			Audit: &pb.Audit{
				UpdatedAt: timestamppb.New(in.Client.UpdatedAt),
				CreatedAt: timestamppb.New(in.CreatedAt),
			},
		},
		Flight: &pb.FlightDTO{
			Id:            in.Flight.Id.String(),
			Name:          in.Flight.Name,
			From:          in.Flight.From,
			To:            in.Flight.To,
			DepartDate:    timestamppb.New(in.Flight.DepartDate),
			Status:        in.Flight.Status,
			AvailableSlot: in.Flight.AvailableSlot,
			CreatedAt:     timestamppb.New(in.Flight.CreatedAt),
			UpdatedAt:     timestamppb.New(in.Flight.UpdatedAt),
		},
	}

	return res
}

func (in *Booking) ToResponseForCreate() *pb.Booking {
	res := &pb.Booking{
		Id:         in.Id.String(),
		ClientId:   in.ClientId,
		FlightId:   in.FlightId,
		Code:       in.Code,
		BookedSlot: in.BookedSlot,
		BookedDate: timestamppb.New(in.BookedDate),
		Status:     in.Status,
		CreatedAt:  timestamppb.New(in.CreatedAt),
		UpdatedAt:  timestamppb.New(in.UpdatedAt),
	}

	return res
}
