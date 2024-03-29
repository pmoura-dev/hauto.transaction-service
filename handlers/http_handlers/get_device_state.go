package http_handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pmoura-dev/hauto.transaction-service/dataaccess"
)

func (h *HandlerWithDB) GetDeviceState(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var request dataaccess.GetDeviceStateParams
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := dataaccess.GetDeviceState(h.Conn, request)
	switch {
	case errors.Is(err, dataaccess.ErrDeviceStateNotFound):
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse := []byte(fmt.Sprintf(`{"device_id": %d, "timestamp": "%s", "state": %s}`,
		response.DeviceID,
		response.Timestamp,
		response.State,
	))

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}
