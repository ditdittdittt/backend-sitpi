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

type CaughtFishHandler struct {
	CFUsecase domain.CaughtFishUsecase
}

type StoreRequest struct {
	FisherID      int64   `json:"fisher_id" validate:"required"`
	FishTypeID    int64   `json:"fish_type_id" validate:"required"`
	Weight        float64 `json:"weight" validate:"required"`
	WeightUnitID  int64   `json:"weight_unit_id" validate:"required"`
	FishingGearID int64   `json:"fishing_gear_id" validate:"required"`
	FishingArea   string  `json:"fishing_area" validate:"required"`

	AuctionWeight       float64 `json:"auction_weight"`
	AuctionWeightUnitID int64   `json:"auction_weight_unit_id"`
}

type GetTotalProductionRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type UpdateRequest struct {
	FisherID      int64   `json:"fisher_id" validate:"required"`
	FishTypeID    int64   `json:"fish_type_id" validate:"required"`
	Weight        float64 `json:"weight" validate:"required"`
	WeightUnitID  int64   `json:"weight_unit_id" validate:"required"`
	FishingGearID int64   `json:"fishing_gear_id" validate:"required"`
	FishingArea   string  `json:"fishing_area" validate:"required"`
}

func NewCaughtFishHandler(router *mux.Router, uc domain.CaughtFishUsecase) {
	handler := &CaughtFishHandler{CFUsecase: uc}
	router.HandleFunc("/caught_fish", handler.FetchCaughtFish).Methods("GET")
	router.HandleFunc("/caught_fish/total_production", handler.GetTotalProduction).Methods("GET")
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
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		logrus.Error(err)
		helper.SetResponse(res, req, response)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		response.Data = err.Error()
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		logrus.Error(err)
		helper.SetResponse(res, req, response)
		return
	}

	err = helper.ValidateRequest(request, response)
	if err != nil {
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		logrus.Error(err)
		helper.SetResponse(res, req, response)
		return
	}

	ctx := req.Context()
	caughtFish := &domain.CaughtFish{
		FisherID:      request.FisherID,
		FishTypeID:    request.FishTypeID,
		Weight:        request.Weight,
		WeightUnitID:  request.WeightUnitID,
		FishingGearID: request.FishingGearID,
		FishingArea:   request.FishingArea,
	}

	auction := &domain.Auction{
		Weight:       request.AuctionWeight,
		WeightUnitID: request.AuctionWeightUnitID,
	}

	err = h.CFUsecase.Store(ctx, caughtFish, auction)
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
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		logrus.Error(err)
		helper.SetResponse(res, req, response)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		logrus.Error(err)
		helper.SetResponse(res, req, response)
		return
	}

	err = helper.ValidateRequest(request, response)
	if err != nil {
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		logrus.Error(err)
		helper.SetResponse(res, req, response)
		return
	}

	ctx := req.Context()
	caughtFish := &domain.CaughtFish{
		ID:            id,
		FisherID:      request.FisherID,
		FishTypeID:    request.FishTypeID,
		Weight:        request.Weight,
		WeightUnitID:  request.WeightUnitID,
		FishingGearID: request.FishingGearID,
		FishingArea:   request.FishingArea,
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

func (h *CaughtFishHandler) GetTotalProduction(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	fromParam := req.URL.Query()["from"]
	if len(fromParam) == 0 || fromParam[0] == "" {
		response.Code = "XX"
		response.Desc = "Missing from parameter"
	}

	toParam := req.URL.Query()["to"]
	if len(toParam) == 0 || toParam[0] == "" {
		response.Code = "XX"
		response.Desc = "Missing to parameter"
	}

	ctx := req.Context()
	totalProduction, err := h.CFUsecase.GetTotalProduction(ctx, fromParam[0], toParam[0])
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get total production caught fish"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get total production caught fish"
		response.Data = totalProduction
	}

	helper.SetResponse(res, req, response)
}
