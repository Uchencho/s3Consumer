package app

import (
	"encoding/json"

	"github.com/Uchencho/commons/pubsub"
	"github.com/Uchencho/s3Consumer/internal/pkg"
	"github.com/Uchencho/s3Consumer/internal/workflow"
	"github.com/pkg/errors"
)

// HandleUploadConsumer handles an SQS event of file upload type
func HandleUploadConsumer() pubsub.ConsumerFunc {
	return func(msg pubsub.Message) error {
		var payload pkg.User

		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			return errors.Wrapf(err, "consumer - unable to decode user request %v", payload)
		}

		handleFile := workflow.HandleZipFile()
		if err := handleFile(payload); err != nil {
			return errors.Wrapf(err, "consumer - unable to handle data %v", payload)
		}
		return nil
	}
}
