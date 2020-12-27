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
	"time"
)

type CaughtFishHandler struct {
	CFUsecase domain.CaughtFishUsecase
}

type StoreRequest struct {
	TpiID       int64   `json:"tpi_id"`
	OfficerID   int64   `json:"officer_id"`
	FisherID    int64   `json:"fisher_id"`
	FishTypeID  int64   `json:"fish_type_id"`
	Weight      float64 `json:"weight"`
	WeightUnit  string  `json:"weight_unit"`
	FishingGear string  `json:"fishing_gear"`
	FishingArea string  `json:"fishing_area"`
}

type UpdateRequest struct {
	ID          int64   `json:"id"`
	TpiID       int64   `json:"tpi_id"`
	OfficerID   int64   `json:"officer_id"`
	FisherID    int64   `json:"fisher_id"`
	FishTypeID  int64   `json:"fish_type_id"`
	Weight      float64 `json:"weight"`
	WeightUnit  string  `json:"weight_unit"`
	FishingGear string  `json:"fishing_gear"`
	FishingArea string  `json:"fishing_area"`
}

func NewCaughtFishHandler(router *mux.Router, uc domain.CaughtFishUsecase) {
	handler := &CaughtFishHandler{CFUsecase: uc}
	router.HandleFunc("/caught_fish", handler.FetchCaughtFish).Methods("GET")
	router.HandleFunc("/caught_fish/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/caught_fish", handler.Store).Methods("POST")
	router.HandleFunc("/caught_fish/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/caught_fish/{id}", handler.Delete).Methods("DELETE")
}

func (h *CaughtFishHandler) FetchCaughtFish(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listCaughtFish, err := h.CFUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch caught fish data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch caught fish data"
		response.Data = listCaughtFish
	}

	helper.SetResponse(res, req, response)
}

func (h *CaughtFishHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	caughtFish, err := h.CFUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID caught fish data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by ID caught fish data"
		response.Data = caughtFish
	}

	helper.SetResponse(res, req, response)
}

func (h *CaughtFishHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	caughtFish := &domain.CaughtFish{
		TpiID:       request.TpiID,
		OfficerID:   request.OfficerID,
		FisherID:    request.FisherID,
		FishTypeID:  request.FishTypeID,
		Weight:      request.Weight,
		WeightUnit:  request.WeightUnit,
		FishingGear: request.FishingGear,
		FishingArea: request.FishingArea,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = h.CFUsecase.Store(ctx, caughtFish)
	if err != nil {
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store caught fish data"
	}

	helper.SetResponse(res, req, response)
}

func (h *CaughtFishHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	caughtFish := &domain.CaughtFish{
		ID:          id,
		TpiID:       request.TpiID,
		OfficerID:   request.OfficerID,
		FisherID:    request.FisherID,
		FishTypeID:  request.FishTypeID,
		Weight:      request.Weight,
		WeightUnit:  request.WeightUnit,
		FishingGear: request.FishingGear,
		FishingArea: request.FishingArea,
		UpdatedAt:   time.Now(),
	}
	err = h.CFUsecase.Update(ctx, caughtFish)

	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update caught fish data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update caught fish data"
	}

	helper.SetResponse(res, req, response)
}

func (h *CaughtFishHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	err := h.CFUsecase.Delete(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete caught fish data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete caught fish data"
	}

	helper.SetResponse(res, req, response)
}
