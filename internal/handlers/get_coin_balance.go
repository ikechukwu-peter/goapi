package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/ikechukwu-peter/goapi/api"
	"github.com/ikechukwu-peter/goapi/internal/tools"
	"github.com/ikechukwu-peter/internal/tools"
	log "github.com/sirupsen/logrus"
)


func GetCoinBalance( w http.ResponseWriter, r r *http.Request){
	var params = api.CoinBalanceParams{}
	var decoder  *schema.Decode = schema.NewDecoder()

	var err error
	err =  decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return 
	}


	var database *tools.DatabaseInterface

	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return 
	}


	var tokenDetails *tools.CoinDetails

	tokenDetails = (*database).GetUserCoins(params.username)

	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return 
		}


		var response = api.CoinBalanceResponse{
			Balance: (*tokenDetails).Coins,
			Code: http.StatusOK
		}

		w.Header().Set("Content-Type", "application/json")
		err= json.NewEncoder(w).Encode((response))
		if tokenDetails == nil {
			log.Error(err)
			api.InternalErrorHandler(w)
			return 
			}
	

}