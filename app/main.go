package main

import (
	"database/sql"
	"fmt"
	_auctionHttpDelivery "github.com/ditdittdittt/backend-sitpi/auction/delivery/http"
	_auctionRepo "github.com/ditdittdittt/backend-sitpi/auction/repository/mysql"
	_auctionUsecase "github.com/ditdittdittt/backend-sitpi/auction/usecase"
	_buyerHttpDelivery "github.com/ditdittdittt/backend-sitpi/buyer/delivery/http"
	_buyerRepo "github.com/ditdittdittt/backend-sitpi/buyer/repository/mysql"
	_buyerUsecase "github.com/ditdittdittt/backend-sitpi/buyer/usecase"
	_caughtFishHttpDelivery "github.com/ditdittdittt/backend-sitpi/caughtFish/delivery/http"
	_caughtFishRepo "github.com/ditdittdittt/backend-sitpi/caughtFish/repository/mysql"
	_caughtFishUsecase "github.com/ditdittdittt/backend-sitpi/caughtFish/usecase"
	_fishTypeHttpDelivery "github.com/ditdittdittt/backend-sitpi/fishType/delivery/http"
	_fishTypeRepo "github.com/ditdittdittt/backend-sitpi/fishType/repository/msyql"
	_fishTypeUsecase "github.com/ditdittdittt/backend-sitpi/fishType/usecase"
	_fisherHttpDelivery "github.com/ditdittdittt/backend-sitpi/fisher/delivery/http"
	_fisherRepo "github.com/ditdittdittt/backend-sitpi/fisher/repository/mysql"
	_fisherUsecase "github.com/ditdittdittt/backend-sitpi/fisher/usecase"
	_fishingGearHttpDelivery "github.com/ditdittdittt/backend-sitpi/fishingGear/delivery/http"
	_fishingGearRepo "github.com/ditdittdittt/backend-sitpi/fishingGear/repository/mysql"
	_fishingGearUsecase "github.com/ditdittdittt/backend-sitpi/fishingGear/usecase"
	_transactionHttpDelivery "github.com/ditdittdittt/backend-sitpi/transaction/delivery/http"
	_transactionRepo "github.com/ditdittdittt/backend-sitpi/transaction/repository/mysql"
	_transactionUsecase "github.com/ditdittdittt/backend-sitpi/transaction/usecase"
	_weightUnitHttpDelivery "github.com/ditdittdittt/backend-sitpi/weightUnit/delivery/http"
	_weightUnitRepo "github.com/ditdittdittt/backend-sitpi/weightUnit/repository/mysql"
	_weightUnitUsecase "github.com/ditdittdittt/backend-sitpi/weightUnit/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

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

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	r := mux.NewRouter()

	auctionRepo := _auctionRepo.NewMysqlAuctionRepository(dbConn)
	auctionUsecase := _auctionUsecase.NewAuctionUsecase(auctionRepo, timeoutContext)
	_auctionHttpDelivery.NewAuctionHandler(r, auctionUsecase)

	caughtFishRepo := _caughtFishRepo.NewMysqlCaughtFishRepository(dbConn)
	caughtFishUsecase := _caughtFishUsecase.NewCaughtFishUsecase(caughtFishRepo, auctionRepo, timeoutContext)
	_caughtFishHttpDelivery.NewCaughtFishHandler(r, caughtFishUsecase)

	fisherRepo := _fisherRepo.NewMysqlFisherRepository(dbConn)
	fisherUsecase := _fisherUsecase.NewFisherUsecase(fisherRepo, timeoutContext)
	_fisherHttpDelivery.NewFisherHandler(r, fisherUsecase)

	transactionRepo := _transactionRepo.NewMysqlTransactionRepository(dbConn)
	transactionUsecase := _transactionUsecase.NewTransactionUsecase(transactionRepo, auctionRepo, timeoutContext)
	_transactionHttpDelivery.NewTransactionHandler(r, transactionUsecase)

	buyerRepo := _buyerRepo.NewMysqlBuyerRepository(dbConn)
	buyerUsecase := _buyerUsecase.NewBuyerUsecase(buyerRepo, timeoutContext)
	_buyerHttpDelivery.NewBuyerHandler(r, buyerUsecase)

	fishTypeRepo := _fishTypeRepo.NewMysqlFishTypeRepository(dbConn)
	fishTypeUsecase := _fishTypeUsecase.NewFishTypeUsecase(fishTypeRepo, timeoutContext)
	_fishTypeHttpDelivery.NewFishTypeHandler(r, fishTypeUsecase)

	fishingGearRepo := _fishingGearRepo.NewMysqlFishingGearRepository(dbConn)
	fishingGearusecase := _fishingGearUsecase.NewFishingGearUsecase(fishingGearRepo, timeoutContext)
	_fishingGearHttpDelivery.NewFishingGearHandler(r, fishingGearusecase)

	weightUnitRepo := _weightUnitRepo.NewMysqlWeightUnitRepository(dbConn)
	weightUnitUsecase := _weightUnitUsecase.NewWeightUnitUsecase(weightUnitRepo, timeoutContext)
	_weightUnitHttpDelivery.NewWeightUnitHandler(r, weightUnitUsecase)

	handler := cors.Default().Handler(r)
	c := cors.New(cors.Options{AllowedMethods: []string{"POST", "GET", "DELETE", "PUT"}})
	handler = c.Handler(handler)

	_ = http.ListenAndServe(viper.GetString("server.address"), handler)
}
