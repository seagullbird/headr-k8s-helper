package main

import (
	"github.com/go-kit/kit/log"
	"github.com/seagullbird/headr-common/mq"
	"github.com/seagullbird/headr-common/mq/receive"
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
		servername = mq.MQSERVERNAME
		username = mq.MQUSERNAME
		passwd = mq.MQSERVERPWD
	)
	conn, err := mq.MakeConn(servername, username, passwd)
	if err != nil {
		logger.Log("error_desc", "mq.MakeConn failed", "error", err)
		return
	}
	receiver, err := receive.NewReceiver(conn, logger)
	if err != nil {
		logger.Log("error_desc", "receive.NewReceiver failed", "error", err)
		return
	}

	// Register listener
	receiver.RegisterListener("new_site_server", makeNewSiteServerListener(logger))

	// Run forever
	forever := make(chan bool)
	<-forever
}