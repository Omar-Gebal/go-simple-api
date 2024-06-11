package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Omar-Gebal/go-simple-api/api"
	log "github.com/sirupsen/logrus"
)

func GetApiDescription(w http.ResponseWriter, r *http.Request) {
	description := api.ApiDescription{
		Version:      "1.0.0",
		Descrtiption: "This is a simple go api, that has accounts and coin balance for each account",
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	var err = json.NewEncoder(w).Encode(description)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
