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

type AuctionHandler struct {
	AUsecase domain.AuctionUsecase
}

type FetchAuctionRequest struct {
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
	Price       int     `json:"price"`
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
	Price       int     `json:"price"`
	Status      int     `json:"status"`
}

type DeleteRequest struct {
	ID int64 `json:"id"`
}

func NewAuctionHandler(router *mux.Router, uc domain.AuctionUsecase) {
	handler := &AuctionHandler{AUsecase: uc}
	router.HandleFunc("/auction/index", handler.FetchAuction).Methods("GET")
	router.HandleFunc("/auction/get_by_id", handler.GetByID).Methods("GET")
	router.HandleFunc("/auction/store", handler.Store).Methods("POST")
	router.HandleFunc("/auction/update", handler.Update).Methods("PUT")
	router.HandleFunc("/auction/delete", handler.Delete).Methods("DELETE")
}

func (h *AuctionHandler) FetchAuction(res http.ResponseWriter, req *http.Request) {
	request := &FetchAuctionRequest{}
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
	listAuction, _, err := h.AUsecase.Fetch(ctx, request.Cursor, request.Num)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch auction data"
		response.Data = err
	}

	response.Code = "00"
	response.Desc = "Success to fetch caught fish data"
	response.Data = listAuction

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) GetByID(res http.ResponseWriter, req *http.Request) {
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
	auction, err := h.AUsecase.GetByID(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to get by ID auction data"
	response.Data = auction

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	auction := &domain.Auction{
		TpiID:       request.TpiID,
		FisherID:    request.FisherID,
		OfficerID:   request.OfficerID,
		FishTypeID:  request.FishTypeID,
		Weight:      request.Weight,
		WeightUnit:  request.WeightUnit,
		FishingGear: request.FishingGear,
		FishingArea: request.FishingArea,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = h.AUsecase.Store(ctx, auction)

	response.Code = "00"
	response.Desc = "Success to store auction data"
	response.Data = auction

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	auction := &domain.Auction{
		ID:          request.ID,
		TpiID:       request.TpiID,
		FisherID:    request.FisherID,
		OfficerID:   request.OfficerID,
		FishTypeID:  request.FishTypeID,
		Weight:      request.Weight,
		WeightUnit:  request.WeightUnit,
		FishingGear: request.FishingGear,
		FishingArea: request.FishingArea,
		Status:      request.Status,
		UpdatedAt:   time.Now(),
	}
	err = h.AUsecase.Update(ctx, auction)

	response.Code = "00"
	response.Desc = "Success to update auction data"
	response.Data = auction

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) Delete(res http.ResponseWriter, req *http.Request) {
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
	err = h.AUsecase.Delete(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to delete auction data"

	helper.SetResponse(res, req, response)
}
