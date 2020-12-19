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

type ResponseError struct {
	Message string `json:"message"`
}

type CaughtFishHandler struct {
	CFUsecase domain.CaughtFishUsecase
}

type FetchCaughtFishRequest struct {
	Cursor string `json:"cursor"`
	Num    int64  `json:"num"`
}

type GetByIDRequest struct {
	ID int64 `json:"id"`
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

type DeleteRequest struct {
	ID int64 `json:"id"`
}

func NewCaughtFishHandler(router *mux.Router, uc domain.CaughtFishUsecase) {
	handler := &CaughtFishHandler{CFUsecase: uc}
	router.HandleFunc("/caught_fish/index", handler.FetchCaughtFish).Methods("GET")
	router.HandleFunc("/caught_fish/get_by_id", handler.GetByID).Methods("GET")
	router.HandleFunc("/caught_fish/store", handler.Store).Methods("POST")
	router.HandleFunc("/caught_fish/update", handler.Update).Methods("PUT")
	router.HandleFunc("/caught_fish/delete", handler.Delete).Methods("DELETE")
}

func (h *CaughtFishHandler) FetchCaughtFish(res http.ResponseWriter, req *http.Request) {
	request := &FetchCaughtFishRequest{}
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
	listCaughtFish, _, err := h.CFUsecase.Fetch(ctx, request.Cursor, request.Num)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch caught fish data"
		response.Data = err
	}

	response.Code = "00"
	response.Desc = "Success to fetch caught fish data"
	response.Data = listCaughtFish

	helper.SetResponse(res, req, response)
}

func (h *CaughtFishHandler) GetByID(res http.ResponseWriter, req *http.Request) {
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
	caughtFish, err := h.CFUsecase.GetByID(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to get by ID caught fish data"
	response.Data = caughtFish

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

	response.Code = "00"
	response.Desc = "Success to store caught fish data"
	response.Data = caughtFish

	helper.SetResponse(res, req, response)
}

func (h *CaughtFishHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	caughtFish := &domain.CaughtFish{
		ID:          request.ID,
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

	response.Code = "00"
	response.Desc = "Success to update caught fish data"
	response.Data = caughtFish

	helper.SetResponse(res, req, response)
}

func (h *CaughtFishHandler) Delete(res http.ResponseWriter, req *http.Request) {
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
	err = h.CFUsecase.Delete(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to delete caught fish data"

	helper.SetResponse(res, req, response)
}
