package dataaccess

import (
	"database/sql"
)

type DeviceControlDataRow struct {
	NaturalID  string
	Controller string
}

func GetDeviceControlData(conn *sql.DB, deviceID int) (DeviceControlDataRow, error) {
	var controlData DeviceControlDataRow

	row := conn.QueryRow(`SELECT * FROM get_device_control_data($1)`, deviceID)
	err := row.Scan(&controlData.NaturalID, &controlData.Controller)
	if err != nil {
		return DeviceControlDataRow{}, err
	}

	return controlData, nil
}
