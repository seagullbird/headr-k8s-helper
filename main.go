package main

import (
	"github.com/go-kit/kit/log"
	mqclient "github.com/seagullbird/headr-common/mq/client"
	"github.com/seagullbird/headr-common/mq/receive"
	"github.com/seagullbird/headr-k8s-helper/client"
	"os"
)

func main() {
	// logging domain
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// mq receiver
	var (
		servername = os.Getenv("RABBITMQ_SERVER")
		username   = os.Getenv("RABBITMQ_USER")
		passwd     = os.Getenv("RABBITMQ_PASS")
	)
	receiver, err := receive.NewReceiver(mqclient.New(servername, username, passwd), logger)
	if err != nil {
		logger.Log("error_desc", "receive.NewReceiver failed", "error", err)
		return
	}
	//	new k8s client
	c, err := client.NewClient(logger)
	if err != nil {
		logger.Log("error_desc", "failed to create k8s client", "error", err)
	}

	// Register listeners
	receiver.RegisterListener("new_site_server", makeNewSiteServerListener(c, logger))
	receiver.RegisterListener("del_site_server", makeDelSiteServerListener(c, logger))
	// Run forever
	forever := make(chan bool)
	<-forever
}
