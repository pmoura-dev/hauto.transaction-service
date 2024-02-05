package handlers

import (
	"database/sql"
	"regexp"

	"github.com/pmoura-dev/gobroker"
)

var (
	topicRE = regexp.MustCompile(`status.([\w\-]+)`)
)

func UpdateDeviceStatus(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	dbConnection := ctx.Params["database"].(*sql.DB)

	naturalID := topicRE.FindStringSubmatch(message.GetTopic())[1]

	_, err := dbConnection.Exec(
		`SELECT * FROM upsert_device_status($1, $2);`,
		naturalID,
		string(message.GetBody()),
	)
	return err
}
