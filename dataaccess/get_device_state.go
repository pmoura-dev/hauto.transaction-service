package dataaccess

import (
	"database/sql"
)

func GetDeviceState(conn *sql.DB, deviceID int) (string, error) {
	var state string

	row := conn.QueryRow(
		`SELECT * FROM get_device_state($1);`,
		deviceID,
	)

	err := row.Scan(&state)
	if err != nil {
		return "", err
	}

	return state, nil
}
