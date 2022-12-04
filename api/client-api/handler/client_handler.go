package client_handler

import (
	client_request "flight-booking/api/client-api/request"
	client_response "flight-booking/api/client-api/response"
	"flight-booking/pb"
	"flight-booking/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

type ClientHandler interface {
	CreateClient(c *gin.Context)
	UpdateClient(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type clientHandler struct {
	clientClient pb.RPCClientClient
}

func NewClientHandler(clientClient pb.RPCClientClient) ClientHandler {
	return &clientHandler{
		clientClient: clientClient,
	}
}

func (h *clientHandler) CreateClient(c *gin.Context) {
	req := client_request.CreateClientRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Client{
		Role:           req.Role,
		Name:           req.Name,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		DateOfBith:     req.DateOfBith,
		IdentityCard:   req.IdentityCard,
		Address:        req.Address,
		MembershipCard: req.MembershipCard,
		Password:       req.Password,
		Status:         req.Status,
	}

	// encrypt pwd
	if len(strings.TrimSpace(req.Password)) > 0 {
		// To encrypt the StringToEncrypt
		encText, err := util.Encrypt(req.Password)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusText(http.StatusInternalServerError),
				"error":  err.Error(),
			})
			return
		}

		pReq.Password = encText
	}

	pRes, err := h.clientClient.CreateClient(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &client_response.ClientResponse{
		Id:             pRes.Id,
		Role:           pRes.Role,
		Name:           pRes.Name,
		Email:          pRes.Email,
		PhoneNumber:    pRes.PhoneNumber,
		DateOfBith:     pRes.DateOfBith,
		IdentityCard:   pRes.IdentityCard,
		Address:        pRes.Address,
		MembershipCard: pRes.MembershipCard,
		Status:         pRes.Status,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *clientHandler) UpdateClient(c *gin.Context) {
	req := client_request.UpdateClientRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReq := &pb.Client{
		Id:             req.Id,
		Role:           req.Role,
		Name:           req.Name,
		Email:          req.Email,
		PhoneNumber:    req.PhoneNumber,
		DateOfBith:     req.DateOfBith,
		IdentityCard:   req.IdentityCard,
		Address:        req.Address,
		MembershipCard: req.MembershipCard,
		Status:         req.Status,
	}

	// encrypt pwd
	if len(strings.TrimSpace(req.Password)) > 0 {
		// To encrypt the StringToEncrypt
		encText, err := util.Encrypt(req.Password)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusText(http.StatusInternalServerError),
				"error":  err.Error(),
			})
			return
		}

		pReq.Password = encText
	}

	pRes, err := h.clientClient.UpdateClient(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	dto := &client_response.ClientResponse{
		Id:             pRes.Id,
		Role:           pRes.Role,
		Name:           pRes.Name,
		Email:          pRes.Email,
		PhoneNumber:    pRes.PhoneNumber,
		DateOfBith:     pRes.DateOfBith,
		IdentityCard:   pRes.IdentityCard,
		Address:        pRes.Address,
		MembershipCard: pRes.MembershipCard,
		Status:         pRes.Status,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": dto,
	})
}

func (h *clientHandler) ChangePassword(c *gin.Context) {
	req := client_request.ChangePasswordRequest{}

	if err := c.ShouldBind(&req); err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			errMessages := make([]string, 0)
			for _, v := range validateErrors {
				errMessages = append(errMessages, v.Error())
			}

			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"status": http.StatusText(http.StatusBadRequest),
				"error":  errMessages,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  err.Error(),
		})

		return
	}

	pReqCheck := &pb.ClientParamId{
		Id: req.Id,
	}
	// Check existed
	pResCheck, err := h.clientClient.FindById(c.Request.Context(), pReqCheck)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	if pResCheck == nil || pResCheck.Id == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusNotFound),
			"error":  "ID does not exist in the system",
		})
		return
	}

	// Check old password not match
	decText, err := util.Decrypt(pResCheck.Password)
	if err != nil {
		fmt.Println("error decrypting your classified text: ", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	if req.OldPassword != decText {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 99,
			"error":  "Old password not match",
		})
		return
	}

	if req.NewPassword != req.ConfirmPassword {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": 99,
			"error":  "New password not match with Confirm password",
		})
		return
	}

	pReq := &pb.ChangePasswordRequest{
		ClientId: req.Id,
	}

	// encrypt pwd
	if len(strings.TrimSpace(req.NewPassword)) > 0 {
		// To encrypt the StringToEncrypt
		encText, err := util.Encrypt(req.NewPassword)
		if err != nil {
			fmt.Println("error encrypting your classified text: ", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusText(http.StatusInternalServerError),
				"error":  err.Error(),
			})
			return
		}

		pReq.NewPassword = encText
	}

	pRes, err := h.clientClient.ChangePassword(c.Request.Context(), pReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusText(http.StatusInternalServerError),
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusText(http.StatusOK),
		"payload": pRes,
	})
}
