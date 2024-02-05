package main

import (
	"log"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/gobroker/brokers"
	"github.com/pmoura-dev/gobroker/middleware"
	"github.com/pmoura-dev/hauto.transaction-service/config"
	"github.com/pmoura-dev/hauto.transaction-service/database"
	"github.com/pmoura-dev/hauto.transaction-service/handlers"
)

const (
	devicesExchange = "devices"

	deviceStatusQueue = "device.status.transaction-service.queue"
)

func main() {

	db, err := database.GetConnection(config.GetDatabaseConfig())
	if err != nil {
		log.Fatal(err)
	}

	b := brokers.NewRabbitMQBroker()
	b.AddExchange(devicesExchange)
	b.AddQueue(deviceStatusQueue).Bind(devicesExchange, "status.#")

	s := gobroker.NewServer(b)
	s.Use(middleware.Logging)

	s.AddConsumer(deviceStatusQueue, handlers.UpdateDeviceStatus).AddParam("database", db)

	if err = s.Run("amqp://guest:guest@localhost:5672"); err != nil {
		log.Fatal(err)
	}
}
