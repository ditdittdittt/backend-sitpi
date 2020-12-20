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

type TransactionHandler struct {
	TUsecase domain.TransactionUsecase
}

type FetchTransactionRequest struct {
	Cursor string `json:"cursor"`
	Num    int64  `json:"num"`
}

type GetByIDRequest struct {
	ID int64 `json:"id"`
}

type StoreRequest struct {
	TpiID            int64  `json:"tpi_id"`
	AuctionID        int64  `json:"auction_id"`
	OfficerID        int64  `json:"officer_id"`
	BuyerID          int64  `json:"buyer_id"`
	DistributionArea string `json:"distribution_area"`
}

type UpdateRequest struct {
	ID               int64  `json:"id"`
	TpiID            int64  `json:"tpi_id"`
	AuctionID        int64  `json:"auction_id"`
	OfficerID        int64  `json:"officer_id"`
	BuyerID          int64  `json:"buyer_id"`
	DistributionArea string `json:"distribution_area"`
}

type DeleteRequest struct {
	ID int64 `json:"id"`
}

func NewTransactionHandler(router *mux.Router, uc domain.TransactionUsecase) {
	handler := &TransactionHandler{TUsecase: uc}
	router.HandleFunc("/transaction/index", handler.FetchTransaction).Methods("GET")
	router.HandleFunc("/transaction/get_by_id", handler.GetByID).Methods("GET")
	router.HandleFunc("/transaction/store", handler.Store).Methods("POST")
	router.HandleFunc("/transaction/update", handler.Update).Methods("PUT")
	router.HandleFunc("/transaction/delete", handler.Delete).Methods("DELETE")
}

func (h *TransactionHandler) FetchTransaction(res http.ResponseWriter, req *http.Request) {
	request := &FetchTransactionRequest{}
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
	listTransaction, _, err := h.TUsecase.Fetch(ctx, request.Cursor, request.Num)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch transaction data"
		response.Data = err
	}

	response.Code = "00"
	response.Desc = "Success to fetch transaction data"
	response.Data = listTransaction

	helper.SetResponse(res, req, response)
}

func (h TransactionHandler) GetByID(res http.ResponseWriter, req *http.Request) {
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
	transaction, err := h.TUsecase.GetByID(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to get by ID transaction data"
	response.Data = transaction

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
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err = h.TUsecase.Store(ctx, transaction)

	response.Code = "00"
	response.Desc = "Success to store transaction data"
	response.Data = transaction

	helper.SetResponse(res, req, response)
}

func (h *TransactionHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	transaction := &domain.Transaction{
		TpiID:            request.TpiID,
		AuctionID:        request.AuctionID,
		OfficerID:        request.OfficerID,
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
		UpdatedAt:        time.Now(),
	}

	err = h.TUsecase.Update(ctx, transaction)

	response.Code = "00"
	response.Desc = "Success to update transaction data"
	response.Data = transaction

	helper.SetResponse(res, req, response)
}

func (h *TransactionHandler) Delete(res http.ResponseWriter, req *http.Request) {
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
	err = h.TUsecase.Delete(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to delete transaction data"

	helper.SetResponse(res, req, response)
}
