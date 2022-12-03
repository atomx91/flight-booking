package client_request

import "github.com/google/uuid"

type ClientRequest struct {
	Id             uuid.UUID
	Role           int32
	Name           string
	Email          string
	PhoneNumber    string
	DateOfBith     string
	IdentityCard   string
	Address        string
	MembershipCard string
	Password       string
	Status         int32
}

type SearchClientRequest struct {
	Id             string
	Role           int32
	Name           string
	Email          string
	PhoneNumber    string
	IdentityCard   string
	MembershipCard string
	Status         int32
}
