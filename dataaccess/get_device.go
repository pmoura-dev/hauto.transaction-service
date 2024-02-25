package dataaccess

import (
	"database/sql"
	"errors"
)

type GetDeviceParams struct {
	DeviceID int `json:"device_id"`
}

type GetDeviceResponse struct {
	ID           int     `json:"id"`
	NaturalID    string  `json:"natural_id"`
	Name         string  `json:"name"`
	Manufacturer *string `json:"manufacturer"`
	Model        *string `json:"model"`
	RoomID       int     `json:"room_id"`
	Controller   string  `json:"controller"`
}

func GetDevice(conn *sql.DB, params GetDeviceParams) (GetDeviceResponse, error) {
	var response GetDeviceResponse

	row := conn.QueryRow(
		`SELECT * FROM get_device($1)`,
		params.DeviceID,
	)

	err := row.Scan(
		&response.ID,
		&response.NaturalID,
		&response.Name,
		&response.Manufacturer,
		&response.Model,
		&response.RoomID,
		&response.Controller,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return GetDeviceResponse{}, ErrDeviceNotFound
	case err != nil:
		return GetDeviceResponse{}, err
	}

	return response, nil
}
