package dataaccess

import (
	"database/sql"
)

func UpsertDeviceState(conn *sql.DB, deviceID string, state string) error {
	_, err := conn.Exec(
		`SELECT * FROM upsert_device_state($1, $2);`,
		deviceID,
		state,
	)
	return err
}
