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

type GetTotalProductionRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func NewCaughtFishHandler(router *mux.Router, uc domain.CaughtFishUsecase) {
	handler := &CaughtFishHandler{CFUsecase: uc}
	router.HandleFunc("/caught_fish", handler.FetchCaughtFish).Methods("GET")
	router.HandleFunc("/caught_fish/total_fisher", handler.GetTotalFisher).Methods("GET")
	router.HandleFunc("/caught_fish/total_production", handler.GetTotalProduction).Methods("GET")
	router.HandleFunc("/caught_fish", handler.Store).Methods("POST")
	router.HandleFunc("/caught_fish/{id:[0-9]+}", handler.GetByID).Methods("GET")
	router.HandleFunc("/caught_fish/{id:[0-9]+}", handler.Update).Methods("PUT")
	router.HandleFunc("/caught_fish/{id:[0-9]+}", handler.Delete).Methods("DELETE")
}

func (h *CaughtFishHandler) FetchCaughtFish(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	fromParam := req.URL.Query()["from"]
	if len(fromParam) == 0 {
		fromParam = append(fromParam, "")
	}

	toParam := req.URL.Query()["to"]
	if len(toParam) == 0 {
		toParam = append(toParam, "")
	}

	fisherParam := req.URL.Query()["fisher_id"]
	if len(fisherParam) == 0 {
		fisherParam = append(fisherParam, "0")
	}

	fishTypeParam := req.URL.Query()["fish_type_id"]
	if len(fishTypeParam) == 0 {
		fishTypeParam = append(fishTypeParam, "0")
	}

	ctx := req.Context()
	request := &domain.FetchCaughtFishRequest{
		From:       fromParam[0],
		To:         toParam[0],
		FisherID:   fisherParam[0],
		FishTypeID: fishTypeParam[0],
	}
	listCaughtFish, err := h.CFUsecase.Fetch(ctx, request)
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
	request := &domain.StoreCaughtFishRequest{}
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
	err = h.CFUsecase.Store(ctx, request)
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
	request := &domain.UpdateCaughtFishRequest{}
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	body, err := helper.ReadRequest(req, response)
	if err != nil {
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		helper.SetResponse(res, req, response)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		helper.SetResponse(res, req, response)
		return
	}

	err = helper.ValidateRequest(request, response)
	if err != nil {
		response.Code = "XX"
		response.Data = "Failed to store caught fish data"
		response.Data = err.Error()
		helper.SetResponse(res, req, response)
		return
	}

	ctx := req.Context()
	err = h.CFUsecase.Update(ctx, id, request)

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
		helper.SetResponse(res, req, response)
		return
	}

	toParam := req.URL.Query()["to"]
	if len(toParam) == 0 || toParam[0] == "" {
		response.Code = "XX"
		response.Desc = "Missing to parameter"
		helper.SetResponse(res, req, response)
		return
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

func (h *CaughtFishHandler) GetTotalFisher(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	fromParam := req.URL.Query()["from"]
	if len(fromParam) == 0 {
		response.Code = "XX"
		response.Desc = "Missing from parameter"
		helper.SetResponse(res, req, response)
		return
	}

	toParam := req.URL.Query()["to"]
	if len(toParam) == 0 {
		response.Code = "XX"
		response.Desc = "Missing to parameter"
		helper.SetResponse(res, req, response)
		return
	}

	ctx := req.Context()
	totalFisher, err := h.CFUsecase.GetTotalFisher(ctx, fromParam[0], toParam[0])
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get total fisher caught fish"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get total fisher caught fish"
		response.Data = totalFisher
	}

	helper.SetResponse(res, req, response)
}
