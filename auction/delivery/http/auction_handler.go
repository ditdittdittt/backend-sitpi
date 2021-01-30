package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ditdittdittt/backend-sitpi/domain"
	_response "github.com/ditdittdittt/backend-sitpi/domain/response"
	"github.com/ditdittdittt/backend-sitpi/helper"
)

type AuctionHandler struct {
	AUsecase domain.AuctionUsecase
}

func NewAuctionHandler(router *mux.Router, uc domain.AuctionUsecase) {
	handler := &AuctionHandler{AUsecase: uc}
	router.HandleFunc("/auction", handler.FetchAuction).Methods("GET")
	router.HandleFunc("/auction/inquiry", handler.Inquiry).Methods("GET")
	router.HandleFunc("/auction/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/auction/{id}", handler.Delete).Methods("DELETE")
}

func (h *AuctionHandler) FetchAuction(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	fromParam := req.URL.Query()["from"]
	if len(fromParam) == 0 {
		fromParam = append(fromParam, "")
	}

	toParam := req.URL.Query()["to"]
	if len(toParam) == 0 {
		toParam = append(toParam, "")
	}

	auctionParam := req.URL.Query()["auction_id"]
	if len(auctionParam) == 0 {
		auctionParam = append(auctionParam, "0")
	}

	fisherParam := req.URL.Query()["fisher_id"]
	if len(fisherParam) == 0 {
		fisherParam = append(fisherParam, "0")
	}

	fishTypeParam := req.URL.Query()["fish_type_id"]
	if len(fishTypeParam) == 0 {
		fishTypeParam = append(fishTypeParam, "0")
	}

	statusParam := req.URL.Query()["status_id"]
	if len(statusParam) == 0 {
		statusParam = append(statusParam, "0")
	}

	ctx := req.Context()
	request := &domain.FetchAuctionRequest{
		From:       fromParam[0],
		To:         toParam[0],
		AuctionID:  auctionParam[0],
		FisherID:   fisherParam[0],
		FishTypeID: fishTypeParam[0],
		StatusID:   statusParam[0],
	}
	listAuction, err := h.AUsecase.Fetch(ctx, request)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to fetch auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to fetch auction data"
		response.Data = listAuction
	}

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) GetByID(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	auction, err := h.AUsecase.GetByID(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to get by ID auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to get by ID auction data"
		response.Data = auction
	}

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) Delete(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	params := mux.Vars(req)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	ctx := req.Context()
	err := h.AUsecase.Delete(ctx, id)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to delete auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to delete auction data"
	}

	helper.SetResponse(res, req, response)
}

func (h *AuctionHandler) Inquiry(res http.ResponseWriter, req *http.Request) {
	response := _response.New()

	ctx := req.Context()
	listAuction, err := h.AUsecase.Inquiry(ctx)
	if err != nil {
		response.Code = "XX"
		response.Desc = "Failed to inquiry auction data"
		response.Data = err.Error()
	} else {
		response.Code = "00"
		response.Desc = "Success to inquiry auction data"
		response.Data = listAuction
	}

	helper.SetResponse(res, req, response)
}
