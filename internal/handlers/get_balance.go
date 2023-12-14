package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SylvanasGr/goapi/api"
	"github.com/SylvanasGr/goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.BalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.BalanceDetails
	tokenDetails = (*database).GetUserBalance(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.BalanceResponse{
		Balance: (*tokenDetails).Balance,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
