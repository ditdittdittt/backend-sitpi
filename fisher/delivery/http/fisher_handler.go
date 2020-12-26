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

type FisherHandler struct {
	FUsecase domain.FisherUsecase
}

type FetchFisherRequest struct {
	Cursor string `json:"cursor"`
	Num    int64  `json:"num"`
}

type GetByIDRequest struct {
	ID int64 `json:"id"`
}

type StoreRequest struct {
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UpdateRequest struct {
	ID      int64  `json:"id"`
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DeleteRequest struct {
	ID int64 `json:"id"`
}

func NewFisherHandler(router *mux.Router, uc domain.FisherUsecase) {
	handler := &FisherHandler{FUsecase: uc}
	router.HandleFunc("/fisher", handler.FetchFisher).Methods("GET")
	router.HandleFunc("/fisher/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/fisher", handler.Store).Methods("POST")
	router.HandleFunc("/fisher/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/fisher/{id}", handler.Delete).Methods("DELETE")
}

func (h *FisherHandler) FetchFisher(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listFisher, err := h.FUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch fisher data"
		response.Data = err
	}

	response.Code = "00"
	response.Desc = "Success to fetch fisher data"
	response.Data = listFisher

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	fisher, _ := h.FUsecase.GetByID(ctx, id)

	response.Code = "00"
	response.Desc = "Success to get by ID fisher data"
	response.Data = fisher

	helper.SetResponse(res, req, response)

}

func (h *FisherHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	fisher := &domain.Fisher{
		Nik:       request.Nik,
		Address:   request.Address,
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = h.FUsecase.Store(ctx, fisher)

	response.Code = "00"
	response.Desc = "Success to store fisher data"
	response.Data = err

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	fisher := &domain.Fisher{
		ID:        id,
		Nik:       request.Nik,
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = h.FUsecase.Update(ctx, fisher)

	response.Code = "00"
	response.Desc = "Success to update fisher data"
	response.Data = err

	helper.SetResponse(res, req, response)
}

func (h *FisherHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	err := h.FUsecase.Delete(ctx, id)

	response.Code = "00"
	response.Desc = "Success to delete fisher data"
	response.Data = err

	helper.SetResponse(res, req, response)
}
