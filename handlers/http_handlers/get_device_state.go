package http_handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pmoura-dev/hauto.transaction-service/dataaccess"
)

func (h *HandlerWithDB) GetDeviceState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceID, err := strconv.Atoi(vars["deviceID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	state, err := dataaccess.GetDeviceState(h.Conn, deviceID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(state))
}
