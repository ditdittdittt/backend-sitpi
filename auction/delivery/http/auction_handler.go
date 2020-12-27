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

type AuctionHandler struct {
	AUsecase domain.AuctionUsecase
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

func NewAuctionHandler(router *mux.Router, uc domain.AuctionUsecase) {
	handler := &AuctionHandler{AUsecase: uc}
	router.HandleFunc("/auction", handler.FetchAuction).Methods("GET")
	router.HandleFunc("/auction/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/auction", handler.Store).Methods("POST")
	router.HandleFunc("/auction/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/auction/{id}", handler.Delete).Methods("DELETE")
}

func (h *AuctionHandler) FetchAuction(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listAuction, err := h.AUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch auction data"
		response.Data = listAuction
	}

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	auction, err := h.AUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by ID auction data"
		response.Data = auction
	}

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
	}

	err = h.AUsecase.Store(ctx, auction)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to store auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store auction data"
		response.Data = auction
	}

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	auction := &domain.Auction{
		ID:          id,
		TpiID:       request.TpiID,
		FisherID:    request.FisherID,
		OfficerID:   request.OfficerID,
		FishTypeID:  request.FishTypeID,
		Weight:      request.Weight,
		WeightUnit:  request.WeightUnit,
		FishingGear: request.FishingGear,
		FishingArea: request.FishingArea,
		Status:      request.Status,
	}

	err = h.AUsecase.Update(ctx, auction)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update auction data"
		response.Data = auction
	}

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	err := h.AUsecase.Delete(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete auction data"
	}

	helper.SetResponse(res, req, response)
}
