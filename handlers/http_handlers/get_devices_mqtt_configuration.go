package http_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pmoura-dev/hauto.transaction-service/dataaccess"
)

type deviceMQTTConfiguration struct {
	Controller string            `json:"controller"`
	Actions    map[string]string `json:"actions"`
	Listeners  map[string]string `json:"listeners"`
}

func (h *HandlerWithDB) GetDevicesMQTTConfiguration(w http.ResponseWriter, r *http.Request) {

	mqttConfigurationRows, err := dataaccess.GetDevicesMQTTConfiguration(h.Conn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := buildResponse(mqttConfigurationRows)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(2)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func buildResponse(rows []dataaccess.DevicesMQTTConfigurationRow) ([]byte, error) {
	response := make(map[int]deviceMQTTConfiguration)

	for _, row := range rows {
		if _, exists := response[row.DeviceID]; !exists {
			response[row.DeviceID] = deviceMQTTConfiguration{
				Controller: row.Controller,
				Actions:    map[string]string{},
				Listeners:  map[string]string{},
			}
		}

		switch row.ConsumerType {
		case "action":
			actions := response[row.DeviceID].Actions
			actions[row.Name] = row.Topic
		case "listener":
			listeners := response[row.DeviceID].Listeners
			listeners[row.Name] = row.Topic
		}
	}

	return json.Marshal(response)
}
