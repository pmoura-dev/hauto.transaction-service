package http_handlers

import (
	"database/sql"
)

type HandlerWithDB struct {
	Conn *sql.DB
}
