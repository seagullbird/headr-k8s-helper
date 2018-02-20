package main

import (
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/receive"
	"github.com/go-kit/kit/log"
	"github.com/streadway/amqp"
	"encoding/json"
)

func makeNewSiteServerListener(logger log.Logger) receive.Listener {
	return func(delivery amqp.Delivery) {
		var event mq.NewSiteEvent
		err := json.Unmarshal(delivery.Body, &event)
		if err != nil {
			logger.Log("error_desc", "Failed to unmarshal event","error", err, "raw-message:", delivery.Body)
			return
		}
		logger.Log("info", "Received newsite event", "event", event)


	}
}
