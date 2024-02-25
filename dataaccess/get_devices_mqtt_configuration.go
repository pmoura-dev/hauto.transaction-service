package dataaccess

import (
	"database/sql"
)

type DevicesMQTTConfigurationRow struct {
	DeviceID     int
	Controller   string
	ConsumerType string
	Name         string
	Topic        string
}

type GetDevicesMQTTConfigurationResponse []DevicesMQTTConfigurationRow

func GetDevicesMQTTConfiguration(conn *sql.DB) (GetDevicesMQTTConfigurationResponse, error) {

	rows, err := conn.Query(`SELECT * FROM get_devices_mqtt_configuration();`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mqttConfigurationRows []DevicesMQTTConfigurationRow
	for rows.Next() {
		var mqttConfiguration DevicesMQTTConfigurationRow
		err := rows.Scan(
			&mqttConfiguration.DeviceID,
			&mqttConfiguration.Controller,
			&mqttConfiguration.ConsumerType,
			&mqttConfiguration.Name,
			&mqttConfiguration.Topic,
		)
		if err != nil {
			return nil, err
		}

		mqttConfigurationRows = append(mqttConfigurationRows, mqttConfiguration)
	}

	return mqttConfigurationRows, nil
}
