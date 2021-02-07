package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
	_response "github.com/ditdittdittt/backend-sitpi/domain/response"
	"github.com/ditdittdittt/backend-sitpi/helper"
)

var mySigningKey = []byte("captainjacksparrowsayshi")

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(router *mux.Router, uc domain.UserUsecase) {
	handler := &UserHandler{UserUsecase: uc}
	router.HandleFunc("/user/login", handler.Login).Methods("POST")
}

func (h *UserHandler) Login(res http.ResponseWriter, req *http.Request) {
	request := &domain.LoginUserRequest{}
	response := _response.New()

	body, err := helper.ReadRequest(req, response)
	if err != nil {
		response.Data = err.Error()
		logrus.Error(err)
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		response.Data = err.Error()
		logrus.Error(err)
	}

	err = helper.ValidateRequest(request, response)
	if err != nil {
		response.Data = err.Error()
		logrus.Error(err)
	}

	ctx := req.Context()
	jwtToken, err := h.UserUsecase.Login(ctx, request)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Login failed"
		response.Data = err.Error()
		helper.SetResponse(res, req, response)
		return
	}

	response.Code = "00"
	response.Desc = "Login success"
	response.Data = jwtToken

	helper.SetResponse(res, req, response)
}

func (h *UserHandler) ChangePassword(res http.ResponseWriter, req *http.Request) {
	request := &domain.ChangePasswordRequest{}
	response := _response.New()

	body, err := helper.ReadRequest(req, response)
	if err != nil {
		response.Data = err.Error()
		logrus.Error(err)
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		response.Data = err.Error()
		logrus.Error(err)
	}

	err = helper.ValidateRequest(request, response)
	if err != nil {
		response.Data = err.Error()
		logrus.Error(err)
	}

	ctx := req.Context()
	err = h.UserUsecase.ChangePassword(ctx, request)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Change password failed"
		response.Data = err.Error()
		helper.SetResponse(res, req, response)
		return
	}

	response.Code = "00"
	response.Desc = "Change password success"
	helper.SetResponse(res, req, response)
}
