package http_handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pmoura-dev/hauto.transaction-service/dataaccess"
)

type GetDeviceControlDataResponse struct {
	NaturalID  string `json:"natural_id"`
	Controller string `json:"controller"`
}

func (h *HandlerWithDB) GetDeviceControlData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	deviceID, err := strconv.Atoi(vars["device_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	controlData, err := dataaccess.GetDeviceControlData(h.Conn, deviceID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(GetDeviceControlDataResponse{
		NaturalID:  controlData.NaturalID,
		Controller: controlData.Controller,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
