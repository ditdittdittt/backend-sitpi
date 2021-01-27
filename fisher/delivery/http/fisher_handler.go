package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/ditdittdittt/backend-sitpi/domain"
	_response "github.com/ditdittdittt/backend-sitpi/domain/response"
	"github.com/ditdittdittt/backend-sitpi/helper"
)

type FisherHandler struct {
	FUsecase domain.FisherUsecase
}

func NewFisherHandler(router *mux.Router, uc domain.FisherUsecase) {
	handler := &FisherHandler{FUsecase: uc}
	router.HandleFunc("/fisher", handler.FetchFisher).Methods("GET")
	router.HandleFunc("/fisher/inquiry", handler.Inquiry).Methods("GET")
	router.HandleFunc("/fisher/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/fisher", handler.Store).Methods("POST")
	router.HandleFunc("/fisher/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/fisher/{id}", handler.Delete).Methods("DELETE")
}

func (h *FisherHandler) FetchFisher(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listFisher, err := h.FUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch fisher data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch fisher data"
		response.Data = listFisher
	}

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	fisher, err := h.FUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID fisher data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by ID fisher data"
		response.Data = fisher
	}

	helper.SetResponse(res, req, response)

}

func (h *FisherHandler) Store(res http.ResponseWriter, req *http.Request) {
	request := &domain.StoreFisherRequest{}
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
	err = h.FUsecase.Store(ctx, request)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to store fisher data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store fisher data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) Update(res http.ResponseWriter, req *http.Request) {
	request := &domain.UpdateFisherRequest{}
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

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
	err = h.FUsecase.Update(ctx, id, request)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update fisher data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update fisher data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	err := h.FUsecase.Delete(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete fisher data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete fisher data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) Inquiry(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listFisher, err := h.FUsecase.Inquiry(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to inquiry fisher data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to inquiry fisher data"
		response.Data = listFisher
	}

	helper.SetResponse(res, req, response)
}
