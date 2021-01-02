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

type FishTypeHandler struct {
	FTUsecase domain.FishTypeUsecase
}

type StoreRequest struct {
	Name string `json:"name"`
}

type UpdateRequest struct {
	Name string `json:"name"`
}

func NewFishTypeHandler(router *mux.Router, uc domain.FishTypeUsecase) {
	handler := &FishTypeHandler{FTUsecase: uc}
	router.HandleFunc("/fish_type", handler.Fetch).Methods("GET")
	router.HandleFunc("/fish_type", handler.Store).Methods("POST")
	router.HandleFunc("/fish_type/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/fish_type/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/fish_type/{id}", handler.Delete).Methods("DELETE")
}

func (h *FishTypeHandler) Fetch(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listFishType, err := h.FTUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch fish type data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch fish type data"
		response.Data = listFishType
	}

	helper.SetResponse(res, req, response)
}

func (h *FishTypeHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	fishType, err := h.FTUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID fish type data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by id fish type data"
		response.Data = fishType
	}

	helper.SetResponse(res, req, response)
}

func (h *FishTypeHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	fishType := &domain.FishType{
		Name: request.Name,
	}

	err = h.FTUsecase.Store(ctx, fishType)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to store fish type data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store fish type data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FishTypeHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	fishType := &domain.FishType{
		ID:   id,
		Name: request.Name,
	}

	err = h.FTUsecase.Update(ctx, fishType)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update fish type data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update fish type data"
	}

	helper.SetResponse(res, req, response)
}

func (h *FishTypeHandler) Delete(res http.ResponseWriter, req *http.Request) {

	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()

	err := h.FTUsecase.Delete(ctx, id)

	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete fish type data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete fish type data"
	}

	helper.SetResponse(res, req, response)
}
