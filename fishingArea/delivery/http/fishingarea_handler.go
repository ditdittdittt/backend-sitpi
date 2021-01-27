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

type FishingAreaHandler struct {
	FAUsecase domain.FishingAreaUsecase
}

func NewFishingAreaHandler(router *mux.Router, uc domain.FishingAreaUsecase) {
	handler := &FishingAreaHandler{FAUsecase: uc}
	router.HandleFunc("/fishing_area", handler.FetchFishingArea).Methods("GET")
	router.HandleFunc("/fishing_area/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/fishing_area", handler.Store).Methods("POST")
	router.HandleFunc("/fishing_area/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/fishing_area/{id}", handler.Delete).Methods("DELETE")
}

func (h *FishingAreaHandler) FetchFishingArea(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listFishingArea, err := h.FAUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch fishing area data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch fishing area data"
		response.Data = listFishingArea
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingAreaHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	fishingArea, err := h.FAUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID fishing area data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by ID fishing area data"
		response.Data = fishingArea
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingAreaHandler) Store(res http.ResponseWriter, req *http.Request) {
	request := &domain.StoreFishingAreaRequest{}
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
	err = h.FAUsecase.Store(ctx, request)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to store fishing area data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store fishing area data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingAreaHandler) Update(res http.ResponseWriter, req *http.Request) {
	request := &domain.UpdateFishingAreaRequest{}
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
	err = h.FAUsecase.Update(ctx, id, request)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update fishing area data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update fishing area data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FishingAreaHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()

	err := h.FAUsecase.Delete(ctx, id)

	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete fishing area data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete fishing area data"
	}

	helper.SetResponse(res, req, response)
}
