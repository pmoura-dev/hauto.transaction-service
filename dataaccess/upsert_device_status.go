package dataaccess

import (
	"database/sql"
)

func UpsertDeviceStatus(conn *sql.DB, naturalID string, state string) error {
	_, err := conn.Exec(
		`SELECT * FROM upsert_device_status($1, $2);`,
		naturalID,
		state,
	)
	return err
}
