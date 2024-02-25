package http_handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/pmoura-dev/hauto.transaction-service/dataaccess"
)

type GetDeviceRequest struct {
	DeviceID int `json:"device_id"`
}

func (h *HandlerWithDB) GetDevice(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var request dataaccess.GetDeviceParams
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := dataaccess.GetDevice(h.Conn, request)
	switch {
	case errors.Is(err, dataaccess.ErrDeviceNotFound):
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	case err != nil:
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}
