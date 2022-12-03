package client_handler

import (
	"context"
	"database/sql"
	client_model "flight-booking/grpc/client-grpc/model"
	client_repo "flight-booking/grpc/client-grpc/repo"
	client_request "flight-booking/grpc/client-grpc/request"
	"flight-booking/pb"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientHandler struct {
	pb.UnimplementedRPCClientServer
	clientRepository client_repo.ClientRepository
	mu               *sync.Mutex
}

func NewClientHandler(clientRepository client_repo.ClientRepository) (*ClientHandler, error) {
	return &ClientHandler{
		clientRepository: clientRepository,
		mu:               &sync.Mutex{},
	}, nil
}

func (h *ClientHandler) FindById(ctx context.Context, in *pb.ClientParamId) (*pb.Client, error) {
	out, err := h.clientRepository.FindById(ctx, uuid.MustParse(in.Id))

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponse(), nil
}

func (h *ClientHandler) CreateClient(ctx context.Context, in *pb.Client) (*pb.Client, error) {
	req := &client_model.Client{
		Id:             uuid.New(),
		Role:           in.Role,
		Name:           in.Name,
		Email:          in.Email,
		PhoneNumber:    in.PhoneNumber,
		DateOfBith:     in.DateOfBith,
		IdentityCard:   in.IdentityCard,
		Address:        in.Address,
		MembershipCard: in.MembershipCard,
		Password:       in.Password,
		Status:         in.Status,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	client, err := h.clientRepository.CreateClient(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return client.ToResponse(), nil
}

func (h *ClientHandler) UpdateClient(ctx context.Context, in *pb.Client) (*pb.Client, error) {
	req, err := h.clientRepository.FindById(ctx, uuid.MustParse(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.Role >= 0 {
		req.Role = in.Role
	}

	if in.Name != "" {
		req.Name = in.Name
	}

	if in.Email != "" {
		req.Email = in.Email
	}

	if in.PhoneNumber != "" {
		req.PhoneNumber = in.PhoneNumber
	}

	if in.DateOfBith != "" {
		req.DateOfBith = in.DateOfBith
	}

	if in.IdentityCard != "" {
		req.IdentityCard = in.IdentityCard
	}

	if in.Address != "" {
		req.Address = in.Address
	}

	if in.MembershipCard != "" {
		req.MembershipCard = in.MembershipCard
	}

	if in.Password != "" {
		req.Password = in.Password
	}

	if in.Status >= 0 {
		req.Status = in.Status
	}

	req.UpdatedAt = time.Now()

	out, err := h.clientRepository.UpdateClient(ctx, req)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return out.ToResponse(), nil
}

func (h *ClientHandler) ChangePassword(ctx context.Context, in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error) {
	req, err := h.clientRepository.FindById(ctx, uuid.MustParse(in.ClientId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if in.NewPassword != "" {
		req.Password = in.NewPassword
	}

	req.UpdatedAt = time.Now()

	custOut, err := h.clientRepository.UpdateClient(ctx, req)

	if custOut != nil && err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &pb.ChangePasswordResponse{
		Code:    0,
		Message: "Success",
	}

	return out, nil
}

func (h *ClientHandler) SearchClient(ctx context.Context, in *pb.SearchClientRequest) (*pb.SearchClientResponse, error) {
	clients, err := h.clientRepository.SearchClient(ctx, &client_request.SearchClientRequest{
		Name:         in.Name,
		Email:        in.Email,
		PhoneNumber:  in.PhoneNumber,
		IdentityCard: in.IdentityCard,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, err
	}

	pRes := &pb.SearchClientResponse{
		Client: []*pb.Client{},
	}

	for _, client := range clients {
		pRes.Client = append(pRes.Client, client.ToResponse())
	}

	if err != nil {
		return nil, err
	}

	return pRes, nil
}
