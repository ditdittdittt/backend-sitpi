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
	TpiID            int64  `json:"tpi_id" validate:"required"`
	AuctionID        int64  `json:"auction_id" validate:"required"`
	OfficerID        int64  `json:"officer_id" validate:"required"`
	BuyerID          int64  `json:"buyer_id" validate:"required"`
	DistributionArea string `json:"distribution_area" validate:"required"`
	Price            int64  `json:"price" validate:"required"`
}

type UpdateRequest struct {
	TpiID            int64  `json:"tpi_id" validate:"required"`
	AuctionID        int64  `json:"auction_id" validate:"required"`
	OfficerID        int64  `json:"officer_id" validate:"required"`
	BuyerID          int64  `json:"buyer_id" validate:"required"`
	DistributionArea string `json:"distribution_area" validate:"required"`
	Price            int64  `json:"price" validate:"required"`
}

func NewTransactionHandler(router *mux.Router, uc domain.TransactionUsecase) {
	handler := &TransactionHandler{TUsecase: uc}
	router.HandleFunc("/transaction", handler.FetchTransaction).Methods("GET")
	router.HandleFunc("/transaction/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/transaction", handler.Store).Methods("POST")
	router.HandleFunc("/transaction/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/transaction/{id}", handler.Delete).Methods("DELETE")
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
		TpiID:            request.TpiID,
		AuctionID:        request.AuctionID,
		OfficerID:        request.OfficerID,
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
		TpiID:            request.TpiID,
		AuctionID:        request.AuctionID,
		OfficerID:        request.OfficerID,
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
