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

type TransactionHandler struct {
	TUsecase domain.TransactionUsecase
}

type StoreRequest struct {
	AuctionID        int64  `json:"auction_id" validate:"required"`
	BuyerID          int64  `json:"buyer_id" validate:"required"`
	DistributionArea string `json:"distribution_area" validate:"required"`
	Price            int64  `json:"price" validate:"required"`
}

type UpdateRequest struct {
	AuctionID        int64  `json:"auction_id" validate:"required"`
	BuyerID          int64  `json:"buyer_id" validate:"required"`
	DistributionArea string `json:"distribution_area" validate:"required"`
	Price            int64  `json:"price" validate:"required"`
}

func NewTransactionHandler(router *mux.Router, uc domain.TransactionUsecase) {
	handler := &TransactionHandler{TUsecase: uc}
	router.HandleFunc("/transaction", handler.FetchTransaction).Methods("GET")
	router.HandleFunc("/transaction/total_buyer", handler.GetTotalBuyer).Methods("GET")
	router.HandleFunc("/transaction/{id:[0-9]+}", handler.GetByID).Methods("GET")
	router.HandleFunc("/transaction", handler.Store).Methods("POST")
	router.HandleFunc("/transaction/{id:[0-9]+}", handler.Update).Methods("PUT")
	router.HandleFunc("/transaction/{id:[0-9]+}", handler.Delete).Methods("DELETE")
}

func (h *TransactionHandler) FetchTransaction(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listTransaction, err := h.TUsecase.Fetch(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch transaction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch transaction data"
		response.Data = listTransaction
	}

	helper.SetResponse(res, req, response)
}

func (h TransactionHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	transaction, err := h.TUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID transaction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by ID transaction data"
		response.Data = transaction
	}

	helper.SetResponse(res, req, response)
}

func (h *TransactionHandler) Store(res http.ResponseWriter, req *http.Request) {
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
	transaction := &domain.Transaction{
		AuctionID:        request.AuctionID,
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
		Price:            request.Price,
	}

	err = h.TUsecase.Store(ctx, transaction)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to store transaction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to store transaction data"
	}

	helper.SetResponse(res, req, response)
}

func (h *TransactionHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	transaction := &domain.Transaction{
		ID:               id,
		AuctionID:        request.AuctionID,
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
		Price:            request.Price,
	}

	err = h.TUsecase.Update(ctx, transaction)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to update transaction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to update transaction data"
	}

	helper.SetResponse(res, req, response)
}

func (h *TransactionHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	err := h.TUsecase.Delete(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete transaction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete transaction data"
	}

	helper.SetResponse(res, req, response)
}

func (h *TransactionHandler) GetTotalBuyer(res http.ResponseWriter, req *http.Request) {
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
	totalFisher, err := h.TUsecase.GetTotalBuyer(ctx, fromParam[0], toParam[0])
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
