package http

import (
	"encoding/json"
	"github.com/ditdittdittt/backend-sitpi/domain"
	_response "github.com/ditdittdittt/backend-sitpi/domain/response"
	"github.com/ditdittdittt/backend-sitpi/helper"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type FisherHandler struct {
	FUsecase domain.FisherUsecase
}

type FetchFisherRequest struct {
	Cursor string `json:"cursor"`
	Num    int64  `json:"num"`
}

type GetByIDRequest struct {
	ID int64 `json:"id"`
}

type StoreRequest struct {
	Nik     string `json:"nik"`
	Address string `json:"address"`
}

type UpdateRequest struct {
	ID      int64  `json:"id"`
	Nik     string `json:"nik"`
	Address string `json:"address"`
}

type DeleteRequest struct {
	ID int64 `json:"id"`
}

func NewFisherHandler(router *mux.Router, uc domain.FisherUsecase) {
	handler := &FisherHandler{FUsecase: uc}
	router.HandleFunc("/fisher/index", handler.FetchFisher).Methods("GET")
	router.HandleFunc("/fisher/get_by_id", handler.GetByID).Methods("GET")
	router.HandleFunc("/fisher/store", handler.Store).Methods("POST")
	router.HandleFunc("/fisher/update", handler.Update).Methods("PUT")
	router.HandleFunc("/fisher/delete", handler.Delete).Methods("DELETE")
}

func (h *FisherHandler) FetchFisher(res http.ResponseWriter, req *http.Request) {
	request := &FetchFisherRequest{}
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
	listFisher, _, err := h.FUsecase.Fetch(ctx, request.Cursor, request.Num)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch fisher data"
		response.Data = err
	}

	response.Code = "00"
	response.Desc = "Success to fetch fisher data"
	response.Data = listFisher

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	request := &GetByIDRequest{}
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
	fisher, err := h.FUsecase.GetByID(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to get by ID fisher data"
	response.Data = fisher

	helper.SetResponse(res, req, response)

}

func (h *FisherHandler) Store(res http.ResponseWriter, req *http.Request) {
	request := &StoreRequest{}
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
	fisher := &domain.Fisher{
		Nik:       request.Nik,
		Address:   request.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = h.FUsecase.Store(ctx, fisher)

	response.Code = "00"
	response.Desc = "Success to store fisher data"
	response.Data = fisher

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) Update(res http.ResponseWriter, req *http.Request) {
	request := &UpdateRequest{}
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
	fisher := &domain.Fisher{
		ID:        request.ID,
		Nik:       request.Nik,
		Address:   request.Address,
		UpdatedAt: time.Now(),
	}
	err = h.FUsecase.Update(ctx, fisher)

	response.Code = "00"
	response.Desc = "Success to update fisher data"
	response.Data = fisher

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) Delete(res http.ResponseWriter, req *http.Request) {
	request := &DeleteRequest{}
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
	err = h.FUsecase.Delete(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to delete fisher data"

	helper.SetResponse(res, req, response)
}
