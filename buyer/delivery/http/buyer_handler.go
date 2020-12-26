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

type BuyerHandler struct {
	BUsecase domain.BuyerUsecase
}

type StoreRequest struct {
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UpdateRequest struct {
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewBuyerHandler(router *mux.Router, uc domain.BuyerUsecase) {
	handler := &BuyerHandler{BUsecase: uc}
	router.HandleFunc("/buyer", handler.FetchBuyer).Methods("GET")
	router.HandleFunc("/buyer/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/buyer", handler.Store).Methods("POST")
	router.HandleFunc("/buyer/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/buyer/{id}", handler.Delete).Methods("DELETE")
}

func (h *BuyerHandler) FetchBuyer(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listBuyer, err := h.BUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch buyer data"
		response.Data = err
	}

	response.Code = "00"
	response.Desc = "Success to fetch buyer data"
	response.Data = listBuyer

	helper.SetResponse(res, req, response)
}

func (h *BuyerHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	buyer, _ := h.BUsecase.GetByID(ctx, id)

	response.Code = "00"
	response.Desc = "Success to get by ID buyer"
	response.Data = buyer

	helper.SetResponse(res, req, response)
}

func (h *BuyerHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	buyer := &domain.Buyer{
		Nik:       request.Nik,
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.BUsecase.Store(ctx, buyer)

	response.Code = "00"
	response.Desc = "Success to store buyer data"
	response.Data = err

	helper.SetResponse(res, req, response)
}

func (h *BuyerHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	buyer := &domain.Buyer{
		ID:        id,
		Nik:       request.Nik,
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.BUsecase.Update(ctx, buyer)

	response.Code = "00"
	response.Desc = "Success to update buyer data"
	response.Data = err

	helper.SetResponse(res, req, response)
}

func (h *BuyerHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()

	err := h.BUsecase.Delete(ctx, id)

	response.Code = "00"
	response.Desc = "Success to delete buyer data"
	response.Data = err

	helper.SetResponse(res, req, response)
}
