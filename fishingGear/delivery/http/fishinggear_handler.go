package http

import (
	"encoding/json"
	"github.com/ditdittdittt/backend-sitpi/domain"
	_response "github.com/ditdittdittt/backend-sitpi/domain/response"
	"github.com/ditdittdittt/backend-sitpi/helper"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type FishingGearHandler struct {
	FGUsecase domain.FishingGearUsecase
}

type StoreRequest struct {
	Name string `json:"name"`
}

type UpdateRequest struct {
	Name string `json:"name"`
}

func NewFishingGearHandler(router *mux.Router, uc domain.FishingGearUsecase) {
	handler := &FishingGearHandler{FGUsecase: uc}
	router.HandleFunc("/fishing_gear", handler.Fetch).Methods("GET")
	router.HandleFunc("/fishing_gear", handler.Store).Methods("POST")
	router.HandleFunc("/fishing_gear/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/fishing_gear/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/fishing_gear/{id}", handler.Delete).Methods("DELETE")
}

func (h *FishingGearHandler) Fetch(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listFishingGear, err := h.FGUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch fishing gear data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch fishing gear data"
		response.Data = listFishingGear
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingGearHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	fishingGear, err := h.FGUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID fishing gear data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by id fishing gear data"
		response.Data = fishingGear
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingGearHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	fishingGear := &domain.FishingGear{
		Name: request.Name,
	}

	err = h.FGUsecase.Store(ctx, fishingGear)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to store fishing gear data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store fishing gear data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingGearHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	fishingGear := &domain.FishingGear{
		ID:   id,
		Name: request.Name,
	}

	err = h.FGUsecase.Update(ctx, fishingGear)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update fishing gear data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update fishing gear data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingGearHandler) Delete(res http.ResponseWriter, req *http.Request) {

	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()

	err := h.FGUsecase.Delete(ctx, id)

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
