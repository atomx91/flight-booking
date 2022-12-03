package client_repo

import (
	"context"
	"flight-booking/database"
	client_model "flight-booking/grpc/client-grpc/model"
	client_request "flight-booking/grpc/client-grpc/request"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Embeded struct

type ClientRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*client_model.Client, error)
	CreateClient(ctx context.Context, model *client_model.Client) (*client_model.Client, error)
	UpdateClient(ctx context.Context, model *client_model.Client) (*client_model.Client, error)
	SearchClient(ctx context.Context, req *client_request.SearchClientRequest) ([]*client_model.Client, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (ClientRepository, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	db = db.Debug()

	err = db.AutoMigrate(
		&client_model.Client{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db}, nil
}

func (m *dbmanager) FindById(ctx context.Context, id uuid.UUID) (*client_model.Client, error) {
	res := client_model.Client{}
	if err := m.Where(&client_model.Client{Id: id}).First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (m *dbmanager) CreateClient(ctx context.Context, model *client_model.Client) (*client_model.Client, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdateClient(ctx context.Context, model *client_model.Client) (*client_model.Client, error) {
	if err := m.Where(&client_model.Client{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) SearchClient(ctx context.Context, req *client_request.SearchClientRequest) ([]*client_model.Client, error) {
	clients := []*client_model.Client{}

	sbWhere := " 1=1 "
	params := []interface{}{}
	if len(strings.TrimSpace(req.Id)) > 0 {
		sbWhere += " AND Id = ? "
		params = append(params, req.Id)
	}
	if len(strings.TrimSpace(req.Name)) > 0 {
		sbWhere += " AND Name = ? "
		params = append(params, req.Name)
	}
	if req.Role >= 0 {
		sbWhere += " AND Role = ? "
		params = append(params, req.Role)
	}
	if len(strings.TrimSpace(req.Email)) > 0 {
		sbWhere += " AND Email = ? "
		params = append(params, req.Email)
	}
	if len(strings.TrimSpace(req.PhoneNumber)) > 0 {
		sbWhere += " AND phone_number = ? "
		params = append(params, req.PhoneNumber)
	}
	if len(strings.TrimSpace(req.IdentityCard)) > 0 {
		sbWhere += " AND identity_card = ? "
		params = append(params, req.IdentityCard)
	}
	if len(strings.TrimSpace(req.MembershipCard)) > 0 {
		sbWhere += " AND membership_card = ? "
		params = append(params, req.MembershipCard)
	}
	if req.Status >= 0 {
		sbWhere += " AND Status = ? "
		params = append(params, req.Status)
	}

	if err := m.Where(sbWhere, params...).Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}
