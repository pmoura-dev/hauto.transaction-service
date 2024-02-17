package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/gobroker/brokers"
	"github.com/pmoura-dev/gobroker/middleware"
	"github.com/pmoura-dev/hauto.transaction-service/config"
	"github.com/pmoura-dev/hauto.transaction-service/database"
	"github.com/pmoura-dev/hauto.transaction-service/handlers/broker_handlers"
	"github.com/pmoura-dev/hauto.transaction-service/handlers/http_handlers"
)

const (
	devicesExchange = "devices"

	deviceStatusQueue = "device.state.transaction-service.queue"
)

func main() {

	db, err := database.GetConnection(config.GetDatabaseConfig())
	if err != nil {
		log.Fatal(err)
	}

	appConfig := config.GetAppConfig()
	r := mux.NewRouter()
	handler := http_handlers.HandlerWithDB{Conn: db}
	r.HandleFunc("/devices/mqtt_configuration", handler.GetDevicesMQTTConfiguration).Methods("GET")
	r.HandleFunc("/devices/{deviceID:[0-9]+}/state", handler.GetDeviceState).Methods("GET")

	go func() {
		if err = http.ListenAndServe(fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port), r); err != nil {
			log.Fatal(err)
		}
	}()

	b := brokers.NewRabbitMQBroker()
	b.AddExchange(devicesExchange)
	b.AddQueue(deviceStatusQueue).Bind(devicesExchange, "state.#")

	s := gobroker.NewServer(b)
	s.Use(middleware.Logging)

	s.AddConsumer(deviceStatusQueue, broker_handlers.UpdateDeviceStatus).AddParam("database", db)

	rabbitMQConfig := config.GetRabbitMQConfig()
	if err = s.Run(
		fmt.Sprintf("amqp://%s:%s@%s:%s",
			rabbitMQConfig.User,
			rabbitMQConfig.Password,
			rabbitMQConfig.Host,
			rabbitMQConfig.Port,
		)); err != nil {
		log.Fatal(err)
	}
}
