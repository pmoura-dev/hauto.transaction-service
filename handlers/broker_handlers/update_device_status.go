package broker_handlers

import (
	"database/sql"
	"regexp"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/hauto.transaction-service/dataaccess"
)

var (
	statusTopicRE = regexp.MustCompile(`state.(\d+)`)
)

func UpdateDeviceState(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	dbConnection := ctx.Params["database"].(*sql.DB)

	deviceID := statusTopicRE.FindStringSubmatch(message.GetTopic())[1]

	return dataaccess.UpsertDeviceState(dbConnection, deviceID, string(message.GetBody()))
}
