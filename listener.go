package main

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/receive"
	"github.com/seagullbird/headr-k8s-helper/client"
	"github.com/streadway/amqp"
)

func makeNewSiteServerListener(c client.Client, logger log.Logger) receive.Listener {
	return func(delivery amqp.Delivery) {
		var event mq.NewSiteEvent
		err := json.Unmarshal(delivery.Body, &event)
		if err != nil {
			logger.Log("error_desc", "Failed to unmarshal event", "error", err, "raw-message:", delivery.Body)
			return
		}
		logger.Log("info", "Received newsite event", "event", event)

		// Create caddy service
		err = c.CreateCaddyService(event.Email, event.SiteName)
		if err != nil {
			logger.Log("error_desc", "Failed to create caddy service", "error", err)
		}
	}
}

func makeDelSiteServerListener(c client.Client, logger log.Logger) receive.Listener {
	return func(delivery amqp.Delivery) {
		var event mq.DelSiteEvent
		err := json.Unmarshal(delivery.Body, &event)
		if err != nil {
			logger.Log("error_desc", "Failed to unmarshal event", "error", err, "raw-message:", delivery.Body)
			return
		}
		logger.Log("info", "Received delsite event", "event", event)

		// Delete caddy service
		err = c.DeleteCaddyService(event.Email, event.SiteName)
		if err != nil {
			logger.Log("error_desc", "Failed to delete caddy service", "error", err)
		}
	}
}
