package dataaccess

import (
	"database/sql"
	"errors"
)

type GetDeviceStateParams struct {
	DeviceID int `json:"device_id"`
}

type GetDeviceStateResponse struct {
	DeviceID  int    `json:"device_id"`
	Timestamp string `json:"timestamp"`
	State     string `json:"state"`
}

func GetDeviceState(conn *sql.DB, params GetDeviceStateParams) (GetDeviceStateResponse, error) {
	var response GetDeviceStateResponse

	row := conn.QueryRow(
		`SELECT * FROM get_device_state($1)`,
		params.DeviceID,
	)

	err := row.Scan(
		&response.DeviceID,
		&response.Timestamp,
		&response.State,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return GetDeviceStateResponse{}, ErrDeviceStateNotFound
	case err != nil:
		return GetDeviceStateResponse{}, err
	}

	return response, nil
}
