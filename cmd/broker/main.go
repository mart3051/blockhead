package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"

	"github.com/cloudfoundry-incubator/blockhead/pkg/broker"
	"github.com/cloudfoundry-incubator/blockhead/pkg/config"
)

func main() {
	logger := lager.NewLogger("blockhead-broker")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	if len(os.Args) < 3 {
		logger.Fatal("main", errors.New("config file and/or service directory missing"))
	}

	configFilepath := os.Args[1]
	servicePath := os.Args[2]
	state, err := config.NewState(configFilepath, servicePath)
	if err != nil {
		logger.Fatal("main", err)
	}

	broker := broker.NewBlockheadBroker(state)
	creds := brokerapi.BrokerCredentials{
		Username: state.Config.Username,
		Password: state.Config.Password,
	}
	brokerAPI := brokerapi.New(broker, logger, creds)

	http.Handle("/", brokerAPI)
	logger.Fatal("http-listen", http.ListenAndServe(fmt.Sprintf(":%d", state.Config.Port), nil))
}
