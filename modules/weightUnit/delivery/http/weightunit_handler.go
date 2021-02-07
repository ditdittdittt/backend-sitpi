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

type WeightUnitHandler struct {
	WUUsecase domain.WeightUnitUsecase
}

type StoreRequest struct {
	Unit string `json:"unit"`
}

type UpdateRequest struct {
	Unit string `json:"unit"`
}

func NewWeightUnitHandler(router *mux.Router, uc domain.WeightUnitUsecase) {
	handler := &WeightUnitHandler{WUUsecase: uc}
	router.HandleFunc("/weight_unit", handler.Fetch).Methods("GET")
	router.HandleFunc("/weight_unit", handler.Store).Methods("POST")
	router.HandleFunc("/weight_unit/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/weight_unit/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/weight_unit/{id}", handler.Delete).Methods("DELETE")

}

func (h *WeightUnitHandler) Fetch(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listWeightUnit, err := h.WUUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch weight unit data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch weight unit data"
		response.Data = listWeightUnit
	}

	helper.SetResponse(res, req, response)
}

func (h *WeightUnitHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	weightUnit, err := h.WUUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID weight unit data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by ID weight unit data"
		response.Data = weightUnit
	}

	helper.SetResponse(res, req, response)
}

func (h *WeightUnitHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	weightUnit := &domain.WeightUnit{
		Unit: request.Unit,
	}

	err = h.WUUsecase.Store(ctx, weightUnit)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to store weight unit data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store weight unit data"
	}

	helper.SetResponse(res, req, response)
}

func (h *WeightUnitHandler) Update(res http.ResponseWriter, req *http.Request) {
	request := &UpdateRequest{}
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
	weightUnit := &domain.WeightUnit{
		ID:   id,
		Unit: request.Unit,
	}

	err = h.WUUsecase.Update(ctx, weightUnit)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update weight unit data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update weight unit data"
	}

	helper.SetResponse(res, req, response)
}

func (h *WeightUnitHandler) Delete(res http.ResponseWriter, req *http.Request) {

	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()

	err := h.WUUsecase.Delete(ctx, id)

	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete fishing gear data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete fishing gear data"
	}

	helper.SetResponse(res, req, response)
}
