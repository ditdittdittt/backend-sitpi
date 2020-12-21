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

type BuyerHandler struct {
	BUsecase domain.BuyerUsecase
}

type FetchBuyerRequest struct {
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

func NewBuyerHandler(router *mux.Router, uc domain.BuyerUsecase) {

}

func (h *BuyerHandler) FetchBuyer(res http.ResponseWriter, req *http.Request) {
	request := &FetchBuyerRequest{}
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
	listBuyer, _, err := h.BUsecase.Fetch(ctx, request.Cursor, request.Num)
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
	buyer, err := h.BUsecase.GetByID(ctx, request.ID)

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
	response.Data = buyer

	helper.SetResponse(res, req, response)
}

func (h *BuyerHandler) Update(res http.ResponseWriter, req *http.Request) {
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
	buyer := &domain.Buyer{
		ID:        request.ID,
		Nik:       request.Nik,
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = h.BUsecase.Update(ctx, buyer)

	response.Code = "00"
	response.Desc = "Success to update buyer data"
	response.Data = buyer

	helper.SetResponse(res, req, response)
}

func (h *BuyerHandler) Delete(res http.ResponseWriter, req *http.Request) {
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

	err = h.BUsecase.Delete(ctx, request.ID)

	response.Code = "00"
	response.Desc = "Success to delete buyer data"

	helper.SetResponse(res, req, response)
}
