package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"

	"github.com/ditdittdittt/backend-sitpi/config"
	_auctionHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/auction/delivery/http"
	_auctionRepo "github.com/ditdittdittt/backend-sitpi/modules/auction/repository/mysql"
	_auctionUsecase "github.com/ditdittdittt/backend-sitpi/modules/auction/usecase"
	_buyerHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/buyer/delivery/http"
	_buyerRepo "github.com/ditdittdittt/backend-sitpi/modules/buyer/repository/mysql"
	_buyerUsecase "github.com/ditdittdittt/backend-sitpi/modules/buyer/usecase"
	_caughtFishHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/caughtFish/delivery/http"
	_caughtFishRepo "github.com/ditdittdittt/backend-sitpi/modules/caughtFish/repository/mysql"
	_caughtFishUsecase "github.com/ditdittdittt/backend-sitpi/modules/caughtFish/usecase"
	_fishTypeHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/fishType/delivery/http"
	_fishTypeRepo "github.com/ditdittdittt/backend-sitpi/modules/fishType/repository/mysql"
	_fishTypeUsecase "github.com/ditdittdittt/backend-sitpi/modules/fishType/usecase"
	_fisherHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/fisher/delivery/http"
	_fisherRepo "github.com/ditdittdittt/backend-sitpi/modules/fisher/repository/mysql"
	_fisherUsecase "github.com/ditdittdittt/backend-sitpi/modules/fisher/usecase"
	_fishingAreaHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/fishingArea/delivery/http"
	_fishingAreaRepo "github.com/ditdittdittt/backend-sitpi/modules/fishingArea/repository/mysql"
	_fishingAreaUsecase "github.com/ditdittdittt/backend-sitpi/modules/fishingArea/usecase"
	_fishingGearHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/fishingGear/delivery/http"
	_fishingGearRepo "github.com/ditdittdittt/backend-sitpi/modules/fishingGear/repository/mysql"
	_fishingGearUsecase "github.com/ditdittdittt/backend-sitpi/modules/fishingGear/usecase"
	_transactionHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/transaction/delivery/http"
	_transactionRepo "github.com/ditdittdittt/backend-sitpi/modules/transaction/repository/mysql"
	_transactionUsecase "github.com/ditdittdittt/backend-sitpi/modules/transaction/usecase"
	_userHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/user/delivery/http"
	_userRepo "github.com/ditdittdittt/backend-sitpi/modules/user/repository/mysql"
	_userUsecase "github.com/ditdittdittt/backend-sitpi/modules/user/usecase"
	_weightUnitHttpDelivery "github.com/ditdittdittt/backend-sitpi/modules/weightUnit/delivery/http"
	_weightUnitRepo "github.com/ditdittdittt/backend-sitpi/modules/weightUnit/repository/mysql"
	_weightUnitUsecase "github.com/ditdittdittt/backend-sitpi/modules/weightUnit/usecase"
)

func main() {
	config.Init()
	dbConn, err := sql.Open(`mysql`, config.Dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	r := mux.NewRouter()

	auctionRepo := _auctionRepo.NewMysqlAuctionRepository(dbConn)
	auctionUsecase := _auctionUsecase.NewAuctionUsecase(auctionRepo, config.TimeoutContext)
	_auctionHttpDelivery.NewAuctionHandler(r, auctionUsecase)

	caughtFishRepo := _caughtFishRepo.NewMysqlCaughtFishRepository(dbConn)
	caughtFishUsecase := _caughtFishUsecase.NewCaughtFishUsecase(caughtFishRepo, auctionRepo, config.TimeoutContext)
	_caughtFishHttpDelivery.NewCaughtFishHandler(r, caughtFishUsecase)

	fisherRepo := _fisherRepo.NewMysqlFisherRepository(dbConn)
	fisherUsecase := _fisherUsecase.NewFisherUsecase(fisherRepo, config.TimeoutContext)
	_fisherHttpDelivery.NewFisherHandler(r, fisherUsecase)

	transactionRepo := _transactionRepo.NewMysqlTransactionRepository(dbConn)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(transactionRepo, auctionRepo, config.TimeoutContext)
	_transactionHttpDelivery.NewTransactionHandler(r, transactionUsecase)

	buyerRepo := _buyerRepo.NewMysqlBuyerRepository(dbConn)
	buyerUsecase := _buyerUsecase.NewBuyerUsecase(buyerRepo, config.TimeoutContext)
	_buyerHttpDelivery.NewBuyerHandler(r, buyerUsecase)

	fishTypeRepo := _fishTypeRepo.NewMysqlFishTypeRepository(dbConn)
	fishTypeUsecase := _fishTypeUsecase.NewFishTypeUsecase(fishTypeRepo, config.TimeoutContext)
	_fishTypeHttpDelivery.NewFishTypeHandler(r, fishTypeUsecase)

	fishingGearRepo := _fishingGearRepo.NewMysqlFishingGearRepository(dbConn)
	fishingGearusecase := _fishingGearUsecase.NewFishingGearUsecase(fishingGearRepo, config.TimeoutContext)
	_fishingGearHttpDelivery.NewFishingGearHandler(r, fishingGearusecase)

	weightUnitRepo := _weightUnitRepo.NewMysqlWeightUnitRepository(dbConn)
	weightUnitUsecase := _weightUnitUsecase.NewWeightUnitUsecase(weightUnitRepo, config.TimeoutContext)
	_weightUnitHttpDelivery.NewWeightUnitHandler(r, weightUnitUsecase)

	fishingAreaRepo := _fishingAreaRepo.NewFishingAreRepository(dbConn)
	fishingAreaUsecase := _fishingAreaUsecase.NewFishingAreaUsecase(fishingAreaRepo, config.TimeoutContext)
	_fishingAreaHttpDelivery.NewFishingAreaHandler(r, fishingAreaUsecase)

	userRepo := _userRepo.NewMysqlUserRepository(dbConn)
	userUsecase := _userUsecase.NewUseUsecase(userRepo, config.TimeoutContext)
	_userHttpDelivery.NewUserHandler(r, userUsecase)

	handler := cors.Default().Handler(r)
	c := cors.New(cors.Options{AllowedMethods: []string{"POST", "GET", "DELETE", "PUT"}})
	handler = c.Handler(handler)

	_ = http.ListenAndServe(viper.GetString("server.address"), handler)
}
