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

func UpdateDeviceStatus(ctx gobroker.ConsumerContext, message gobroker.Message) error {
	dbConnection := ctx.Params["database"].(*sql.DB)

	naturalID := statusTopicRE.FindStringSubmatch(message.GetTopic())[1]

	return dataaccess.UpsertDeviceStatus(dbConnection, naturalID, string(message.GetBody()))
}
